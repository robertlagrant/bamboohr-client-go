package bamboohr_client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func CallJsonApi(url string, apikey string, method string) ([]byte, error) {
	response, err := CallJsonApiInternal(url, apikey, method)
	if err != nil {
		return nil, fmt.Errorf("1 Got error %s", err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("2 Got error %s", err.Error())
	}

	defer response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("3 Got error %s", err.Error())
	}
	return body, nil
}

func CallJsonApiList(url string, apikey string, method string) ([]map[string]interface{}, error) {
	response, err := CallJsonApiInternal(url, apikey, method)
	if err != nil {
		return nil, fmt.Errorf("1 Got error %s", err.Error())
	}

	body, err := decodeJsonListResponse(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("2 Got error %s", err.Error())
	}
	return body, nil
}

func CallJsonApiMap(url string, apikey string, method string) (map[string]interface{}, error) {
	response, err := CallJsonApiInternal(url, apikey, method)
	if err != nil {
		return nil, fmt.Errorf("3 Got error %s", err.Error())
	}

	body, err := decodeJsonMapResponse(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("4 Got error %s", err.Error())
	}
	return body, nil
}

func CallJsonApiInternal(url string, apikey string, method string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("5 Got error %s", err.Error())
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
	return response, nil
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
