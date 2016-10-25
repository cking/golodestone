package golodestone

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

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

	if urlParts.Host != "" {
		if !strings.HasSuffix(urlParts.Host, "finalfantasyxiv.com") {
			return "", fmt.Errorf("<%v> is an invalid Final Fantasy XIV host", urlParts.Host)
		}

		if !strings.HasPrefix(urlParts.Path, "/lodestone/") {
			return "", fmt.Errorf("<%v> is an invalid Lodestone path (missing </lodestone/> prefix)", urlParts.Path)
		}
	} else {
		urlParts.Path = reURLPath.ReplaceAllString(urlParts.Path, `/lodestone/$2`)
	}

	urlParts.Scheme = "http"
	urlParts.Host = "eu.finalfantasyxiv.com"

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
