package mailosaur_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MailosaurClient struct {
	apiKey  string
	baseURL string
	servers *ServersOperations
}

func NewMailosaurClient(apiKey string, baseURL string) *MailosaurClient {
	if baseURL == "" {
		baseURL = "https://mailosaur.com/api/" // TODO api? /?
	}
	client := &MailosaurClient{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
	client.servers = newServersOperations(client)
	return client
}

func (client *MailosaurClient) Servers() *ServersOperations {
	return client.servers
}

func (client *MailosaurClient) get(apiPath string) ([]byte, error) {
	return client.doRequest("GET", apiPath, nil)
}

func (client *MailosaurClient) post(apiPath string, data interface{}) ([]byte, error) {
	return client.doRequest("POST", apiPath, data)
}

func (client *MailosaurClient) put(apiPath string, data interface{}) ([]byte, error) {
	return client.doRequest("PUT", apiPath, data)
}

func (client *MailosaurClient) delete(apiPath string) ([]byte, error) {
	return client.doRequest("DELETE", apiPath, nil)
}

func (client *MailosaurClient) doRequest(method, apiPath string, data interface{}) ([]byte, error) {
	var content []byte
	if data != nil {
		var err error
		content, err = json.Marshal(data)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, client.baseURL+apiPath, bytes.NewReader(content))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.apiKey, "")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := http.DefaultClient.Do(req) // TODO: default client
	if err != nil {
		return nil, err
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
		return nil, ex
	}
	return body, err
}
