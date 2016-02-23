package icat4json

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//ErrAPI - API error icat for JSON
type ErrAPI string

func (e ErrAPI) Error() string {
	return "API error: " + string(e)
}

//Tool options for icat
const (
	ToolICATW = "icatw"
	ToolICATH = "icath"
	urlAPI    = "https://isec-myjvn-feed1.ipa.go.jp/IPARssReader.php"
)

//Get returs JSON data for icat
func Get(tool string) (string, error) {
	values := url.Values{
		"tool": {tool},
	}
	ux := strconv.FormatInt(time.Now().Unix(), 10)

	resp, err := http.Get(fmt.Sprintf("%s?%s&%s", urlAPI, ux, values.Encode()))
	if err != nil {
		return "", ErrAPI(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", ErrAPI(err.Error())
	}
	return string(body), nil
}
