package structs

//Request represents a request
type Request struct {
	Method string `json:"method"`
	Path   []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"path"`
	Header []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"header"`
	QueryString []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"query-string"`
	URL     string `json:"url"`
	Payload string `json:"payload"`
}
