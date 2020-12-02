package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
)
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
    fmt.Println("RESTfulServ. on:5000, Controller:",r.URL.Path[1:])
}
func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting Restful services...")
    fmt.Println("Using port:5000")
    err := http.ListenAndServe(":5000", nil)
    log.Print(err)
    errorHandler(err)
}
func errorHandler(err error){
if err!=nil {
    fmt.Println(err)
    os.Exit(1)
}
}
