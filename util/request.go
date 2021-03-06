package util

import (
	"io/ioutil"
	"net/http"

	"encoding/json"

	"github.com/jiangshengwu/aliyun-sdk-for-go/log"
)

type AliyunRequest struct {
	Url string
}

// Http get request
func (request *AliyunRequest) DoGetRequest() (string, error) {
	resp, err := http.Get(request.Url)
	if err != nil {
		// handle error
		log.Error(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Error(err)
		return "", err
	}
	result := string(body)
	var errResp ErrorResponse
	json.Unmarshal([]byte(result), &errResp)
	if errResp.Message != "" {
		err = &SdkError{errResp, request.Url}
	}
	return result, err
}

// Get formatted query string from map
func GetQueryFromMap(params map[string]string) string {
	query := ""
	for k, v := range params {
		query += "&" + k + "=" + v
	}
	if len(query) > 0 {
		query = query[1:]
	}
	return query
}
