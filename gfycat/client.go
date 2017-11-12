package gfycat

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	// BaseURI represents GFYCat's API URL
	BaseURI = "https://api.gfycat.com/v1"
)

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
