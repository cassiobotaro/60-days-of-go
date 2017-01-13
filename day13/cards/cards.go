package cards

// Card is a item in todo list
type Card struct {
	Title string `json:"title" valid:"alphanum,required"`
	Text  string `json:"text" valid:"alphanum,required"`
	Done  bool   `json:"done"`
	ID    int64  `json:"id,omitempty"`
}
