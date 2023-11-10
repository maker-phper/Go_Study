package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type issue struct {
	URL           string
	RepositoryUrl string `json:"repository_url"`
	ID            int
	Number        int
	Title         string
	User          *user
	State         string
	Locked        bool
	CreatedAt     time.Time `json:"created_at"`
	Body          string
}

type user struct {
	Login     string
	UserId    int    `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	UserType  string `json:"type"`
}

var IssueIds = make(map[int]int)

const Accept = "application/vnd.github+json"
const Authorization = "Bearer ghp_DYA7u23EfKS"
const ApiVersion = "2022-11-28"
const IssuesUrl = "https://api.github.com/repos/maker-phper/Go_Study/issues"

var action = flag.String("a", "list", "list")
var title = flag.String("t", " ", "")
var message = flag.String("m", " ", "")
var number = flag.Int("n", 0, "")

func main() {
	flag.Parse()
	switch *action {
	case "list":
		readIssues()
	case "create":
		if *title == "" || *message == "" {
			fmt.Println("title and message can't empty")
			break
		}
		//createIssues()

	case "update":
		if _, ok := IssueIds[*number]; !ok {
			fmt.Println("please enter a valid number")
			break
		}
		if *title == "" || *message == "" {
			fmt.Println("title and message can't empty")
			break
		}
	default:
		fmt.Println("please enter -h for more information")
	}

}

func returnIssuesJsonRes(body []byte) (*[]issue, error) {
	var result []issue
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func readIssues() {
	body, err := httpRequest(IssuesUrl, "GET", nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	var issuesMap *[]issue
	issuesMap, err = returnIssuesJsonRes(body)
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
	fmt.Printf("%-s\t %-s\t %-s\t %s\n", "Number", "Title", "Content", "Create_at")
	for index, items := range *issuesMap {
		IssueIds[index] = items.Number
		fmt.Printf("%.3d\t %s\t %.55s\t %.19v\n", items.Number, items.Title, items.Body, items.CreatedAt)
	}
}

func httpRequest(reqUrl, method string, form map[string]string) ([]byte, error) {
	client := http.Client{}
	jsonData, err := json.Marshal(form)
	req, err := http.NewRequest(method, reqUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	//req.Header = header
	req.Header = http.Header{
		"Accept":               {Accept},
		"Authorization":        {Authorization},
		"X-GitHub-Api-Version": {ApiVersion},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue error %s", resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}

func createIssues() {
	params := make(map[string]string, 5)
	params["title"] = "testIssues2"
	params["body"] = "create second new Issues"
	form := url.Values{}
	for k, v := range params {
		form.Add(k, v)
	}
	http, err := httpRequest(IssuesUrl, "POST", params)
	if err != nil {
		log.Print(err)
		os.Exit(2)
	}
	fmt.Println(http)
}

func updateIssues() {
	updateUrl := IssuesUrl + "/" + strconv.Itoa(IssueIds[1])
	params := make(map[string]string, 5)
	params["title"] = "create first issues"
	params["body"] = "add new issues body"
	http, err := httpRequest(updateUrl, "PATCH", params)
	if err != nil {
		log.Print(err)
		os.Exit(3)
	}
	fmt.Println(string(http))
}

func deleteIssues() {
	updateUrl := IssuesUrl + "/" + strconv.Itoa(IssueIds[0])
	http, err := httpRequest(updateUrl, "DELETE", nil)
	if err != nil {
		log.Print(err)
		os.Exit(4)
	}
	fmt.Println(string(http))
}

func openTextEditor() {
	//api don't support delete issue
}
