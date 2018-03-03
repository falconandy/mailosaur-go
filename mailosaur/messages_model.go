package mailosaur

import "time"

type MessageListResult struct {
	Items []*MessageSummary
}

type MessageSummary struct {
	ID          string
	Server      string
	Rcpt        []MessageAddress
	From        []MessageAddress
	To          []MessageAddress
	Cc          []MessageAddress
	Bcc         []MessageAddress
	Received    time.Time
	Subject     string
	Summary     string
	Attachments int
}

type MessageAddress struct {
	Name  string
	Email string
	Phone string
}

type Message struct {
	ID          string
	Rcpt        []MessageAddress
	From        []MessageAddress
	To          []MessageAddress
	Cc          []MessageAddress
	Bcc         []MessageAddress
	Received    time.Time
	Subject     string
	Html        MessageContent
	Text        MessageContent
	Attachments []Attachment
	Metadata    Metadata
	Server      string
}

type MessageContent struct {
	Links  []Link
	Images []Image
	Body   string
}

type Attachment struct {
	ID          string
	ContentType string
	FileName    string
	ContentId   string
	Length      int
	Url         string
}

type Metadata struct {
	Headers []MessageHeader
}

type Link struct {
	Href string
	Text string
}

type Image struct {
	Src string
	Alt string
}

type MessageHeader struct {
	Field string
	Value string
}

type SearchCriteria struct {
	SentTo  string
	Subject string
	Body    string
}
