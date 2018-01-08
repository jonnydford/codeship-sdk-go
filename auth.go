package codeship

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// AuthService is an interface for managing authentication with the Codeship V2 API.
type AuthService interface {
	Get(context.Context, *AuthRequest) (*Auth, *Response, error) // POST
}

// AuthRequest respresents a request for authentication.
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse represents a response from the AuthRequest
type AuthResponse struct {
	AccessToken   string             `json:"access_token"`
	Organizations *OrganizationsRoot `json:"organizations"`
	ExpiresAt     string             `json:"expires_at"`
}

// OrganizationsRoot represents the list of organisations from an AuthResponse
type OrganizationsRoot struct {
	orguuid   string  `json:"uuid"`
	orgname   string  `json:"name"`
	orgscopes *Scopes `json:"scopes"`
}

// Scopes represents the list of scopes from the OrganizationRoot
type Scopes struct {
	orgscopes []string `json:"scopes"`
}

func authenticate() {

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("POST", APIEndpointAuth, payload)

	req.Header.Add("content-type", jsonType)
	req.SetBasicAuth(user, password)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
