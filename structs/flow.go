package structs

//Flow ...
type Flow struct {
	Step       int        `json:"step"`
	Validation Validation `json:"validation"`
	Next       int        `json:"next,omitempty"`
	Request    Request    `json:"request"`
}
