package gfycat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GFYClientGrant represents the locally stored API OAuth client credentials
type GFYClientGrant struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
}

// GFYAccessToken represents an exchanged OAuth token
type GFYAccessToken struct {
	TokenType string    `json:"token_type"`
	Scope     string    `json:"scope"`
	Expires   int       `json:"expires_in"`
	Token     string    `json:"access_token"`
	Created   time.Time `json:"-"`
}

// Exchange turns local credentials into an active token
func (gfy *GFYClientGrant) Exchange() (*GFYAccessToken, error) {
	body, err := ToJSON(gfy)

	if err != nil {
		return nil, err
	}

	oauth := fmt.Sprintf("%s/%s/token", BaseURI, "oauth")
	resp, err := http.Post(oauth, "application/json", body)

	if err != nil {
		return nil, err
	}

	token := &GFYAccessToken{Created: time.Now()}
	apiresp := APIResponse(*resp)
	err = apiresp.ToType(token)

	return token, err
}

// NewClientGrantFromFile loads a file and returns a GFYClientGrant
func NewClientGrantFromFile(path string) (*GFYClientGrant, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	grant := &GFYClientGrant{}

	err = json.Unmarshal(file, grant)
	if err != nil {
		return nil, err
	}

	return grant, nil
}
