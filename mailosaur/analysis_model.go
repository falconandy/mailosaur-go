package mailosaur

type SpamAnalysisResult struct {
	SpamFilterResults SpamFilterResults
	Score             float64
}

type SpamAssassinRule struct {
	Score       float64
	Rule        string
	Description string
}

type SpamFilterResults struct {
	SpamAssassin []SpamAssassinRule
}
