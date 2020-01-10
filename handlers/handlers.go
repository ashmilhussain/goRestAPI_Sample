package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go , You are visiting Directory :/ %q", html.EscapeString(r.URL.Path))
}

func TestRest(w http.ResponseWriter, r *http.Request) {

	type TestJSON struct {
		Name string
		Data string
	}

	type TestJSONs []TestJSON

	testVariable := TestJSONs{
		TestJSON{Name: "Test1"},
		TestJSON{Name: "Test1"},
	}

	json.NewEncoder(w).Encode(testVariable)
}
