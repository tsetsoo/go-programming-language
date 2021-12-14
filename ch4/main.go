// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"local/ch4/github"
// )

// func searchTask() {
// 	result, err := github.SearchIssues(os.Args[1:])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%d issues:\n", result.TotalCount)
// 	lessThanMonthOld := []*github.Issue{}
// 	monthAgo := time.Now().AddDate(0, -1, 0)
// 	lessThanYearOld := []*github.Issue{}
// 	yearAgo := time.Now().AddDate(-1, 0, 0)
// 	moreThanYearOld := []*github.Issue{}
// 	for _, item := range result.Items {
// 		fmt.Println(item.CreatedAt.Date())
// 		if item.CreatedAt.After(monthAgo) {
// 			lessThanMonthOld = append(lessThanMonthOld, item)
// 		} else if item.CreatedAt.After(yearAgo) {
// 			lessThanYearOld = append(lessThanYearOld, item)
// 		} else {
// 			moreThanYearOld = append(moreThanYearOld, item)
// 		}
// 	}
// 	fmt.Println("Less than a month old")
// 	for _, item := range lessThanMonthOld {
// 		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
// 	}
// 	fmt.Println("Less than a year old")
// 	for _, item := range lessThanYearOld {
// 		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
// 	}
// 	fmt.Println("More than a year old")
// 	for _, item := range moreThanYearOld {
// 		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
// 	}
// }

// func main() {
// 	// result, err := github.ReadIssues()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// b, err := json.MarshalIndent(*result, "", "  ")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(string(b))
// 	issueBody, err := github.CaptureInputFromEditor()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	issue := github.Issue{
// 		Title: "Test issue",
// 		Body:  string(issueBody),
// 	}
// 	err = github.CreateIssue(&issue)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
package main

import (
	"fmt"
	"testing"
)

func benchmarkBool(b *testing.B) {
	s := make(map[int]bool)

	for i := 0; i < 10000*b.N; i++ {
		s[2*i] = false
	}
}

func benchmarkStruct(b *testing.B) {
	s := make(map[int]struct{})

	for i := 0; i < 10000*b.N; i++ {
		s[2*i] = struct{}{}
	}
}

func main() {
	boolRes := testing.Benchmark(benchmarkBool)
	fmt.Println("bool:", boolRes.MemString())

	structRes := testing.Benchmark(benchmarkStruct)
	fmt.Println("struct{}:", structRes.MemString())

	fmt.Println("ratio:", float32(boolRes.AllocedBytesPerOp())/float32(structRes.AllocedBytesPerOp()))
}
