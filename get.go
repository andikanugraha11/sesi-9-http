package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	//log.Println(string(body))
	var respBody map[string]interface{}

	json.Unmarshal(body, &respBody)
	log.Println(respBody["title"])
}
