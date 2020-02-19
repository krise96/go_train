package routers

import (
	"fmt"
	"io"
	"net/http"
)

func HandleHello() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if _, err := io.WriteString(res, "Hello World"); err != nil {
			fmt.Print("Some error")
		}
	}

}
