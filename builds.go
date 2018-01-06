package codeship

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// BuildsService is an interface for managing builds with the Codeship V2 API.
type BuildsService interface {
	List(context.Context, *ListOptions) ([]Domain, *Response, error)                                 // GET
	Get(context.Context, string) (*Domain, *Response, error)                                         // GET
	Create(context.Context, *DomainCreateRequest) (*Domain, *Response, error)                        // POST
	Update(context.Context, string, int, *DomainRecordEditRequest) (*DomainRecord, *Response, error) // PUT
}

// Builds represents
type Builds struct {
	Ref       string `json:"ref"`        // "heads/master"
	CommitSha string `json:"commit_sha"` // "5927b8c40deecc656138f61b7e0d77e9cca835ac"
}

func createbuild() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + projectuuid + "/builds"

	payload := strings.NewReader("{\"ref\":\"heads/master\",\"commit_sha\":\"5927b8c40deecc656138f61b7e0d77e9cca835ac\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func getbuild() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + projectuuid + "/builds/" + builduuid

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func listbuilds() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + projectuuid + "/builds"

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func getbuildpipelines() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + projectuuid + "/builds/" + builduuid + "/pipelines"
	"https://psfmbbh2bj3yymna6-mock.stoplight-proxy.io/v2/organizations/%7Borganization_uuid%7D/projects/%7Bproject_uuid%7D/builds/%7Bbuild_uuid%7D/pipelines"

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func stopbuild() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + projectuuid + "/builds/" + builduuid + "/stop"

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("POST", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func restartbuild() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + projectuuid + "/builds/" + builduuid + "/restart"

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("POST", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func getbuildservices() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + projectuuid + "/builds/" + builduuid + "/services"

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func getbuildsteps() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + projectuuid + "/builds/" + builduuid + "/steps"

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
