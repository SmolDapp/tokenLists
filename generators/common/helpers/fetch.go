package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/logs"
)

func FetchJSON[T any](uri string) (data T) {
	var resp *http.Response
	var err error

	if strings.Contains(uri, `api.portals.fi`) {
		req, _ := http.NewRequest("GET", uri, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
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
		logs.Error(err)
		return data
	}

	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`failed to fetch: ` + uri)
		return data
	}

	if err := json.Unmarshal(body, &data); err != nil {
		logs.Error(err)
		return data
	}
	return data
}
