package mailosaur

import "fmt"

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

func (op *MessagesOperations) ListPage(server string, page int, itemsPerPage int) (MessageListResult, error) {
	apiPath := op.getPagingApiPath("messages?server="+server, page, itemsPerPage)
	var messageList MessageListResult
	err := op.client.get(apiPath, &messageList)
	return messageList, err
}

func (op *MessagesOperations) DeleteAll(server string) error {
	return op.client.delete("messages?server=" + server)
}

func (op *MessagesOperations) Search(server string, criteria SearchCriteria) (MessageListResult, error) {
	var messageList MessageListResult
	err := op.client.post("messages/search?server="+server, criteria, &messageList)
	return messageList, err
}

func (op *MessagesOperations) SearchPage(server string, criteria SearchCriteria, page int, itemsPerPage int) (MessageListResult, error) {
	apiPath := op.getPagingApiPath("messages/search?server="+server, page, itemsPerPage)
	var messageList MessageListResult
	err := op.client.post(apiPath, criteria, &messageList)
	return messageList, err
}

func (op *MessagesOperations) WaitFor(server string, criteria SearchCriteria) (*Message, error) {
	var message *Message
	err := op.client.post("messages/await?server="+server, criteria, &message)
	return message, err
}

func (op *MessagesOperations) getPagingApiPath(baseApiPath string, page int, itemsPerPage int) string {
	result := baseApiPath
	if page != 0 {
		result += fmt.Sprintf("&page=%d", page)
	}
	if itemsPerPage != 0 {
		result += fmt.Sprintf("&itemsPerPage=%d", itemsPerPage)
	}
	return result
}
