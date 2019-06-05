package request

type GetUser struct {
	Name  string `json:"name",omitempty`
	Token string `json:"token",omitempty`
}

type CreateUser struct {
	Name string `json:"name"`
}

type StartGame struct {
	Token string `json:"token",omitempty`
}
