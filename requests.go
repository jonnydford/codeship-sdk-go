package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

func urljoin(APIEndpointBase string, APIEndpointPath string) string {
	u := path.Join(APIEndpointBase, APIEndpointPath)
	return u.String()
}

func (client *Client) httpget(ctx context.Context, url string) {
	payload := strings.NewReader("{}")
	req, _ := http.NewRequest("GET", url)
	req.Header.Set("Authorization", "Bearer "+client.pat)
	req.Header.Add("content-type", jsonType)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(res)
	// fmt.Println(string(body))
}

func (client *Client) httpput(ctx context.Context, url string, payload string) {
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("content-type", jsonType)
	req.Header.Set("Authorization", "Bearer "+client.pat)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(res)
	// fmt.Println(string(body))
}

func (client *Client) httppost(ctx context.Context, url string, payload string) {
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Set("Authorization", "Bearer "+client.pat)
	req.Header.Add("content-type", jsonType)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(res)
	// fmt.Println(string(body))
}
