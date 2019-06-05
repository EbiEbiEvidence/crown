package request

type GetUser struct {
	Name  string `json:"name",omitempty`
	Token string `json:"token",omitempty`
}

type CreateUser struct {
	Name string `json:"name"`
}
