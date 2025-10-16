package entity

type License struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	SpdxID      string `json:"spdx_id"`
	path        string
	Summary     string   `json:"description"`
	Body        string   `json:"body"`
	Permissions []string `json:"permissions"`
	Conditions  []string `json:"conditions"`
	Limitations []string `json:"limitations"`
}

func (l License) Title() string       { return l.Name }
func (l License) Description() string { return l.Summary }
func (l License) FilterValue() string { return l.Name }
