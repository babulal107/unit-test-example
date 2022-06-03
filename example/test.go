package main

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func main() {
	article := Article{}
	val := "{\"title\":\"this is test article \",\"body\":\"   \"}"
	ba := []byte(val)
	_ = json.Unmarshal(ba, &article)

	fmt.Println("title : ", article.Title)
	if len(article.Body) == 0 {
		fmt.Println("error body is empty")
	}
	fmt.Println("body : ", article.Body)
}
