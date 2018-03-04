package mailosaur

type ServersOperations struct {
	client *MailosaurClient
}

func newServersOperations(client *MailosaurClient) *ServersOperations {
	return &ServersOperations{
		client: client,
	}
}

// Returns a list of your virtual SMTP servers. Servers are returned sorted in alphabetical order.
func (op *ServersOperations) List() (ServerListResult, error) {
	var serverList ServerListResult
	err := op.client.get("servers", &serverList)
	return serverList, err
}

// Creates a new virtual SMTP server and returns it.
func (op *ServersOperations) Create(options ServerCreateOptions) (*Server, error) {
	var server *Server
	err := op.client.post("servers", options, &server)
	return server, err
}

// Retrieves the detail for a single server.
// Simply supply the unique identifier for the required server.
func (op *ServersOperations) Get(id string) (*Server, error) {
	var server *Server
	err := op.client.get("servers/"+id, &server)
	return server, err
}

// Updates a single server and returns it.
func (op *ServersOperations) Update(id string, server *Server) (*Server, error) {
	var newServer *Server
	err := op.client.put("servers/"+id, server, &newServer)
	return newServer, err
}

// Permanently deletes a server. This operation cannot be undone.
// Also deletes all messages and associated attachments within the server.
func (op *ServersOperations) Delete(id string) error {
	return op.client.delete("servers/" + id)
}
