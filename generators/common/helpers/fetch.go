package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/migratooor/tokenLists/generators/common/logs"
)

func FetchJSON[T any](uri string) (data T) {
	var resp *http.Response
	var err error

	u, err := url.Parse(uri)
	if err != nil {
		logs.Error(err)
		return data
	}

	if u.Hostname() == `api.portals.fi` ||
		u.Hostname() == `api.1inch.io` {
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		resp, err = http.DefaultClient.Do(req)
	} else if u.Hostname() == `api.1inch.dev` {
		req, _ := http.NewRequest("GET", uri, nil)
		onInchBearerFromEnv := os.Getenv("BEARER_FOR_1INCH")
		req.Header.Set("Authorization", "Bearer "+onInchBearerFromEnv)
		req.Header.Set("Content-Type", "application/json")
		resp, err = http.DefaultClient.Do(req)
	} else {
		resp, err = http.Get(uri)
	}
	if err != nil {
		logs.Error(err)
		return data
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(`Error reading body for URI ` + uri + `: ` + err.Error())
		return data
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`Error status code for URI ` + uri + `: ` + resp.Status)
		return data
	}

	if err := json.Unmarshal(body, &data); err != nil {
		logs.Error(`Error unmarshal body for URI ` + uri + `: ` + err.Error())
		return data
	}
	return data
}

type FetchJSONPayload struct{}

func FetchJSONPost[T any](uri string) (data T) {
	payloadData := FetchJSONPayload{}
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		logs.Error(err)
		return data
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		logs.Error(err)
		return data
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.Error(err)
		return data
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(`Error reading body for URI ` + uri + `: ` + err.Error())
		return data
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`Error status code for URI ` + uri + `: ` + resp.Status)
		return data
	}

	if err := json.Unmarshal(respBody, &data); err != nil {
		logs.Error(`Error unmarshal body for URI ` + uri + `: ` + err.Error())
		return data
	}
	return data
}
