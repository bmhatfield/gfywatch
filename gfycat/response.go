package gfycat

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// APIResponse represents an HTTP response from the GFYCat API
type APIResponse http.Response

// ToType will unmarshal a GfyCat API response into the provided object
func (gfy *APIResponse) ToType(destination interface{}) error {
	responseBody, err := ioutil.ReadAll(gfy.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, destination)
	if err != nil {
		return err
	}

	return nil
}

// ToJSON will encode an object to it's JSON representation and provide it as a io.Reader
func ToJSON(t interface{}) (io.Reader, error) {
	data, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)

	return reader, nil
}
