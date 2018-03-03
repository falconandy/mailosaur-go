package mailosaur_go

type ServerListResult struct {
	Items []*Server
}

type Server struct {
	ID              string
	Password        string
	Name            string
	Users           []string
	Messages        int
	ForwardingRules []ForwardingRule
}

type ForwardingRule struct {
	Field     string
	Operator  string
	Value     string
	ForwardTo string
}

type ServerCreateOptions struct {
	Name string
}
