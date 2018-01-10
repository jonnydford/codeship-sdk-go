package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Builds represents
type Builds struct {
	Ref       string `json:"ref"`        // "heads/master"
	CommitSha string `json:"commit_sha"` // "5927b8c40deecc656138f61b7e0d77e9cca835ac"
}

func createbuild(orguuid string, projectuuid string) {
	payload := strings.NewReader("{\"ref\":\"heads/master\",\"commit_sha\":\"5927b8c40deecc656138f61b7e0d77e9cca835ac\"}")
	url, _ := urljoin(APIEndpointBase, APIEndpointCreateBuild)
	httppost(ctx, url, payload)
}

func getbuild(orguuid string, projectuuid string, builduuid string) {
	url, _ := urljoin(APIEndpointBase, APIEndpointGetBuild)
	httpget(ctx, url)

	req, _ := http.NewRequest("GET", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}

func listbuilds(orguuid string, projectuuid string) {
	url, _ := urljoin(APIEndpointBase, APIEndpointListBuilds)
	httpget(ctx, url)
}

func getbuildpipelines(orguuid string, projectuuid string, builduuid string) {
	url, _ := urljoin(APIEndpointBase, APIEndpointGetBuildPipelines)
	httpget(ctx, url)
}

func stopbuild(orguuid string, projectuuid string, builduuid string) {
	url, _ := urljoin(APIEndpointBase, APIEndpointStopBuild)
	httppost(ctx, url)
}

func restartbuild(orguuid string, projectuuid string, builduuid string) {
	url, _ := urljoin(APIEndpointBase, APIEndpointRestartBuild)
	httppost(ctx, url)
}

func getbuildservices(orguuid string, projectuuid string, builduuid string) {
	url, _ := urljoin(APIEndpointBase, APIEndpointGetBuildServices)
	httpget(ctx, url)
}

func getbuildsteps(orguuid string, projectuuid string, builduuid string) {
	url, _ := urljoin(APIEndpointBase, APIEndpointGetBuildSteps)
	httpget(ctx, url)
}
