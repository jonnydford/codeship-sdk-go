package codeship

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ProjectsService is an interface for managing projects with the Codeship V2 API.
type ProjectsService interface {
	List(context.Context, *ListOptions) ([]Domain, *Response, error)                                 // GET
	Get(context.Context, string) (*Domain, *Response, error)                                         // GET
	Create(context.Context, *DomainCreateRequest) (*Domain, *Response, error)                        // POST
	Update(context.Context, string, int, *DomainRecordEditRequest) (*DomainRecord, *Response, error) // PUT
}

// Builds respresents builds
type Builds struct {
	Repourl              string   `json:"repository_url"` // "https://github.com/codeship/documentation.git
	Buildtype            string   `json:"type"`           // "basic"
	Setupcommands        []string `json:"setup_commands"` // [ "nvm install 6" ]
	Teamids              int      `json:"team_ids"`       // [ 99173438, 64874098, 74618189, 92403869 ]
	Testpipelines        *TestPipelines
	Noticationrules      *NotificationRules
	Environmentvariables *EnvironmentVariables
	Deploymentpipelines  *DeploymentPipelines
}

// TestPipelines represents your builds' test pipelines
type TestPipelines struct {
	Name     string   `json:"name"`     // "Tests"
	Commands []string `json:"commands"` // [ "npm test" ]
}

// NotificationRules represents your builds' notification rules
type NotificationRules struct {
	Notifier      string `json:"notifier"` // "webhook"
	Target        string `json:"target"`
	Branch        string `json:"branch"` // "development"
	Options       *Options
	BuildStatuses []string `json:"build_statuses"`
	BranchMatch   string   `json:"branch_match"`
}

// Options represents your Builds' Notification options
type Options struct {
	Key  string `json:"key"`
	URL  string `json:"url"` // "https://localhost/webhook"
	Room string `json:"room"`
}

// EnvironmentalVariables represents your builds' environmental variables
type EnvironmentalVariables struct {
	Name  string `json:"name"`  // "env_name_1"
	Value string `json:"value"` // "foo"
}

// DeploymentPipelines represents your builds' deployment pipelines
type DeploymentPipelines struct {
	Branch   *Branch
	Config   []string `json:"config"`   // [ "./deploy_master.sh" ]
	Position int      `json:"position"` // 1
}

// Branch represents your builds' deployment pipelines' branch
type Branch struct {
	BranchName string `json:"branch_name"` // "master"
	MatchMode  string `json:"match_mode"`  // "exact"
}

func createproject() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects"

	payload := strings.NewReader("{\"repository_url\":\"https://github.com/codeship/documentation.git\",\"type\":\"basic\",\"team_ids\":[99173438,64874098,74618189,92403869],\"notification_rules\":[{\"notifier\":\"webhook\",\"branch\":\"development\",\"options\":{\"url\":\"https://localhost/webhook\"}}],\"deployment_pipelines\":[{\"branch\":{\"branch_name\":\"master\",\"match_mode\":\"exact\"},\"config\":[\"./deploy_master.sh\"],\"postion\":1}],\"setup_commands\":[\"nvm install 6\"],\"environment_variables\":[{\"name\":\"env_name_1\",\"value\":\"foo\"},{\"name\":\"env_name_2\",\"value\":\"bar\"}],\"test_pipelines\":[{\"name\":\"Tests\",\"commands\":[\"npm test\"]}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func getproject() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + builduuid

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func listprojects() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects"

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func updateproject() {

	url := defaultBaseURL + "/organizations/" + orguuid + "/projects/" + builduuid

	payload := strings.NewReader("{\"type\":\"string (required)\",\"setup_commands\":[\"string\"],\"team_ids\":[\"number\"],\"notification_rules\":[{\"notifier\":\"string (optional)\",\"email_target\":\"string (optional)\",\"branch\":\"string (optional)\",\"options\":{\"slack\":{\"key\":\"string (optional)\"},\"webhook\":{\"url\":\"string (optional)\"},\"hipchat\":{\"key\":\"string (optional)\",\"room\":\"string (optional)\"},\"campfire\":{\"key\":\"string (optional)\",\"room\":\"string (optional)\",\"domain\":\"string (optional)\"},\"grove\":{\"key\":\"string (optional)\"},\"flowdock\":{\"key\":\"string (optional)\"}}}],\"environment_variables\":[{\"name\":\"string (optional)\",\"value\":\"string (optional)\"}]}")

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
