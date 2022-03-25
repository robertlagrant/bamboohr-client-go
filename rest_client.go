package bamboohr_client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func CallJsonApi(url string, apikey string, method string) ([]byte, error) {
	req, _ := http.NewRequest(method, url, nil)
	body, err := CallJsonApiInternal(req, apikey)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}

	return body, nil
}

func CallJsonApiWithPayload(url string, apikey string, method string, payload *strings.Reader) ([]byte, error) {
	req, _ := http.NewRequest(method, url, payload)
	body, err := CallJsonApiInternal(req, apikey)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}

	return body, nil
}

func CallJsonApiInternal(req *http.Request, apikey string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apikey, ".")
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("6 Got error %s", err.Error())
	}
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("7 Got error HTTP response %s", response.Status)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("2 Got error %s", err.Error())
	}

	return body, nil
}

func decodeJsonMapResponse(reader io.Reader) (map[string]interface{}, error) {
	var m map[string]interface{}

	return m, json.NewDecoder(reader).Decode(&m)
}

func decodeJsonListResponse(reader io.Reader) ([]map[string]interface{}, error) {
	var l []map[string]interface{}

	return l, json.NewDecoder(reader).Decode(&l)
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
