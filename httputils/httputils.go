package httputils

import (
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/*
ParseAllAssetsSequential takes a string of text (typically from a http.Response.Body)
and return the urls for the page <script> <link> and <img> tag.
The method runs lineearly.
In benchmark you will see that ParseAllAssets is generally faster and GetAssets is faster still
*/
func ParseAllAssetsSequential(body string) (js, img, css []string) {
	jsfiles := GetJS(body)
	cssfiles := GetCSS(body)
	imgfiles := GetIMG(body)
	return jsfiles, imgfiles, cssfiles
}

/*
GetAssets takes a string of test from an http.Response.Body and returns the
urls for the page <script>, <link>, and <img> tags.
It makes use of the goquery library and is currently the fastest method
*/
func GetAssets(body string) (js, img, css []string) {
	utfBody := strings.NewReader(body)
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Println("Unable to parse document with goquery.  Make sure it is utf8")
		return ParseAllAssets(body)
	}
	goroutine := false
	jsfiles := []string{}
	imgfiles := []string{}
	cssfiles := []string{}

	if goroutine {
		c1 := make(chan []string)
		c2 := make(chan []string)
		c3 := make(chan []string)

		go func() { c1 <- getAttr(doc, ScriptTag, SrcAttribute) }()
		go func() { c2 <- getAttr(doc, ImageTag, SrcAttribute) }()
		go func() { c3 <- getAttr(doc, LinkTag, HrefAttribute) }()

		for i := 0; i < AssetTypesCount; i++ {
			select {
			case jsfiles = <-c1:
			case imgfiles = <-c2:
			case cssfiles = <-c3:
			}
		}
	} else {
		jsfiles = getAttr(doc, ScriptTag, SrcAttribute)
		imgfiles = getAttr(doc, ImageTag, SrcAttribute)
		cssfiles = getAttr(doc, LinkTag, HrefAttribute)
	}

	return jsfiles, imgfiles, cssfiles
}

/*
geAttr takes a *goquery.Document a html tag and attr
and returns a list of those attributes
*/
func getAttr(doc *goquery.Document, tag, attr string) []string {
	files := []string{}
	doc.Find(tag).Each(func(_ int, s *goquery.Selection) {
		value, exists := s.Attr(attr)
		if exists {
			files = append(files, value)
		}
	})
	return files
}

/*
ParseAllAssets takes string of text (typically from a http.Response.Body)
and return the urls for the page <script> <link> and <img> tag.
The method uses separate go routines for each asset class.
It is faster than ParseAllAssetsSequentially, but still slower than GetAssets
*/
func ParseAllAssets(body string) (js, img, css []string) {
	// make some channels
	c1 := make(chan []string)
	c2 := make(chan []string)
	c3 := make(chan []string)

	// kick off our anonymous go routines.
	go func() { c1 <- GetJS(body) }()
	go func() { c2 <- GetIMG(body) }()
	go func() { c3 <- GetCSS(body) }()

	// collect our results
	jsfiles := []string{}
	imgfiles := []string{}
	cssfiles := []string{}

	for i := 0; i < AssetTypesCount; i++ {
		select {
		case jsfiles = <-c1:
		case imgfiles = <-c2:
		case cssfiles = <-c3:
		}
	}

	return jsfiles, imgfiles, cssfiles
}

// GetJS uses regex to parse a body of text and return the script src attributes
func GetJS(body string) []string {
	return runregex(ScriptSrcPattern, body)
}

// GetCSS uses regex to parse a body of text and return the <link> href attributes
func GetCSS(body string) []string {
	return runregex(LinkHrefPattern, body)
}

// GetIMG uses regex to parse a body of text and return the <img> src attributes
func GetIMG(body string) []string {
	backgroundimgs := runregex(BackgroundImagePattern, body)
	imgs := runregex(ImageSrcPattern, body)
	imgs = append(imgs, backgroundimgs...)
	return imgs
}

// Take a regex expression that returns the matched object
// and return an array of the matched text
func runregex(expr, body string) []string {
	r, _ := regexp.Compile(expr)
	match := r.FindAllStringSubmatch(body, RegexMatchLimit)
	files := make([]string, 0)
	for j := 0; j < len(match); j++ {
		files = append(files, match[j][1])
	}
	return files
}
