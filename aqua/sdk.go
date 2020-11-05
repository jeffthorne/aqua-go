// Package aqua provides interface into Aqua Enterprise CSP Platform
package aqua

import (
	"crypto/tls"
	"encoding/json"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// Aqua is the main interface used to interact with Aqua's Enterprise platform API endpoint.
type Aqua struct {
	Host               string
	Port               int
	ID                 string
	Password           string
	URL                string
	Secure             bool
	Token              string `json:"token"`
	InsecureSkipVerify bool
	RestClient         gorequest.SuperAgent
	LicenseType        string `json:"license_type"`
	Scopes             []string
	User               map[string]interface{} `json:"user"`
	License            map[string]interface{} `json:"license"`
}

// NewCSP - is used to obtain functioning Aqua Enterprise API endpoint
// secureEndpoint is optional. if supplied expecting 2 bool values.
// First value is whether or not Aqua is listenting on secure URL
// Second is whether or not to InsecureSkipVerify
func NewCSP(host string, port int, id string, password string, secureEndpoint ...bool) (*Aqua, error) {

	aqua := Aqua{Host: host, Port: port, ID: id, Password: password, Secure: true, InsecureSkipVerify: true}
	aqua.RestClient = *gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: aqua.InsecureSkipVerify})

	if len(secureEndpoint) > 0 {
		aqua.Secure = secureEndpoint[0]
		if len(secureEndpoint) > 1 {
			aqua.InsecureSkipVerify = secureEndpoint[1]
		}
	}

	if aqua.Secure {
		aqua.URL = fmt.Sprintf("https://%s:%d/api", host, port)
	} else {
		aqua.URL = fmt.Sprintf("http://%s:%d/api", host, port)
	}

	connected, message := authenticate(&aqua)

	if connected {
		aqua.RestClient.Set("Authorization", "Bearer "+aqua.Token)
		return &aqua, nil
	}

	return &aqua, fmt.Errorf(message)
}

func authenticate(aqua *Aqua) (bool, string) {

	params := `{"id":"` + aqua.ID + `", "password":"` + aqua.Password + `"}`
	resp, body, err := api(aqua, "login", params, "v1", "POST")

	if err != nil {
		return false, ""
	}

	if resp.StatusCode == 200 {
		_ = json.Unmarshal([]byte(body), &aqua)
		return true, ""
	}

	return false, fmt.Sprintf("Failed with status: %s", resp.Status)

}

func api(aqua *Aqua, call, params, apiVersion, method string) (gorequest.Response, string, []error) {
	url := fmt.Sprintf("%s/%s/%s", aqua.URL, apiVersion, call)
	var resp gorequest.Response
	var body string
	var err []error

	switch method {
	case "GET":
		resp, body, err = aqua.RestClient.Clone().Get(url).Query(params).End()
	case "POST":
		resp, body, err = aqua.RestClient.Clone().Post(url).Send(params).End()
	case "PUT":
		resp, body, err = aqua.RestClient.Clone().Put(url).Send(params).End()
	case "DELETE":
		resp, body, err = aqua.RestClient.Clone().Delete(url).Send(params).End()
	}



	return resp, body, err

}
