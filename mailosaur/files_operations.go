package mailosaur

type FilesOperations struct {
	client *MailosaurClient
}

func newFilesOperations(client *MailosaurClient) *FilesOperations {
	return &FilesOperations{
		client: client,
	}
}

// Downloads a single attachment. Simply supply the unique identifier for the required attachment.
func (op *FilesOperations) GetAttachment(id string) ([]byte, error) {
	var attachment []byte
	err := op.client.get("files/attachments/"+id, &attachment)
	return attachment, err
}

//Downloads an EML file representing the specified email. Simply supply the unique identifier for the required email.
func (op *FilesOperations) GetEmail(id string) ([]byte, error) {
	var email []byte
	err := op.client.get("files/email/"+id, &email)
	return email, err
}
