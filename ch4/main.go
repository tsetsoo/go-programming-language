package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"tsvetelinpantev.com/go-programming-language/ch4/github"
)

func searchTask() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	lessThanMonthOld := []*github.Issue{}
	monthAgo := time.Now().AddDate(0, -1, 0)
	lessThanYearOld := []*github.Issue{}
	yearAgo := time.Now().AddDate(-1, 0, 0)
	moreThanYearOld := []*github.Issue{}
	for _, item := range result.Items {
		fmt.Println(item.CreatedAt.Date())
		if item.CreatedAt.After(monthAgo) {
			lessThanMonthOld = append(lessThanMonthOld, item)
		} else if item.CreatedAt.After(yearAgo) {
			lessThanYearOld = append(lessThanYearOld, item)
		} else {
			moreThanYearOld = append(moreThanYearOld, item)
		}
	}
	fmt.Println("Less than a month old")
	for _, item := range lessThanMonthOld {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("Less than a year old")
	for _, item := range lessThanYearOld {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("More than a year old")
	for _, item := range moreThanYearOld {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanLines)
	fmt.Println("What do you want to do? (read/create/update/close)")
	for input.Scan() {
		switch input.Text() {
		case "read":
			result, err := github.ReadIssues()
			if err != nil {
				log.Fatal(err)
			}
			b, err := json.MarshalIndent(*result, "", "  ")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		case "create":
			fmt.Println("Issue title: ")
			input.Scan()
			issueTitle := input.Text()
			issueBody, err := github.CaptureInputFromEditor()
			if err != nil {
				log.Fatal(err)
			}
			issue := github.Issue{
				Title: issueTitle,
				Body:  string(issueBody),
			}
			err = github.CreateIssue(&issue)
			if err != nil {
				log.Fatal(err)
			}
		case "update":
			fmt.Println("Issue number: ")
			input.Scan()
			issueNumber, err := strconv.Atoi(input.Text())
			if err != nil {
				log.Fatal(err)
			}
			issue, err := github.ReadIssue(issueNumber)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("New issue title: (current is %s)\n", issue.Title)
			input.Scan()
			issueTitle := input.Text()
			issueBody, err := github.CaptureInputFromEditor()
			if err != nil {
				log.Fatal(err)
			}
			issue.Title = issueTitle
			issue.Body = string(issueBody)

			err = github.UpdateIssue(issueNumber, issue)
			if err != nil {
				log.Fatal(err)
			}
		case "close":
			fmt.Println("Issue number: ")
			input.Scan()
			issueNumber, err := strconv.Atoi(input.Text())
			if err != nil {
				log.Fatal(err)
			}
			issue, err := github.ReadIssue(issueNumber)
			if err != nil {
				log.Fatal(err)
			}
			issue.State = "closed"
			err = github.UpdateIssue(issueNumber, issue)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("What do you want to do? (read/create/update/close)")
	}

}

// package main

// import (
// 	"fmt"
// 	"testing"
// )

// func benchmarkBool(b *testing.B) {
// 	s := make(map[int]bool)

// 	for i := 0; i < 10000*b.N; i++ {
// 		s[2*i] = false
// 	}
// }

// func benchmarkStruct(b *testing.B) {
// 	s := make(map[int]struct{})

// 	for i := 0; i < 10000*b.N; i++ {
// 		s[2*i] = struct{}{}
// 	}
// }

// func main() {
// 	boolRes := testing.Benchmark(benchmarkBool)
// 	fmt.Println("bool:", boolRes.MemString())

// 	structRes := testing.Benchmark(benchmarkStruct)
// 	fmt.Println("struct{}:", structRes.MemString())

// 	fmt.Println("ratio:", float32(boolRes.AllocedBytesPerOp())/float32(structRes.AllocedBytesPerOp()))
// }
