package mailosaur_go

type ServersOperations struct {
	client *MailosaurClient
}

func newServersOperations(client *MailosaurClient) *ServersOperations {
	return &ServersOperations{
		client: client,
	}
}

func (op *ServersOperations) List() (ServerListResult, error) {
	var serverList ServerListResult
	err := op.client.get("servers", &serverList)
	return serverList, err
}

func (op *ServersOperations) Create(options ServerCreateOptions) (*Server, error) {
	var server *Server
	err := op.client.post("servers", options, &server)
	return server, err
}

func (op *ServersOperations) Get(id string) (*Server, error) {
	var server *Server
	err := op.client.get("servers/"+id, &server)
	return server, err
}

func (op *ServersOperations) Update(id string, server *Server) (*Server, error) {
	var newServer *Server
	err := op.client.put("servers/"+id, server, &newServer)
	return newServer, err
}

func (op *ServersOperations) Delete(id string) error {
	return op.client.delete("servers/" + id)
}
