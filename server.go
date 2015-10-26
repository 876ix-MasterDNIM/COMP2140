/**
 * COMP2140 Project
 *
 */

/**
*Main Package to start server and handle requests.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var handle *requestHandler = &requestHandler{}
	fmt.Println("Server started on port 8000.")

	http.Handle("/login", handle)
	http.ListenAndServe(":8000", new(requestHandler))
}

/**
 * SERVER IS PRETTY MUCH NON FUNCTIONAL AT THE MOMENT.
 */
type requestHandler struct{}

/**
 * Demo Function to handle http requests
 * @param  {*requestHandler} handler *requestHandler [description]
 * @return {[type]}         [description]
 */
func (handler *requestHandler) ServeHTTP(response http.ResponseWriter,
	request *http.Request) {
	var path string = string(request.URL.Path[1:])
	//fmt.Println(path)
	webpageData, error := ioutil.ReadFile(path)

	if error == nil {
		var contentType string = addContentType(path)

		response.Header().Set("Content-Type", contentType)
		response.Write(webpageData)
	} else {
		response.WriteHeader(404)
		response.Write([]byte(http.StatusText(404)))
	}
	fmt.Println(request.URL.Path)
}

func addContentType(path string) string {
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/js; charset=utf-8"
	} else if strings.HasSuffix(path, ".html") {
		contentType = "text/html; charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpeg; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}
	return contentType
}
