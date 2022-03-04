package hello

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")

	resp, err := netClient.Get("http://example.com/")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Fprint(w, string(body))
}
