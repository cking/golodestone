package golodestone

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

var reURLPath = regexp.MustCompile(`^/?(lodestone/)?(.*)`)

// BuildURL Create the full Lodestone URL
func BuildURL(path string) (string, error) {
	// make sure we have a valid
	urlParts, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	urlParts.Scheme = "http"
	urlParts.Host = "eu.finalfantasyxiv.com"
	urlParts.Path = reURLPath.ReplaceAllString(urlParts.Path, `/lodestone/$2`)
	return urlParts.String(), nil
}

// QueryLodestone Query the Final Fantasy XIV Lodestone for Information
func QueryLodestone(urlPath string) (*html.Node, error) {
	// make sure we have a valid
	url, err := BuildURL(urlPath)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	node, ok := scrape.Find(root, scrape.ById("main"))
	if !ok {
		return nil, errors.New("Couldn't find main node")
	}
	return node, nil
}
