package mailosaur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	defaultBaseUrl = "https://mailosaur.com"
)

type MailosaurClient struct {
	apiKey     string
	baseUrl    string
	httpClient *http.Client

	servers  *ServersOperations
	messages *MessagesOperations
	analysis *AnalysisOperations
	files    *FilesOperations
}

func NewMailosaurClient(apiKey string) *MailosaurClient {
	client := &MailosaurClient{
		apiKey:     apiKey,
		baseUrl:    defaultBaseUrl,
		httpClient: http.DefaultClient,
	}
	client.servers = newServersOperations(client)
	client.messages = newMessagesOperations(client)
	client.analysis = newAnalysisOperations(client)
	client.files = newFilesOperations(client)
	return client
}

func (client *MailosaurClient) SetBaseUrl(baseUrl string) {
	if baseUrl != "" {
		client.baseUrl = strings.TrimSuffix(baseUrl, "/")
	} else {
		client.baseUrl = defaultBaseUrl
	}
}

func (client *MailosaurClient) SetHttpClient(httpClient *http.Client) {
	if httpClient != nil {
		client.httpClient = httpClient
	} else {
		client.httpClient = http.DefaultClient
	}
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
	return client.doApiRequest("GET", apiPath, nil, responseData)
}

func (client *MailosaurClient) post(apiPath string, requestData interface{}, responseData interface{}) error {
	return client.doApiRequest("POST", apiPath, requestData, responseData)
}

func (client *MailosaurClient) put(apiPath string, requestData interface{}, responseData interface{}) error {
	return client.doApiRequest("PUT", apiPath, requestData, responseData)
}

func (client *MailosaurClient) delete(apiPath string) error {
	return client.doApiRequest("DELETE", apiPath, nil, nil)
}

func (client *MailosaurClient) doApiRequest(method, apiPath string, requestData interface{}, responseData interface{}) error {
	req, err := client.createApiRequest(method, apiPath, requestData)
	if err != nil {
		return err
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	return client.processApiResponse(resp, responseData)
}

func (client *MailosaurClient) createApiRequest(method, apiPath string, requestData interface{}) (*http.Request, error) {
	var content []byte
	if requestData != nil {
		var err error
		content, err = json.Marshal(requestData)
		if err != nil {
			return nil, err
		}
	}
	apiUrl := fmt.Sprintf("%s/api/%s", client.baseUrl, apiPath)
	req, err := http.NewRequest(method, apiUrl, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.apiKey, "")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	return req, nil
}

func (client *MailosaurClient) processApiResponse(resp *http.Response, responseData interface{}) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 204 {
		return nil
	}

	if resp.StatusCode == 200 {
		if responseBytes, ok := responseData.(*[]byte); ok {
			*responseBytes = body
			return nil
		}
		return json.Unmarshal(body, responseData)
	}

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

	return fmt.Errorf("unsupported status code: %s", resp.Status)
}
