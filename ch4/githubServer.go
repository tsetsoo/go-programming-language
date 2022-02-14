package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"tsvetelinpantev.com/go-programming-language/ch4/github"
)

var issueList = template.Must(template.New("issuelist").Parse(`

<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Milestone</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.Milestone.HTMLURL}}'>{{.Milestone.Number}}:{{.Milestone.Title}}</a></td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		svgFunction(w)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func svgFunction(w io.Writer) {
	result, err := github.ReadIssues()
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
