package mailosaur_go

import (
	"encoding/json"
)

type ServersOperations struct {
	client *MailosaurClient
}

type serversListResponse struct {
	Items []*Server
}

func newServersOperations(client *MailosaurClient) *ServersOperations {
	return &ServersOperations{
		client: client,
	}
}

func (op *ServersOperations) List() ([]*Server, error) {
	data, err := op.client.get("servers")
	if err != nil {
		return nil, err
	}
	var listResponse serversListResponse
	err = json.Unmarshal(data, &listResponse)
	if err != nil {
		return nil, err
	}
	return listResponse.Items, nil
}

func (op *ServersOperations) Get(id string) (*Server, error) {
	data, err := op.client.get("servers/" + id)
	if err != nil {
		return nil, err
	}
	var server *Server
	err = json.Unmarshal(data, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (op *ServersOperations) Create(options ServerCreateOptions) (*Server, error) {
	data, err := op.client.post("servers", options)
	if err != nil {
		return nil, err
	}
	var server *Server
	err = json.Unmarshal(data, &server)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func (op *ServersOperations) Update(id string, server *Server) (*Server, error) {
	data, err := op.client.put("servers/"+id, server)
	if err != nil {
		return nil, err
	}
	var newServer *Server
	err = json.Unmarshal(data, &newServer)
	if err != nil {
		return nil, err
	}
	return newServer, nil
}

func (op *ServersOperations) Delete(id string) error {
	_, err := op.client.delete("servers/" + id)
	return err
}
