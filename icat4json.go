package icat4json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//ErrAPI - API error icat for JSON
type ErrAPI string

func (e ErrAPI) Error() string {
	return "API error: " + string(e)
}

//Msg - result message from icat
type msg struct {
	Messageid string `json:"messageid"`
	Message   string `json:"message"`
}

func (m msg) String() string {
	return fmt.Sprintf("%s (%s)", m.Message, m.Messageid)
}

//Item - itemdata from icat
type Item struct {
	Title      string    `json:"item_title"`
	Link       string    `json:"item_link"`
	Date       time.Time `json:"item_date"`
	Identifier []string  `json:"item_identifier"`
}

//ICAT - data from icat
type ICAT struct {
	Itemdata []Item    `json:"itemdata"`
	Title    string    `json:"docTitle"`
	Fix      string    `json:"docTitleFix"`
	Link     string    `json:"docLink"`
	Date     time.Time `json:"docDate"`
}

//JSON - string with JSON format
type JSON string

func (j JSON) String() string {
	return string(j)
}

//getError returns status if error
func (j JSON) getError() error {
	if strings.Index(j.String(), "messageid") < 0 {
		return nil
	}
	var m msg
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return ErrAPI(err.Error()) //another format?
	}
	return ErrAPI(m.String())
}

//Decode returns icat data
func (j JSON) Decode() (*ICAT, error) {
	if strings.Index(j.String(), "itemdata") < 0 {
		return nil, ErrAPI("not icat data")
	}
	var i ICAT
	if err := json.Unmarshal([]byte(j), &i); err != nil {
		return nil, ErrAPI(err.Error())
	}
	return &i, nil
}

//Tool options for icat
const (
	ToolICATW = "icatw"
	ToolICATH = "icath"
	urlAPI    = "https://isec-myjvn-feed1.ipa.go.jp/IPARssReader.php"
)

//Get returns JSON string from icat
func Get(tool string) (JSON, error) {
	values := url.Values{
		"tool": {tool},
	}
	ux := strconv.FormatInt(time.Now().Unix(), 10)

	resp, err := http.Get(fmt.Sprintf("%s?%s&%s", urlAPI, ux, values.Encode()))
	if err != nil {
		return "", ErrAPI(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", ErrAPI(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", ErrAPI(err.Error())
	}
	j := JSON(body)
	if err := j.getError(); err != nil {
		return j, err
	}
	return j, nil
}
