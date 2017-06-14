package httpclient

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	Requestinfo struct {
		Url      string
		User     string
		Password string
	}
)

func PostJsontoTarget(info *Requestinfo, body []byte) (string, error) {
	log.Printf("Post body: %s", string(body))

	client := &http.Client{}
	req, err := http.NewRequest("POST", info.Url, bytes.NewReader(body))
	if err != nil {
		log.Printf("")
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")

	if info.User != "" && info.Password != "" {
		req.SetBasicAuth(info.User, info.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("")
		return "", err
	}
	defer resp.Body.Close()
	ret := getResponse(resp)
	if resp.StatusCode != 200 {
		return "", errors.New(ret)
	}

	return ret, nil
}

func getResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	var ret string
	if err == nil {
		ret = string(b)
	} else {
		ret = "Fatal!! ERROR!!"
	}
	return ret
}
