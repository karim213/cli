package main

import (
	"os"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"
	"bytes"
)


func getVideos(getCmd *flag.FlagSet, all *bool, id *string) {

	getCmd.Parse(os.Args[2:]) // parse all the entrees to be understood

	Url := fmt.Sprintf("http://localhost:8080/videos?all=%t&id=%s", *all, *id) // construct the url with the parameters
	resp, err := http.Get(Url) // send the request
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body) // reading the response
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body) // cast data to string
	log.Printf(sb) // print response
}




func saveVideos(addCmd *flag.FlagSet, id *string, title *string, description *string, imageUrl *string, url *string) {

	addCmd.Parse(os.Args[2:]) // parse all the entrees to be understood

	// construct the body of the request as Json format
	postBody, _ := json.Marshal(map[string]string{
		"id": *id,
		"title": *title,
		"description": *description,
		"imageUrl": *imageUrl,
		"url": *url,
	})

	// convert the encoded JSON data to a type implemented by the io.Reader interface
	requestBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:8080/videos", "application/json", requestBody) // send the request
	if err != nil {
		log.Fatalf("An error Occured %v", err)
	}

	defer resp.Body.Close() // close the response body
	body, err := ioutil.ReadAll(resp.Body) // read the response
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body) // cast data to string
   	log.Printf(sb) // print response
}
