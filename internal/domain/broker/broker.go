package broker

type Result struct {
	Name   string   `json:"name"`
	Access []string `json:"access"`
}

type Company struct {
	Company string   `json:"company"`
	Results []Result `json:"results"`
}

type Companies struct {
	Result []Company `json:"result"`
}
