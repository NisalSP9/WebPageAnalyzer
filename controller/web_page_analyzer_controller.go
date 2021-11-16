package controller

import (
	"github.com/NisalSP9/WebPageAnalyzer/models"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetHTMLPage(URL string) (models.HTMLPAGEINFOR,error) {
	var resData models.HTMLPAGEINFOR
	response, err := http.Get(URL)
	if err != nil {
		log.Printf("Error while REST call : %v\n", err.Error())
		return resData,err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error while ReadAll : %v\n", err.Error())
		return resData,err
	}
	reader := strings.NewReader(string(responseData))
	tokenizer := html.NewTokenizer(reader)
	hTagMap := make(map[string]int)
	pageStatus := true
	HTMLTagCount := 0
	htmlVersion := ""
	inaccessible := 0
	internal := 0
	external := 0
	hasLoginForm := false
	for pageStatus {
		tokenType := tokenizer.Next()
		tokenizer.AllowCDATA(true)
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return resData,err
			}
			log.Printf("Error: %v\n", tokenizer.Err())
		}
		tag, hasAttr := tokenizer.TagName()
		if strings.Contains(string(tokenizer.Raw()),"4.01"){
			htmlVersion = "4.01"
		}else if strings.Contains(string(tokenizer.Raw()),"1.0"){
			htmlVersion = "1.0"
		}else if strings.Contains(string(tokenizer.Raw()),"1.1"){
			htmlVersion = "1.1"
		}else {
			htmlVersion = "5"
		}
		if strings.Contains(string(tag), "html") {
			HTMLTagCount += 1
		}
		if HTMLTagCount == 2 {
			pageStatus = false
		}
		count, prs := hTagMap[string(tag)]
		if tokenType == html.StartTagToken {
			if "title" == string(tag) {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					resData.PageTitle = tokenizer.Token().Data
				}
			}
			if strings.Contains(string(tag), "h1") ||
				strings.Contains(string(tag), "h2") ||
				strings.Contains(string(tag), "h3") ||
				strings.Contains(string(tag), "h4") ||
				strings.Contains(string(tag), "h5") ||
				strings.Contains(string(tag), "h6") {
				if !prs {
					hTagMap[string(tag)] = 1
				} else {
					hTagMap[string(tag)] = count + 1
				}
			}
		}
		if strings.Contains(string(tag), "form") {
			hasLoginForm = true
		}
		if hasAttr{
			for {
				_, attrValue, moreAttr := tokenizer.TagAttr()
				strAttrValue := string(attrValue)
				if strings.Contains(strAttrValue, "http") {
					URLLink := strings.Split(URL, ".")
					if strings.Contains(strAttrValue, ","){
						URLArray := strings.Split(strAttrValue,",")
						for _,e := range URLArray{
							leftTrimedE := strings.TrimLeft(e," ")
							strLink := strings.Split(leftTrimedE," ")[0]
							resLink := strings.Split(strLink, ".")
							if len(resLink) > 1 && len(URLLink) > 1 {
								if resLink[0] == URLLink[0] {
									internal += 1
								} else {
									external += 1
								}
							}
							r, err := http.Get(strLink)
							if err != nil {
								log.Printf("Error while check access 1: %v\n", err.Error())
								break
							}
							if r.StatusCode != 200 {
								inaccessible += 1
							}
						}
					}else {
						strLink := strAttrValue
						resLink := strings.Split(strLink, ".")
						if len(resLink) > 1 && len(URLLink) > 1 {
							if resLink[0] == URLLink[0] {
								internal += 1
							} else {
								external += 1
							}
						}
						r, err := http.Get(strAttrValue)
						if err != nil {
							log.Printf("Error while check access 2: %v\n", err.Error())
							break
						}
						if r.StatusCode != 200 {
							inaccessible += 1
						}
					}
				}
				if !moreAttr {
					break
				}
			}
		}

	}
	resData.Headings = hTagMap
	resData.Inaccessible = inaccessible
	resData.Internal = internal
	resData.External = external
	resData.LoginForm = hasLoginForm
	resData.HTMLVersion = htmlVersion
	return resData,err
}
