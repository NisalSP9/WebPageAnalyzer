package models

type HTMLPAGEINFOR struct {
	HTMLVersion  string `json:"HTMLVersion"`
	PageTitle    string `json:"pageTitle"`
	Headings     map[string]int `json:"headings"`
	Internal     int    `json:"internal"`
	External     int    `json:"external"`
	Inaccessible int    `json:"inaccessible"`
	LoginForm    bool   `json:"loginForm"`
}
