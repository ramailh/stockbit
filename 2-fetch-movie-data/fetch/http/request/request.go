package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	data []byte
}

func (res *result) Map() map[string]interface{} {
	result := make(map[string]interface{})
	if err := json.Unmarshal(res.data, &result); err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func (res *result) ArrayOfMap() []map[string]interface{} {
	var result []map[string]interface{}
	if err := json.Unmarshal(res.data, &result); err != nil {
		log.Println(err)
		return nil
	}

	return result
}

func Get(url string) (*result, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result{data: respJSON}, nil
}

func Post(url string, data interface{}) (*result, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	respJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &result{data: respJSON}, nil
}
