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
	"net/http"
)

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
	response.Write([]byte(request.URL.Path))
	fmt.Println(request.URL.Path)
}
func main() {
	var handle *requestHandler = &requestHandler{}
	fmt.Println("Server started on port 8000.")
	http.ListenAndServe(":8000", handle)
}
