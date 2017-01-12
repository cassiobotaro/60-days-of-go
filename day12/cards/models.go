package cards

// Card is a item in todo list
type Card struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Done  bool   `json:"done"`
}
