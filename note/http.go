package note

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HttpGet(url string){
	// return response and error value
	resp, err := http.Get(url)
	if err != nil {
		// print and call os.Exit(1)
		log.Fatalln(err)
	}
	// forgetting to close the response body can cause resource leaks in a long running programs
	defer resp.Body.Close()

	// The resp.Body implements the io.Reader interface allowing us to use ioutil.ReadAll function
	// since response body is an io.Reader, you could read the data chunk by chunk and process it as a stream of data
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Print Result
	log.Println(string(body))
}


func HttpPostJson(url string, data map[string]string){
	// marshalling a map, get []byte
	requestBody, err := json.Marshal(data)
	if err != nil{
		log.Fatalln(err)
	}

	// func post need url + content type(json) + io.Reader
	// []byte doesn't implement this interface, use bytes.NewBuffer
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// process resp, print result
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln(err)
	}
	log.Println(string(body))

}

func HttpCustomPostJson(timeoutSecond time.Duration, url string, data map[string]string) {
	// request body
	requestBody, err := json.Marshal(data)
	if err != nil{
		log.Fatalln(err)
	}

	// create a new client, customize Timeout
	timeout := timeoutSecond
	client := http.Client{
		Timeout: timeout,
	}

	// create a new request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("content-type", "application/json")

	// execute POST, call client.Do with request
	resp, err := client.Do(request)
	if err != nil{
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// process resp, print result
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln(err)
	}
	log.Println(string(body))

}