package mailosaur

type AnalysisOperations struct {
	client *MailosaurClient
}

func newAnalysisOperations(client *MailosaurClient) *AnalysisOperations {
	return &AnalysisOperations{
		client: client,
	}
}

func (op *AnalysisOperations) Spam(emailId string) (SpamAnalysisResult, error) {
	var spamResult SpamAnalysisResult
	err := op.client.get("analysis/spam/"+emailId, &spamResult)
	return spamResult, err
}
