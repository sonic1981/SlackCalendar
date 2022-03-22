package SlackCalender

import (
	"fmt"
	"net/http"
)

// HelloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!\n")

	channels := GetChannelsObj()

	for index, element := range channels.Channels {

		fmt.Println(index, element.Name)

	}

	fmt.Fprint(w, "Done!\n")
}
