package main

import (
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"net/http"
	"time"
)

//type Item struct {
//	Title   string `json:"title" xml:"title"`
//	Link    string `json:"link" xml:"link"`
//	Desc    string `json:"desc" xml:"description"`
//	PubDate string `json:"pub_date" xml:"pubDate"`
//}
//
//type Channel struct {
//	Title string `json:"title" xml:"title"`
//	Link  string `json:"link" xml:"link"`
//	Desc  string `json:"desc" xml:"description"`
//	Lang  string `json:"lang" xml:"language"`
//	Items []Item `json:"items" xml:"item"`
//}

type NewItem struct {
	Title       string    `json:"title" xml:"title"`
	Link        string    `json:"link" xml:"link"`
	Description string    `json:"description" xml:"description"`
	Author      string    `json:"author" xml:"author"`
	Enclosure   Enclosure `json:"enclosure" xml:"enclosure"`
	Source      string    `json:"source" xml:"source"`
	PubDate     string    `json:"pub_date" xml:"pubDate"`
}

//read enclosure tag properties like url, length, type
type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type NewChannel struct {
	Title         string    `json:"title" xml:"title"`
	Link          string    `json:"link" xml:"link"`
	PubDate       string    `json:"pub_date" xml:"pubDate"`
	Language      string    `json:"language" xml:"language"`
	LastBuildDate string    `json:"last_build_date" xml:"lastBuildDate"`
	Items         []NewItem `json:"items" xml:"item"`
}

type Rss struct {
	Channel NewChannel `json:"channel" xml:"channel"`
}

func main() {

	//addrURL := "https://www.straitstimes.com/latest/web_rss.xml"
	addrURL2 := "https://www.straitstimes.com/latest/ios_homescreen_rss.xml"
	GetFeeds(addrURL2)

}

func GetFeeds(url string) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		errMsg := fmt.Sprintf("error when initialize http %s request URL %s", http.MethodGet, url)
		fmt.Println(err, errMsg)
		return
	}

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 2
	retryClient.RetryWaitMin = 3 * time.Second
	retryClient.RetryWaitMax = 9 * time.Second
	httpCli := retryClient.StandardClient()
	resp, err := httpCli.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("req URL: %s", req.URL)
		fmt.Println(err, errMsg)
		return
	}
	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("actual status code: %d", resp.StatusCode)
		fmt.Println(errMsg)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("error occurs when reading response body. BodyWriter:%v", resp.Body)
		fmt.Println(err, errMsg)
		return
	}

	rss := Rss{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println("Error Unmarshal to xml", err)
		return
	}

	if len(rss.Channel.Items) <= 0 {
		fmt.Println("no feeds data found while pull from rss")
		return
	}

	// prepare rss feeds []item data
	var rssFeeds []NewItem
	rssFeeds = rss.Channel.Items
	//fmt.Println("feeds data : ", rssFeeds)

	//respJson, err := json.Marshal(rssFeeds)
	//if err != nil {
	//	fmt.Println("Error Marshal to json", err)
	//	return
	//}

	//fmt.Println("resp : ", string(respJson))

	for index, item := range rssFeeds {
		fmt.Println("item : ", item)
		fmt.Println("index : ", index, "url :", item.Enclosure.URL)
	}
}
