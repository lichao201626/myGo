package main

import (
	"bytes"
	"fmt"
	"github.com/elastic/go-elasticsearch/client"
)

func main() {
	c, _ := client.New(client.WithHost("http://127.0.0.1:9200/user"))
	body := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"user": "kimchy",
			},
		},
	}
	resp, err := c.Search(body)
	fmt.Println(c)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range resp.Response.Header {
		fmt.Println(v)
	}
	fmt.Println(resp.Response.Request)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Response.Body)
	fmt.Println(buf)

	/* 	ch := make(chan string)
	   	close(ch)
	   	chan<- resp */
	/* 	for _, v := range resp.Response.Body {
		fmt.Println(v)
	} */
	/* 	body := map[string]interface{}{
	   		"query": map[string]interface{}{
	   			"term": map[string]interface{}{
	   				"user": "kimchy",
	   			},
	   		},
	   	}
	   	fmt.Println(c)
	   	resp, err := c.Msearch(body) */
}
