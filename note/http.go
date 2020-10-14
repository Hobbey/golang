package note

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HttpGet(url string){
	// http.Get 返回 *Response, Response 里有Body等方法
	resp, err := http.Get(url)
	if err != nil {
		// print and call os.Exit(1)
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// 处理得到的数据, Body 是 io.ReadCloser,
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Print Result
	log.Println(string(body))
}
