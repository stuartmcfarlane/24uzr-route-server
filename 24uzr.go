package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    // "gonum.org/v1/gonum/graph/path"

    "24uzr-route-server/transport"

)

func findShortestRoute(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "application/json")
    query := r.URL.Query()
    route := transport.Route{ Start: query["start"][0], End: query["end"][0] }
    var graph transport.Graph
    json.NewDecoder(r.Body).Decode(&graph)
    log.Println("params", graph)
    json.NewEncoder(w).Encode(route);
}

func main() {

    router := mux.NewRouter()

    router.HandleFunc("/route/shortest", findShortestRoute).Methods("POST")

    log.Fatal(http.ListenAndServe(":3002", router))
}
