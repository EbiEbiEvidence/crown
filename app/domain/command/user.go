package command

type UserCommandModel struct {
	Name  string
	Token string
}

func NewUserCommandModel(name string, token string) *UserCommandModel {
	return &UserCommandModel{
		Name:  name,
		Token: token,
	}
}
