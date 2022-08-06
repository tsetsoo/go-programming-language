package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

var issueList = template.Must(template.New("issuelist").Parse(`

<table>
<tr style='text-align: left'>
  <th>Item</th>
  <th>Price</th>
</tr>
{{ range $key, $value := . }}
<tr>
  <td>{{$key}}</td>
  <td>{{$value}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/create", db.create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := issueList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price not a number: %q\n", req.URL.Query().Get("price"))
		return
	}
	mu.Lock()
	defer mu.Unlock()
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}
	db[item] = dollars(price)
	w.WriteHeader(http.StatusCreated) // 201
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, err := strconv.ParseFloat(req.URL.Query().Get("price"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "price not a number: %q\n", req.URL.Query().Get("price"))
		return
	}
	mu.Lock()
	defer mu.Unlock()
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item does not exist: %q\n", item)
		return
	}
	db[item] = dollars(price)
	w.WriteHeader(http.StatusOK) // 200
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	defer mu.Unlock()
	delete(db, item)
	w.WriteHeader(http.StatusNoContent) // 204
}
