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

type GetHighScores struct {
	Token string `json:"token",omitempty`
}

type SubmitHighScores struct {
	Token          string `json:"token",omitempty`
	Start          int64  `json:"start",omitempty`
	Age            int64  `json:"age",omitempty`
	ChurchScore    int64  `json:"church_score",omitempty`
	CommersScore   int64  `json:"commers_score",omitempty`
	MerchantsScore int64  `json:"merchants_score",omitempty`
	MilitaryScore  int64  `json:"military_score",omitempty`
	BonusScore     int64  `json:"bonus_score",omitempty`
}
