package mailosaur

type MessagesOperations struct {
	client *MailosaurClient
}

func newMessagesOperations(client *MailosaurClient) *MessagesOperations {
	return &MessagesOperations{
		client: client,
	}
}

func (op *MessagesOperations) Get(id string) (*Message, error) {
	var message *Message
	err := op.client.get("messages/"+id, &message)
	return message, err
}

func (op *MessagesOperations) Delete(id string) error {
	return op.client.delete("messages/" + id)
}

// TODO paging
func (op *MessagesOperations) List(server string) (MessageListResult, error) {
	var messageList MessageListResult
	err := op.client.get("messages?server="+server, &messageList)
	return messageList, err
}

func (op *MessagesOperations) DeleteAll(server string) error {
	return op.client.delete("messages?server=" + server)
}

//TODO pagination
func (op *MessagesOperations) Search(server string, criteria SearchCriteria) (MessageListResult, error) {
	var messageList MessageListResult
	err := op.client.post("messages/search?server="+server, criteria, &messageList)
	return messageList, err
}

func (op *MessagesOperations) WaitFor(server string, criteria SearchCriteria) (*Message, error) {
	var message *Message
	err := op.client.post("messages/await?server="+server, criteria, &message)
	return message, err
}
