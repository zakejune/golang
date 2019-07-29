package main()
import {
    "ftm"
    "net/http"
}
func hello(res http.ResponseWriter,req *http.Request) {
    fmt.Fprint(res,"Hello,golang")

}

func main() {
    http.HandleFunc("/",hello)
    http.ListenAndServe("localhost:4000",nil)
}
