package api

import (
    "fmt"
    "net/http"
)

// type helloWorldHandler struct{
    
// }

// func (h *helloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//  fmt.Fprintf(w, "hey")
// }

func DemoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "called demo handler")
}