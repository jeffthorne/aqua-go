package aqua

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/parnurzeal/gorequest"
)

// Aqua is the main interface used to interact with Aqua's Enterprise platform API endpoint.
type Aqua struct {
	Host               string
	Port               int
	User               string
	Password           string
	RestClient         http.Client
	URL                string
	Secure             bool
	Token              string
	InsecureSkipVerify bool
}

// NewClient - is used to obtain functioning API endpoint
func NewClient(host string, port int, user string, password string, secureEndpoint ...bool) (Aqua, error) {

	aqua := Aqua{Host: host, Port: port, User: user, Password: password, Secure: true, InsecureSkipVerify: true}

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
		return aqua, nil
	}

	return aqua, fmt.Errorf(message)
}

func authenticate(aqua *Aqua) (bool, string) {

	url := fmt.Sprintf("%s/v1/login", aqua.URL)
	params := `{"id":"` + aqua.User + `", "password":"` + aqua.Password + `"}`
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: aqua.InsecureSkipVerify})

	resp, body, err := request.Post(url).Param("abilities", "1").Send(params).End()

	if err != nil {
		return false, ""
	}

	if resp.StatusCode == 200 {
		var raw map[string]interface{}
		_ = json.Unmarshal([]byte(body), &raw)
		aqua.Token = raw["token"].(string)
		return true, ""
	}

	return false, fmt.Sprintf("Failed with status: %s", resp.Status)

}
