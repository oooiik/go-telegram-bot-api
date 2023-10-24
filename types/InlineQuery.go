package types

type InlineQuery struct {
	ID       string   `json:"id"`
	From     User     `json:"from"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
	Location Location `json:"location,omitempty"`
}
