package apptest

import (
	"bytes"
	crownsAppServer "crowns/app/server"
	"crowns/config"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

func UnmarshallRequest(responseStruct interface{}, r *http.Response) (err error) {
	if r.Body != nil {
		err = json.NewDecoder(r.Body).Decode(&responseStruct)
	}
	return err
}

func ReadTextFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func ExecQueryFile(db *sqlx.DB, path string) error {
	queryContent, err := ReadTextFile(path)
	if err != nil {
		return err
	}
	_, err = db.Exec(queryContent)
	return err
}

func Prepare() *httprouter.Router {
	server := &crownsAppServer.SimpleServer{}
	server.Init(config.Load("test"))
	router := server.Router()
	db := server.GetDbForTest()

	appTestDir := os.Getenv("APP_TEST_DIR")
	if appTestDir == "" {
		panic(errors.New("please set APP_TEST_DIR"))
	}

	err := ExecQueryFile(db, appTestDir+"/migration/down.sql")
	if err != nil {
		panic(err)
	}

	err = ExecQueryFile(db, appTestDir+"/migration/up.sql")
	if err != nil {
		panic(err)
	}

	return router
}

func SendRequest(router *httprouter.Router, method string, path string, request string) (res *http.Response, err error) {
	var req *http.Request
	if method == "GET" {
		req, err = http.NewRequest(method, path, nil)
	} else {
		req, err = http.NewRequest(method, path, bytes.NewBufferString(request))
	}

	if err != nil {
		return nil, err
	}
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec.Result(), err
}
