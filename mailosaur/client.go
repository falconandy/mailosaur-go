package mailosaur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type MailosaurClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client

	servers  *ServersOperations
	messages *MessagesOperations
	analysis *AnalysisOperations
	files    *FilesOperations
}

func NewMailosaurClient(apiKey string, baseURL string) *MailosaurClient {
	if baseURL == "" {
		baseURL = "https://mailosaur.com"
	}
	client := &MailosaurClient{
		apiKey:     apiKey,
		baseURL:    strings.TrimSuffix(baseURL, "/"),
		httpClient: http.DefaultClient,
	}
	client.servers = newServersOperations(client)
	client.messages = newMessagesOperations(client)
	client.analysis = newAnalysisOperations(client)
	client.files = newFilesOperations(client)
	return client
}

func (client *MailosaurClient) SetHttpClient(httpClient *http.Client) {
	client.httpClient = httpClient
}

func (client *MailosaurClient) Servers() *ServersOperations {
	return client.servers
}

func (client *MailosaurClient) Messages() *MessagesOperations {
	return client.messages
}

func (client *MailosaurClient) Analysis() *AnalysisOperations {
	return client.analysis
}

func (client *MailosaurClient) Files() *FilesOperations {
	return client.files
}

func (client *MailosaurClient) get(apiPath string, responseData interface{}) error {
	return client.doRequest("GET", apiPath, nil, responseData)
}

func (client *MailosaurClient) post(apiPath string, requestData interface{}, responseData interface{}) error {
	return client.doRequest("POST", apiPath, requestData, responseData)
}

func (client *MailosaurClient) put(apiPath string, requestData interface{}, responseData interface{}) error {
	return client.doRequest("PUT", apiPath, requestData, responseData)
}

func (client *MailosaurClient) delete(apiPath string) error {
	return client.doRequest("DELETE", apiPath, nil, nil)
}

func (client *MailosaurClient) doRequest(method, apiPath string, requestData interface{}, responseData interface{}) error {
	var content []byte
	if requestData != nil {
		var err error
		content, err = json.Marshal(requestData)
		if err != nil {
			return err
		}
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s/api/%s", client.baseURL, apiPath), bytes.NewReader(content))
	if err != nil {
		return err
	}
	req.SetBasicAuth(client.apiKey, "")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if 400 <= resp.StatusCode && resp.StatusCode < 600 {
		ex := MailosaurException{
			Message: fmt.Sprintf("Operation returned an invalid status code '%s'", resp.Status),
		}
		var mErr MailosaurError
		err := json.Unmarshal(body, &mErr)
		if err == nil {
			ex.MailosaurError = mErr
		}
		return ex
	}
	if resp.StatusCode == 204 {
		return nil
	}
	if responseBytes, ok := responseData.(*[]byte); ok {
		*responseBytes = body
		return nil
	}
	return json.Unmarshal(body, responseData)
}
