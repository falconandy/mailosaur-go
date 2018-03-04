package mailosaur

import "time"

type Attachment struct {
	Id          string
	ContentType string
	FileName    string
	ContentId   string
	Length      int
	Url         string
}

type Image struct {
	Src string
	Alt string
}

type Link struct {
	Href string
	Text string
}

type Message struct {
	Id          string
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

type MessageAddress struct {
	Name  string
	Email string
	Phone string
}

type MessageContent struct {
	Links  []Link
	Images []Image
	Body   string
}

type MessageHeader struct {
	Field string
	Value string
}

type MessageListResult struct {
	Items []*MessageSummary
}

type MessageSummary struct {
	Id          string
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

type Metadata struct {
	Headers []MessageHeader
}

type SearchCriteria struct {
	SentTo  string
	Subject string
	Body    string
}
