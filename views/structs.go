package views

type Response struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}

type PostRequest struct {
	Name string `json:"name"`
	Task string `json:"task"`
}
