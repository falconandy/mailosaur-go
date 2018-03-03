package mailosaur

type MailosaurError struct {
	Type     string
	Messages map[string]string
	Model    interface{}
}

type MailosaurException struct {
	Message        string
	MailosaurError MailosaurError
}

func (ex MailosaurException) Error() string {
	return ex.Message
}
