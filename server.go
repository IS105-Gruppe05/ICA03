package main

import (
	"net/http"
	"time"

)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<font color=\"green\">Hvordan g\u00E5r det?</font><br/>"))
	w.Write([]byte(time.Now().Format(time.RFC850)))
	w.Write([]byte("\u23f0"))
}
