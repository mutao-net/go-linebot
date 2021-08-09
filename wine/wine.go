package wine

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type ResponesResults struct {
	Items []Item `json:"Items"`
}

type Item struct {
	Item struct {
		ItemName    string `json:"itemName"`
		ItemCaption string `json:"itemCaption"`
		ItemPrice   int    `json:"itemPrice"`
		ShopName    string `json:"shopName"`
		ItemURL     string `json:"itemUrl"`
	} `json:"Item"`
}

func run() ResponesResults {
	url := "https://app.rakuten.co.jp/services/api/IchibaItem/Search/20170706"
	params := setParams()
	return getItems(url + params)
}

func setParams() string {
	queries := []string{}
	params := map[string]string{
		"applicationId": os.Getenv("APPLICATION_ID"),
		"genreId":       "100317",
		"sort":          "+reviewAverage",
		"page":          "1",
		"hits":          "2"}
	for k, v := range params {
		query := k + "=" + url.QueryEscape(v)
		queries = append(queries, query)
	}
	return "?" + strings.Join(queries, "&")
}

func getItems(url string) ResponesResults {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	var results ResponesResults
	json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		log.Fatal(err)
	}
	for _, result := range results.Items {
		fmt.Println(result.Item.ItemName)
	}
	return results
}
