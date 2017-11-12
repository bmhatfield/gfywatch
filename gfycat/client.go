package gfycat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	// BaseURI represents GFYCat's API URL
	BaseURI = "https://api.gfycat.com/v1"

	// OAuthSubURI represents GFYCat's Oauth Subdirectory
	OAuthSubURI = "oauth"
)

func responseToJSON(resp *http.Response, t interface{}) error {
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, t)
	if err != nil {
		return err
	}

	return nil
}

func typeToJSONReader(t interface{}) (io.Reader, error) {
	data, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)

	return reader, nil
}

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
	body, err := typeToJSONReader(gfy)

	if err != nil {
		return nil, err
	}

	oauth := fmt.Sprintf("%s/%s/token", BaseURI, OAuthSubURI)
	resp, err := http.Post(oauth, "application/json", body)

	if err != nil {
		return nil, err
	}

	token := &GFYAccessToken{Created: time.Now()}

	err = responseToJSON(resp, token)

	return token, err
}

// GFYClient represents an authenticated request dispatcher to the GFYCat API
type GFYClient struct {
	Grant  *GFYClientGrant
	Token  *GFYAccessToken
	client *http.Client
}

// request returns an authorized request client
func (gfy *GFYClient) request(method string, URL string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, URL, body)

	if err != nil {
		return nil, err
	}

	tokenType := strings.Title(gfy.Token.TokenType)
	token := fmt.Sprintf("%s %s", tokenType, gfy.Token.Token)

	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := gfy.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		fmt.Printf("Response code was: %d - body was: %s", resp.StatusCode, responseBody)
	}

	return resp, nil
}

// NewGFYClient accepts a grant and returns an authenticated client
func NewGFYClient(grant *GFYClientGrant) (*GFYClient, error) {
	token, err := grant.Exchange()

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	gfy := &GFYClient{Grant: grant, Token: token, client: client}

	return gfy, nil
}
