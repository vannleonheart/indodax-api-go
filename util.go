package indodax

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func sendRequest(endpoint string, data *map[string]interface{}, headers *map[string]string) (*[]byte, error) {
	var reqBody *strings.Reader

	httpClient := http.Client{}

	if data != nil {
		mb := mapBody(*data)

		reqBody = strings.NewReader(mb.Encode())
	}

	httpRequest, err := http.NewRequest(http.MethodPost, endpoint, reqBody)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, value := range *headers {
			httpRequest.Header.Set(key, value)
		}
	}

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}

func mapBody(data map[string]interface{}) *url.Values {
	uv := url.Values{}

	for k, v := range data {
		uv.Add(k, fmt.Sprintf("%v", v))
	}

	return &uv
}

func generateSign(data map[string]interface{}, key string) string {
	uv := mapBody(data)

	h := hmac.New(sha512.New, []byte(key))

	h.Write([]byte(uv.Encode()))

	return hex.EncodeToString(h.Sum(nil))
}
