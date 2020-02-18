package graphql

import (
	"log"

	"github.com/imroc/req"
)

type UserInfo struct {
	Data struct {
		CurrentUser struct {
			Name         string `json:"name"`
			StudentEmail string `json:"studentEmail"`
		} `json:"currentUser"`
	} `json:"data"`
}

func GetGraphUserInfo(session *req.Req) (string, string) {

	url := "https://www.makeschool.com/graphql"
	query := "{ currentUser {name studentEmail} }"
	request := map[string]string{"query": query}

	resp, err := session.Post(url, req.BodyJSON(request))
	if err != nil {
		log.Fatal(err)
	}

	var data UserInfo
	resp.ToJSON(&data)

	return data.Data.CurrentUser.Name, data.Data.CurrentUser.StudentEmail
}