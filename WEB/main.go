package main 
import (
	"fmt"
	"net/http"
	"os"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
	log "github.com/llimllib/loglevel"
)

type Link struct {
	url string
	text string
	depth int
}

type HttpError struct {
	original string
}

func LinkReader(resp *http.Response,depth int) []Link {
	page := html.NewTokenizer(resp,Body)
	links := []Link{}
}