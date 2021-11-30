package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
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

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bodyIo)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-type", "application/json")

	res, err := client.Do(req)
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
