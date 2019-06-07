package apptest

import (
	"crowns/app/domain/response"
	"fmt"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func createUser(router *httprouter.Router, userName string) (*response.User, error) {
	r, err := SendRequest(
		router,
		"POST",
		"/user/create",
		`{
			"name": "`+userName+`"
		}`)
	if err != nil {
		return nil, err
	}

	res := response.User{}
	err = UnmarshallRequest(&res, r)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func getUserByToken(router *httprouter.Router, token string) (*response.User, error) {
	r, err := SendRequest(
		router,
		"POST",
		"/user/get",
		`{
			"token": "`+token+`"
		}`)
	if err != nil {
		return nil, err
	}

	res := response.User{}
	err = UnmarshallRequest(&res, r)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func submitHighScore(router *httprouter.Router, token string, score int64) ([]response.HighScore, error) {
	r, err := SendRequest(
		router,
		"POST",
		"/highscores/submit",
		`{
			"token": "`+token+`",
			"score": `+strconv.FormatInt(score, 10)+`
		}`)
	if err != nil {
		return nil, err
	}

	res := []response.HighScore{}
	err = UnmarshallRequest(&res, r)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getHighScores(router *httprouter.Router) ([]response.HighScore, error) {
	r, err := SendRequest(
		router,
		"GET",
		"/highscores",
		"")
	if err != nil {
		return nil, err
	}

	res := []response.HighScore{}
	err = UnmarshallRequest(&res, r)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func TestCreateUser(t *testing.T) {
	router := Prepare()

	userNames := []string{"taro", "hanako", "jiro"}
	for _, userName := range userNames {
		user, err := createUser(router, userName)
		if err != nil {
			t.Errorf(err.Error())
		}
		if user.Name != userName {
			t.Errorf("name you got is invalid")
		}
	}
}

func TestGetUserByToken(t *testing.T) {
	router := Prepare()

	userNames := []string{"taro", "hanako", "jiro"}
	users := []*response.User{}
	userTokens := make(map[string]struct{})
	for _, userName := range userNames {
		user, err := createUser(router, userName)
		if err != nil {
			t.Errorf(err.Error())
		}
		userTokens[user.Token] = struct{}{}
		users = append(users, user)
	}

	if len(userTokens) != len(users) {
		fmt.Println(userTokens)
		t.Errorf("user token is not unique")
	}

	for u, user := range users {
		userByToken, err := getUserByToken(router, user.Token)
		users[u] = userByToken
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}

func TestHighScores(t *testing.T) {
	type userAndScores struct {
		user   *response.User
		scores []int64
	}

	router := Prepare()
	userNames := []string{"いぬ", "ねこ", "とら", "さる", "きじ", "たか", "おおかみ"}
	highScoreLists := [][]int64{
		{10, 60, 30, 90, 20, 15, 80, 90},
		{100, 30, 80, 70, 90, 100, 110, 100},
		{40, 20, 110, 40, 50, 60, 100},
		{20, 70, 80, 10, 90, 80, 130},
		{110, 20, 110, 40, 50, 60, 100},
		{40, 90, 80, 90, 150, 200, 30},
		{40, 20, 50, 140, 20, 30, 150},
	}

	for u, userName := range userNames {
		user, err := createUser(router, userName)
		if err != nil {
			t.Errorf(err.Error())
		}

		for _, score := range highScoreLists[u] {
			submitHighScore(router, user.Token, score)
		}
	}

	res, err := getHighScores(router)
	if err != nil {
		t.Errorf(err.Error())
	}

	truthHighScores := []int64{200, 150, 130, 110, 110}
	for h, highScore := range res {
		if truthHighScores[h] != highScore.Score {
			t.Errorf("excepted: %d, actual: %d", truthHighScores[h], highScore.Score)
		}
	}
}
