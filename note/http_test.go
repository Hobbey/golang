package note

import "testing"

func TestHttpGet (t *testing.T){
	url := "https://httpbin.org/get"
	HttpGet(url)
}