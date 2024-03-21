package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	errorUtil "github.com/thhuang/kakao/util/error"
	"github.com/thhuang/kakao/util/keyword"
)

type QueryParam struct {
	Key   string
	Value string
}

// MockHTTPRequest creates a mock HTTP request with the specified method, path, body, and optional query parameters.
// It supports both string and JSON bodies.
func MockHTTPRequest(method string, path string, body interface{}, queryParams ...QueryParam) *http.Request {
	if !strings.HasPrefix(path, "/") {
		panic(fmt.Sprintf("invalid path: %v", path))
	}

	reqBodyReader := bytes.NewBufferString("")
	if body != nil {
		switch v := body.(type) {
		case string:
			reqBodyReader = bytes.NewBufferString(v)
		default:
			reqBodyReader = new(bytes.Buffer)
			if err := json.NewEncoder(reqBodyReader).Encode(body); err != nil {
				panic(err)
			}
		}
	}
	req, _ := http.NewRequest(method, path, reqBodyReader)
	req.Header.Add("Content-Type", "application/json; charset=uft-8")

	if len(queryParams) != 0 {
		q := req.URL.Query()
		for _, p := range queryParams {
			q.Add(p.Key, p.Value)
		}
		req.URL.RawQuery = q.Encode()
	}

	return req
}

// DecodeResBody reads and decodes the JSON body of an HTTP response into a map.
func DecodeResBody(res *http.Response) map[string]interface{} {
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	resBody := map[string]interface{}{}
	if err := json.Unmarshal(body, &resBody); err != nil {
		panic(err)
	}
	return resBody
}

// GetResErrorCode extracts an error code from the response body.
func GetResErrorCode(resBody map[string]interface{}) errorUtil.ErrorCode {
	code, ok := resBody[keyword.Code]
	if !ok {
		panic(errors.New("error code not found"))
	}
	return errorUtil.ErrorCode(int(code.(float64)))
}
