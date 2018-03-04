package mailosaur

type ForwardingRule struct {
	Field     string
	Operator  string
	Value     string
	ForwardTo string
}

type Server struct {
	Id              string
	Password        string
	Name            string
	Users           []string
	Messages        int
	ForwardingRules []ForwardingRule
}

type ServerCreateOptions struct {
	Name string
}

type ServerListResult struct {
	Items []*Server
}
