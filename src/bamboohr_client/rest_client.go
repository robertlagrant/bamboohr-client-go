package bamboohr_client

import (
	"fmt"
	"io"
	"encoding/json"
    "net/http"
	"time"
)

func CallBamboohrJson(url, apikey string, method string) (map[string]interface{}, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apikey, ".")
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("Got error HTTP response %s", response.Status)
	}
	body, err := decodeJsonMapResponse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	return body, nil
}

func decodeJsonMapResponse(reader io.Reader) (map[string]interface{}, error) {
	var m map[string]interface{}
 
	return m, json.NewDecoder(reader).Decode(&m)
 }