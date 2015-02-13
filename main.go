package mixpanel

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const baseURL = "http://api.mixpanel.com"

type Profile struct {
	DistinctId string
	Ip         string
	Time       int
	IgnoreTime bool
	Values     map[string]string
	Unset      []string
	Union      map[string][]string
}
type SetProfile struct {
	Token      string            `json:"$token"`
	DistinctId string            `json:"$distinct_id"`
	Ip         string            `json:"$ip,omitempty"`
	Time       int               `json:"$time,omitempty"`
	IgnoreTime bool              `json:"$ignore_time,omitempty"`
	Set        map[string]string `json:"$set"`
}
type SetOnceProfile struct {
	Token      string            `json:"$token"`
	DistinctId string            `json:"$distinct_id"`
	Ip         string            `json:"$ip,omitempty"`
	Time       int               `json:"$time,omitempty"`
	IgnoreTime bool              `json:"$ignore_time,omitempty"`
	SetOnce    map[string]string `json:"$set_once"`
}
type AddProfile struct {
	Token      string            `json:"$token"`
	DistinctId string            `json:"$distinct_id"`
	Ip         string            `json:"$ip,omitempty"`
	Time       int               `json:"$time,omitempty"`
	IgnoreTime bool              `json:"$ignore_time,omitempty"`
	Add        map[string]string `json:"$add"`
}
type UnionProfile struct {
	Token      string              `json:"$token"`
	DistinctId string              `json:"$distinct_id"`
	Ip         string              `json:"$ip,omitempty"`
	Time       int                 `json:"$time,omitempty"`
	IgnoreTime bool                `json:"$ignore_time,omitempty"`
	Union      map[string][]string `json:"$union"`
}
type UnsetProfile struct {
	Token      string   `json:"$token"`
	DistinctId string   `json:"$distinct_id"`
	Ip         string   `json:"$ip,omitempty"`
	Time       int      `json:"$time,omitempty"`
	IgnoreTime bool     `json:"$ignore_time,omitempty"`
	Unset      []string `json:"$unset"`
}
type DeleteProfile struct {
	Token      string `json:"$token"`
	DistinctId string `json:"$distinct_id"`
	Delete     string `json:"$delete"`
}

type Event struct {
	Event      string            `json:"event"`
	Properties map[string]string `json:"properties"`
}
type Mixpanel struct {
	Token string
}

func New(token string) *Mixpanel {
	return &Mixpanel{Token: token}
}

func (m *Mixpanel) Track(event *Event) error {
	if _, ok := event.Properties["token"]; ok {
		return errors.New("You can't put token as a properties")
	}
	event.Properties["token"] = m.Token
	data, err := marshal(event)
	if err != nil {
		return err
	}
	err = request("/track", data)
	if err != nil {
		return err
	}

	return nil
}
func marshal(data interface{}) (std string, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	std = base64.StdEncoding.EncodeToString(jsonData)
	return std, nil
}
func request(reqType string, data string) error {
	url := baseURL + reqType + "?data=" + data
	_, err := http.Get(url)
	if err != nil {
		return err
	}
	return nil

}
func (m *Mixpanel) Set(profile *Profile) error {
	set := SetProfile{
		Token:      m.Token,
		DistinctId: profile.DistinctId,
		Ip:         profile.Ip,
		Time:       profile.Time,
		IgnoreTime: profile.IgnoreTime,
		Set:        profile.Values,
	}
	data, err := marshal(set)
	if err != nil {
		return err
	}
	err = request("/engage", data)
	if err != nil {
		return err
	}
	return nil
}
func (m *Mixpanel) Unset(profile *Profile) error {
	set := UnsetProfile{
		Token:      m.Token,
		DistinctId: profile.DistinctId,
		Ip:         profile.Ip,
		Time:       profile.Time,
		IgnoreTime: profile.IgnoreTime,
		Unset:      profile.Unset,
	}
	data, err := marshal(set)
	if err != nil {
		return err
	}
	err = request("/engage", data)
	if err != nil {
		return err
	}
	return nil
}
func (m *Mixpanel) SetOnce(profile *Profile) error {
	set := SetOnceProfile{
		Token:      m.Token,
		DistinctId: profile.DistinctId,
		Ip:         profile.Ip,
		Time:       profile.Time,
		IgnoreTime: profile.IgnoreTime,
		SetOnce:    profile.Values,
	}
	data, err := marshal(set)
	if err != nil {
		return err
	}
	err = request("/engage", data)
	if err != nil {
		return err
	}
	return nil
}
func (m *Mixpanel) Add(profile *Profile) error {
	set := AddProfile{
		Token:      m.Token,
		DistinctId: profile.DistinctId,
		Ip:         profile.Ip,
		Time:       profile.Time,
		IgnoreTime: profile.IgnoreTime,
		Add:        profile.Values,
	}
	data, err := marshal(set)
	if err != nil {
		return err
	}
	err = request("/engage", data)
	if err != nil {
		return err
	}
	return nil
}
func (m *Mixpanel) Union(profile *Profile) error {
	set := UnionProfile{
		Token:      m.Token,
		DistinctId: profile.DistinctId,
		Ip:         profile.Ip,
		Time:       profile.Time,
		IgnoreTime: profile.IgnoreTime,
		Union:      profile.Union,
	}
	data, err := marshal(set)
	if err != nil {
		return err
	}
	err = request("/engage", data)
	if err != nil {
		return err
	}
	return nil
}
func (m *Mixpanel) Delete(profile *Profile) error {
	set := DeleteProfile{
		Token:      m.Token,
		DistinctId: profile.DistinctId,
		Delete:     "",
	}
	data, err := marshal(set)
	if err != nil {
		return err
	}
	err = request("/engage", data)
	if err != nil {
		return err
	}
	return nil
}
