package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

var (
	githubRawHost = "https://raw.githubusercontent.com"
	githubAPIHost = "https://api.github.com"
)

type CaseLabs struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CaseData struct {
	Slug   string   `json:"slug"`
	Spec   string   `json:"spec"`
	Repo   string   `json:"repo"`
	Start  []string `json:"start"`
	Vars   Vars     `json:"vars"`
	Grade  []Grade  `json:"grade"`
	Finish []string `json:"finish"`
}
type Vars map[string]interface{}

type Grade struct {
	Name   string `json:"name"`
	On     string `json:"on"`
	Script string `json:"script"`
	Expect string `json:"expect"`
}

func getBody(url string) (body []byte, httpCode int) {
	// set context with timeout to prevent goroutine leak
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	// request with context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// setup http client
	client := getHTTPClient()

	// do request with setup
	resp, err := client.Do(req)

	// error raise if timeout
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// read all body it should be json format
	// error if cant parse to json
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data, resp.StatusCode
}

func FileCaseParser(file []byte) (result CaseData, err error) {
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// GetCase GitHub file with specific url format
// https://raw.githubusercontent.com /<repo>/<branch>/<slug>/case.yaml
func GetCase(repo string, branch string, slug string) (result CaseData, httpCode int) {
	course := strings.Split(slug, "-")[0]
	repos := strings.Split(repo, "/")

	pathUrl := path.Join(repos[1], repos[2], branch, course, slug, "case.yaml")
	url := githubRawHost + "/" + pathUrl

	data, httpCode := getBody(url)

	err := yaml.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return result, httpCode
}

// GetReadme on specific slug will return byte array
// name instruction file must be README.md
func GetReadme(repo string, branch string, slug string) (readme []byte, httpCode int) {
	course := strings.Split(slug, "-")[0]

	pathUrl := path.Join(repo, branch, course, slug, "README.md")
	url := githubRawHost + "/" + pathUrl

	readme, httpCode = getBody(url)

	return
}

// GetListLab return array of string list available lab
// request format  https://api.github.com/repos/Kyuubang/philo-sample-case/contents/linux
func GetListLab(repo string, course string) (labList []string, httpCode int) {
	var labs []CaseLabs

	pathUrl := path.Join("repos", repo, "contents", course)
	url := githubAPIHost + "/" + pathUrl

	data, httpCode := getBody(url)

	if httpCode != 200 {
		return
	}

	err := json.Unmarshal(data, &labs)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, lab := range labs {
		if lab.Type == "dir" {
			labList = append(labList, lab.Name)
		}
	}
	return
}

// TODO: create function download lab asset
// DownloadDirContent download recursive content
// on specific directory
//func DownloadDirContent() {}

// GetRepoContent GitHub on specific folder return interface
// will be return level 1 directory list
//func getRepoContent(repo string, filePath string) (result map[string]interface{}) {
//	url := path.Join(githubAPIHost, "repos", repo, "contents", filePath)
//
//	data := getBody(url)
//
//	err := json.Unmarshal(data, &result)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	return result
//}
