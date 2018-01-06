package codeship

import (
	"net/http"
	"net/url"
)

// APIEndpoint constants
const (
	APIEndpointBase              = "https://api.codeship.com/v2"
	APIEndpointAuth              = "auth"
	APIEndpointCreateBuild       = "/organizations/%s/projects/%s/builds"              // orguuid / projectuuid
	APIEndpointGetBuild          = "/organizations/%s/projects/%s/builds/%s"           // orguuid / projectuuid / builduuid
	APIEndpointListBuilds        = "/organizations/%s/projects/%s/builds"              // orguuid / projectuuid
	APIEndpointGetBuildPipelines = "/organizations/%s/projects/%s/builds/%s/pipelines" // orguuid / projectuuid / builduuid
	APIEndpointStopBuild         = "/organizations/%s/projects/%s/builds/%s/stop"      // orguuid / projectuuid / builduuid
	APIEndpointRestartBuild      = "/organizations/%s/projects/%s/builds/%s/restart"   // orguuid / projectuuid / builduuid
	APIEndpointGetBuildServices  = "/organizations/%s/projects/%s/builds/%s/services"  // orguuid / projectuuid / builduuid
	APIEndpointGetBuildSteps     = "/organizations/%s/projects/%s/builds/%s/steps"     // orguuid / projectuuid / builduuid
	APIEndpointCreateProject     = "/organizations/%s/projects"                        // orguuid
	APIEndpointGetProject        = "/organizations/%s/projects/%s"                     // orguuid / projectuuid
	APIEndpointListProjects      = "/organizations/%s/projects"                        // orguuid
	APIEndpointUpdateProject     = "/organizations/%s/projects/%s"                     // orguuid / projectuuid
)

// Header Types constants
const (
	jsonType  = "application/json"
	plainType = "text/plain"
)

// Client type
type Client struct {
	channelSecret string
	channelToken  string
	endpointBase  *url.URL     // default APIEndpointBase
	httpClient    *http.Client // default http.DefaultClient
}
