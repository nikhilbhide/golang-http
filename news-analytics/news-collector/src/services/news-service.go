package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type newsheadlines struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	} `json:"articles"`
}

func main() {
	fmt.Println("Starting the application...")
	response, err := http.Get("https://newsapi.org/v2/top-headlines?country=in&apiKey=456f513b38d14b4a8c96fd267274c62d")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		res := newsheadlines{}

		if err := json.Unmarshal(data, &res); err != nil {
			panic(err)
		}
		//fmt.Println(res)

		for _,article:= range res.Articles  {
			fmt.Println(article.Author)
			fmt.Println(article.PublishedAt)
			fmt.Println(article.Title)
			fmt.Println(article.Description)


			//fmt.Println(article.Content)
		}
	}
}
