package links

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

func Extract(urlToCheck string) ([]string, error) {
	originalUrl, err := url.Parse(urlToCheck)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(urlToCheck)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", urlToCheck, resp.Status)
	}
	forwardedHost := resp.Request.URL.Host
	if forwardedHost != originalUrl.Host {
		return []string{}, nil
	}
	var fileName string
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	dirPath := currentPath + "/crawled"
	fullPath := resp.Request.URL.Path
	if len(fullPath) > 0 {
		fullPath = removeTrailingSlash(fullPath)
		fullPathSplit := strings.Split(fullPath, "/")
		if len(fullPathSplit) > 1 {
			fileName = fullPathSplit[len(fullPathSplit)-1]
			dirPath += strings.Join(fullPathSplit[:len(fullPathSplit)-1], "/")
			fmt.Println(dirPath)
			createDir(dirPath)
		} else {
			fileName = fullPath
		}
		if !strings.HasSuffix(fileName, ".html") {
			fileName = fileName + ".html"
		}
		dirPath += "/"
	} else {
		createDir(dirPath)
		fileName = "/home.html"
	}

	doc, err := html.Parse(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", urlToCheck, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if link.Host != originalUrl.Host {
					continue // skip links outside of domain
				}
				linkString := link.String()
				linkParts := strings.Split(linkString, "#")
				if len(linkParts) > 1 {
					linkString = strings.Join(linkParts[:len(linkParts)-1], "#") // remove anchor tag
				}
				linkString = removeTrailingSlash(linkString)
				fileExtension := filepath.Ext(linkString)
				if len(fileExtension) > 0 && fileExtension != ".html" {
					continue //skip downloadable links
				}
				links = append(links, linkString)
			}
		}
	}
	f, err := os.Create(dirPath + fileName)
	if err != nil {
		return nil, fmt.Errorf("creating file at %s: %v", dirPath+fileName, err)
	}
	defer f.Close()
	fmt.Println("path is ", dirPath+fileName)
	var htmlPageBuffer bytes.Buffer
	err = html.Render(&htmlPageBuffer, doc)
	if err != nil {
		return nil, fmt.Errorf("parsing html: %v", err)
	}
	stringHtm := htmlPageBuffer.String()
	// fmt.Println(stringHtm)
	_, err = f.WriteString(stringHtm)
	if err != nil {
		return nil, fmt.Errorf("writing to file: %v", err)
	}
	f.Sync()
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func removeTrailingSlash(withPossibleTrailingSlash string) string {
	j := len(withPossibleTrailingSlash)
	_, size := utf8.DecodeLastRuneInString(withPossibleTrailingSlash[:j])
	j -= size

	lastByRune := withPossibleTrailingSlash[j:]
	if lastByRune == "/" {
		withPossibleTrailingSlash = withPossibleTrailingSlash[:j]
	}
	return withPossibleTrailingSlash
}

func createDir(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
