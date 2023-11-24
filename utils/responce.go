package utils

import (
	"Assessment/consts"
	"Assessment/tapcontext"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type FieldsMap map[string]interface{}

type ErrResponse struct { //to-do: it must be renamed to a generic response struct
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

// ReturnResponse forms the http response in json format
func ReturnResponse(w http.ResponseWriter, statusCode int, status interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	en := json.NewEncoder(w)
	_ = en.Encode(status)
}

// ErrorResponse returns generic error response
func ErrorResponse(ctx tapcontext.TContext, w http.ResponseWriter, responseErrorMessage string, statusCode int, logError error, fields FieldsMap) {
	w.Header().Set("Content-Type", "application/json")
	var buf = new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	_ = encoder.Encode(ErrResponse{Message: responseErrorMessage})
	w.WriteHeader(statusCode)
	_, _ = w.Write(buf.Bytes())
}

func GetHTTP(ctx tapcontext.TContext, url string, response interface{}, headers map[string]string) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf(GetError(consts.ErrGetFailed, ctx.Locale), err.Error())
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("%d %s", resp.StatusCode, string(body))
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return err
	}

	return nil

}
