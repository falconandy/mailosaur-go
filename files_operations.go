package mailosaur_go

type FilesOperations struct {
	client *MailosaurClient
}

func newFilesOperations(client *MailosaurClient) *FilesOperations {
	return &FilesOperations{
		client: client,
	}
}

func (op *FilesOperations) GetAttachment(id string) ([]byte, error) {
	var attachment []byte
	err := op.client.get("files/attachments/"+id, &attachment)
	return attachment, err
}

func (op *FilesOperations) GetEmail(id string) ([]byte, error) {
	var email []byte
	err := op.client.get("files/email/"+id, &email)
	return email, err
}
