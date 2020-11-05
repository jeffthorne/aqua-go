package aqua

import (
	"encoding/json"
	"fmt"
)

type Registry struct {
	Name string 			`json:"name"`
	Type string 			`json:"type"`
	URL string 				`json:"url"`
	Username string 		`json:"username"`
	Password string 		`json:password"`
	AutoPull bool 			`json:"auto_pull"`
	AutoPullMax int64		`json:"auto_pull_max"`
	AutoPullTime string 	`json:"auto_pull_time"`
	Prefixes []string 		`json:"prefixes"`
	Description string 		`json:"description"`
}
func (aqua *Aqua) CreateRegistry(name, description, registryType, url, username, password string, prefixes []string, auto_pull bool, auto_pull_max int64, autoPullTime string) error{

	registry := Registry{Name: name, Description: description, Type: registryType, URL: url, Username: username, Password: password,
		 				 AutoPull: auto_pull, AutoPullMax: auto_pull_max, AutoPullTime: autoPullTime, Prefixes: prefixes}

	jsonBytes, err := json.Marshal(&registry)
	if err != nil{
		return err
	}

	resp, body, errors := api(aqua, "registries", string(jsonBytes), "v1", "POST")

	if errors != nil{
		return fmt.Errorf(fmt.Sprint(errors))
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204{
		//_ = json.Unmarshal([]byte(body), &results)
		return fmt.Errorf(body)

	}else if resp.StatusCode == 204 {
		fmt.Println("Registry successfully created")
	}



	return nil

}

func (aqua *Aqua) UpdateRegistry(name, description, registryType, url, username, password string, prefixes []string, auto_pull bool, auto_pull_max int64, autoPullTime string) error{

	registry := Registry{Name: name, Description: description, Type: registryType, URL: url, Username: username, Password: password,
		AutoPull: auto_pull, AutoPullMax: auto_pull_max, AutoPullTime: autoPullTime, Prefixes: prefixes}

	jsonBytes, err := json.Marshal(&registry)
	if err != nil{
		return err
	}

	resp, body, errors := api(aqua, fmt.Sprintf("registries/%s", name), string(jsonBytes), "v1", "PUT")

	if errors != nil{
		return fmt.Errorf(fmt.Sprint(errors))
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204{
		//_ = json.Unmarshal([]byte(body), &results)
		return fmt.Errorf(body)

	}else if resp.StatusCode == 204 {
		fmt.Println("Registry successfully updated")
	}



	return nil

}


func (aqua *Aqua) GetRegistry(name string) (string, error) {
	resp, body, errors := api(aqua, fmt.Sprintf("registries/%s", name), "", "v1", "GET")

	if errors != nil{
		return "", fmt.Errorf(fmt.Sprint(errors))
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204{
		return "", fmt.Errorf(body)
	}

		if resp.StatusCode == 200 {
			return body, nil
		}

	return "", fmt.Errorf("Failed with status: %s", resp.Status)
}



