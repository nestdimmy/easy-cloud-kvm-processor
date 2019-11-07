package controllers

import (
	"fmt"
	"net/http"
)

var Test = func(w http.ResponseWriter, r *http.Request) {
	fmt.Print("API works")
}
