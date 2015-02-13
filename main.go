package mixpanel

import (
	"net/http"
	"encoding/base64"
	"errors"
	"encoding/json"
)
const baseURL = "http://api.mixpanel.com"

type Event struct {
	Event string `json:"event"`
	Properties map[string]string `json:"properties"`
}
type Mixpanel struct {
	Token string
}
func New(token string) *Mixpanel{
	return &Mixpanel{Token: token}
}

func (m *Mixpanel) track(event *Event) error {
	if _, ok := event.Properties["token"]; ok {
		return errors.New("You can't put token as a properties")
	}
	event.Properties["token"] = m.Token
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}
	data := base64.StdEncoding.EncodeToString(jsonData)

	err = request("/track", data)
	if err != nil {
		return err
	}


	return nil
}
func request(reqType string, data string) error {
	url := baseURL + reqType+ "?data="+data
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	return nil

}
