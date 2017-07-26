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
	URL     string `json:"url"`
	Payload string `json:"payload"`
}
