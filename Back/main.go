package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mmcdole/gofeed"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle actual GET request
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, World!"}`))
}

type RssFeedHandler struct{}

type RssFeedItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type Response struct {
	Title       string `json:"title"`
	Description string `json:"description"`

	FeedDataItems []RssFeedItem `json:"feedDataItems"`
}

func (h *RssFeedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle actual GET request
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte(`{"message": "Hello, World!"}`))

	feed, err := gofeed.NewParser().ParseURL("https://zenn.dev/spiegel/feed")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	feedDataItems := []RssFeedItem{}

	for _, item := range feed.Items {
		feedDataItems = append(feedDataItems, RssFeedItem{Title: item.Title, Description: item.Description, Link: item.Link})
	}

	response := Response{Title: feed.Title, Description: feed.Description, FeedDataItems: feedDataItems}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("response: %v", jsonResponse)
	//w.Write([]byte(` {"message": feed.Title}`))
	w.Write(jsonResponse)
}

func main() {
	//スタート時間・処理時間表示
	startTime := time.Now()
	fmt.Printf("start time: %v \n", startTime)
	defer func() {
		fmt.Printf("\n processing time: %v", time.Since(startTime).Milliseconds())
	}()

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	HelloHandler := HelloHandler{}

	// APIのエンドポイントを設定
	http.Handle("/hello", &HelloHandler)

	RssFeedHandler := RssFeedHandler{}
	http.Handle("/get_feed", &RssFeedHandler)

	server.ListenAndServe()
}
