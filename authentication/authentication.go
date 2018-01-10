package authentication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// AuthService is an interface for managing authentication with the Codeship V2 API.
//type AuthService interface {
//	Get(context.Context, *AuthRequest) (*Auth, *Response, error) // POST
//}

// // AuthRequest respresents a request for authentication.
// type AuthRequest struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// AuthResponse details the response from the Authenticate function
type authresponse struct {
	ExpiresAt     int    `json:"expires_at"`
	AccessToken   string `json:"access_token"`
	Organizations []struct {
		UUID   string   `json:"uuid"`
		Name   string   `json:"name"`
		Scopes []string `json:"scopes"`
	} `json:"organizations"`
}

// APIEndpoint constants
const (
	APIEndpointBase = "https://api.codeship.com/v2"
	APIEndpointAuth = "/auth"
)

// Authenticate takes a username and password, gets basic auth and returns AuthResponse
func Authenticate(user string, password string) {
	payload := strings.NewReader("{}")
	req, _ := http.NewRequest("POST", APIEndpointBase+APIEndpointAuth, payload)
	req.Header.Add("content-type", "application/json")
	req.SetBasicAuth(user, password)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var response authresponse
	json.Unmarshal([]byte(body), &response)
	Accesstoken := response.AccessToken
	fmt.Println(Accesstoken)
	return
}
