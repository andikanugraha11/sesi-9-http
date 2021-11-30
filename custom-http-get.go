package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		panic(err)
	}


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
