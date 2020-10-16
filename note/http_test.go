package note

import (
	"testing"
	"time"
)

func TestHttpGet (t *testing.T){
	url := "https://httpbin.org/get"
	HttpGet(url)
}

func TestHttpPostJson(t *testing.T) {
	url := "https://httpbin.org/post"
	data := map[string]string{
		"aaa": "111",
		"bbb": "222",
	}
	HttpPostJson(url, data)
}

func TestHttpCustomPostJson(t *testing.T) {
	// ???
	timeout := time.Duration(time.Second * 5)
	url := "https://httpbin.org/post"
	data := map[string]string{
		"ccc": "333",
		"ddd": "333",
	}
	HttpCustomPostJson(timeout, url, data)
}