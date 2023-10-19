// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
	Age       string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	printAgeCategories(result)

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		duration := time.Since(item.CreatedAt)
		if duration.Hours() < 30*24 {
			item.Age = "less than a month"
		} else if duration.Hours() < 365*24 {
			item.Age = "less than a year"
		} else {
			item.Age = "more than a year"
		}

		fmt.Printf("#%-5d %9.9s %-20.17s %.55s \n",
			item.Number, item.User.Login, item.Age, item.Title)
	}
}

//SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//refer https://github.com/torbiak/gopl/blob/master/ex4.10/main.go
func printAgeCategories(result *IssuesSearchResult) {
	lessAMonth := make([]*Issue, 0)
	lessAYear := make([]*Issue, 0)
	moreAYear := make([]*Issue, 0)
	now := time.Now()

	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	format := "#%-5d %9.9s %.10v\n"

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(month):
			lessAMonth = append(lessAMonth, item)
		case item.CreatedAt.After(year):
			lessAYear = append(lessAYear, item)
		default:
			moreAYear = append(moreAYear, item)
		}
	}

	fmt.Println("This is less a month issue")
	for _, item := range lessAMonth {
		fmt.Printf(format, item.Number, item.Title, item.CreatedAt)
	}

	fmt.Println("This is less a year issue")
	for _, item := range lessAYear {
		fmt.Printf(format, item.Number, item.Title, item.CreatedAt)
	}

	fmt.Println("This is more than a year issue")
	for _, item := range moreAYear {
		fmt.Printf(format, item.Number, item.Title, item.CreatedAt)
	}
}
