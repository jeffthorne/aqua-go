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
	RestClient         gorequest.SuperAgent
	URL                string
	Secure             bool
	Token              string `json:"token"`
	InsecureSkipVerify bool
	LicenseType        string `json:"license_type"`
	Scopes             []string
	User               map[string]interface{} `json:"user"`
	License            map[string]interface{} `json:"license"`
}

// NewCSP - is used to obtain functioning Aqua Enterprise API endpoint
func NewCSP(host string, port int, id string, password string, secureEndpoint ...bool) (Aqua, error) {

	aqua := Aqua{Host: host, Port: port, ID: id, Password: password, Secure: true, InsecureSkipVerify: true}

	if len(secureEndpoint) > 0 {
		aqua.Secure = secureEndpoint[0]
		if len(secureEndpoint) > 1 {
			aqua.InsecureSkipVerify = secureEndpoint[1]
		}
	}

	aqua.RestClient = *gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: aqua.InsecureSkipVerify})
	if aqua.Secure {
		aqua.URL = fmt.Sprintf("https://%s:%d/api", host, port)
	} else {
		aqua.URL = fmt.Sprintf("http://%s:%d/api", host, port)
	}

	connected, message := authenticate(&aqua)

	if connected {
		return aqua, nil
	}

	return aqua, fmt.Errorf(message)
}

func authenticate(aqua *Aqua) (bool, string) {

	url := fmt.Sprintf("%s/v1/login", aqua.URL)
	params := `{"id":"` + aqua.ID + `", "password":"` + aqua.Password + `"}`
	resp, body, err := aqua.RestClient.Post(url).Send(params).End()

	if err != nil {
		return false, ""
	}

	if resp.StatusCode == 200 {
		//var raw map[string]interface{}
		_ = json.Unmarshal([]byte(body), &aqua)
		return true, ""
	}

	return false, fmt.Sprintf("Failed with status: %s", resp.Status)

}
