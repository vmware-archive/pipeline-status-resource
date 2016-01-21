package concourse

import (
	"log"
	"net/http"
)

type RequestDoer interface {
	DoRequest(method, url string) *http.Response
}

type ConcourseRequestDoer struct {
	Url      string
	Username string
	Password string
}

func NewRequestDoer(url, username, password string) RequestDoer {
	return &ConcourseRequestDoer{Url: url, Username: username, Password: password}
}

func (api *ConcourseRequestDoer) DoRequest(method, path string) *http.Response {
	req, err := http.NewRequest(method, api.Url+path, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(api.Username, api.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal(err)
	}

	return resp
}
