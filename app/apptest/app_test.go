package apptest

import (
	"crowns/app/domain/request"
	"crowns/app/domain/response"
	"encoding/json"
	"fmt"
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

func submitHighScore(router *httprouter.Router, token string, reqStruct request.SubmitHighScores) ([]response.HighScore, error) {
	req, err := json.Marshal(request.SubmitHighScores{
		Token:          token,
		Start:          reqStruct.Start,
		Age:            reqStruct.Age,
		ChurchScore:    reqStruct.ChurchScore,
		CommersScore:   reqStruct.CommersScore,
		MerchantsScore: reqStruct.MerchantsScore,
		MilitaryScore:  reqStruct.MilitaryScore,
		BonusScore:     reqStruct.BonusScore,
	})
	if err != nil {
		return nil, err
	}

	r, err := SendRequest(
		router,
		"POST",
		"/highscores/submit",
		string(req))
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
	highScoreLists := [][]request.SubmitHighScores{
		{
			request.SubmitHighScores{
				Start:          2193,
				Age:            19,
				ChurchScore:    82,
				CommersScore:   96,
				MerchantsScore: 85,
				MilitaryScore:  90,
				BonusScore:     261,
			},
			request.SubmitHighScores{
				Start:          2194,
				Age:            33,
				ChurchScore:    64,
				CommersScore:   17,
				MerchantsScore: 36,
				MilitaryScore:  99,
				BonusScore:     576,
			},
			request.SubmitHighScores{
				Start:          2029,
				Age:            13,
				ChurchScore:    30,
				CommersScore:   90,
				MerchantsScore: 40,
				MilitaryScore:  35,
				BonusScore:     527,
			},
			request.SubmitHighScores{
				Start:          2144,
				Age:            19,
				ChurchScore:    1,
				CommersScore:   92,
				MerchantsScore: 32,
				MilitaryScore:  11,
				BonusScore:     775,
			},
		},
		{
			request.SubmitHighScores{
				Start:          1956,
				Age:            49,
				ChurchScore:    24,
				CommersScore:   33,
				MerchantsScore: 43,
				MilitaryScore:  22,
				BonusScore:     665,
			},
			request.SubmitHighScores{
				Start:          2034,
				Age:            6,
				ChurchScore:    32,
				CommersScore:   93,
				MerchantsScore: 96,
				MilitaryScore:  17,
				BonusScore:     649,
			},
		},
		{
			request.SubmitHighScores{
				Start:          2064,
				Age:            15,
				ChurchScore:    94,
				CommersScore:   20,
				MerchantsScore: 41,
				MilitaryScore:  19,
				BonusScore:     405,
			},
			request.SubmitHighScores{
				Start:          1917,
				Age:            10,
				ChurchScore:    5,
				CommersScore:   85,
				MerchantsScore: 43,
				MilitaryScore:  44,
				BonusScore:     698,
			},
		},
		{
			request.SubmitHighScores{
				Start:          2052,
				Age:            44,
				ChurchScore:    85,
				CommersScore:   14,
				MerchantsScore: 98,
				MilitaryScore:  8,
				BonusScore:     115,
			},
		},
		{
			request.SubmitHighScores{
				Start:          2084,
				Age:            34,
				ChurchScore:    78,
				CommersScore:   2,
				MerchantsScore: 12,
				MilitaryScore:  32,
				BonusScore:     430,
			},
		},
		{
			request.SubmitHighScores{
				Start:          1905,
				Age:            10,
				ChurchScore:    65,
				CommersScore:   81,
				MerchantsScore: 30,
				MilitaryScore:  75,
				BonusScore:     454,
			},
		},
		{
			request.SubmitHighScores{
				Start:          2065,
				Age:            21,
				ChurchScore:    48,
				CommersScore:   16,
				MerchantsScore: 49,
				MilitaryScore:  36,
				BonusScore:     898,
			},
		},
	}

	for u, userName := range userNames {
		user, err := createUser(router, userName)
		if err != nil {
			t.Errorf(err.Error())
		}

		for _, reqStruct := range highScoreLists[u] {
			submitHighScore(router, user.Token, reqStruct)
		}
	}

	expectedIds := []int64{9, 10, 12, 4, 8}
	res, err := getHighScores(router)
	if err != nil {
		t.Errorf(err.Error())
	}

	for h, highScore := range res {
		if highScore.HighScoreID != expectedIds[h] {
			t.Errorf("expected %d, actual %d", highScore.HighScoreID, expectedIds[h])
		}
	}
}
