package codeship

import (
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
