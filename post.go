package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {

	bodyRequest := map[string]interface{}{
		"title": "title example",
		"userId": 1,
		"body": "body example",
	}

	bodyJson, err := json.Marshal(bodyRequest)
	if err != nil {
		panic(err)
	}
	bodyIo := bytes.NewBuffer(bodyJson)
	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bodyIo)
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
	log.Println(respBody)
}
