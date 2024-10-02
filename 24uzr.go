package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"24uzr-route-server/services"
	"24uzr-route-server/transport"
)

const RACE_TIME_HOURS = 10.0
const RACE_TIME_SECONDS = RACE_TIME_HOURS * 60 * 60

func findShortestRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	query := r.URL.Query()
	route := transport.Route{Start: query["start"][0], End: query["end"][0]}
	var graph transport.Graph
	json.NewDecoder(r.Body).Decode(&graph)
	route = services.FindShortestPath(route, graph)
	json.NewEncoder(w).Encode(route)
}

func findAllRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	query := r.URL.Query()
	var time float64
	if len(query["time"]) == 0 {
		time = RACE_TIME_SECONDS
	} else {
		parsedTime, err := strconv.ParseFloat(query["time"][0], 64)
		if err != nil {
			time = RACE_TIME_SECONDS
		} else {
			time = parsedTime
		}
	}
	var count = 10
	log.Println("query", query)
	if len(query["count"]) > 0 {
		parsedCount, err := strconv.Atoi(query["count"][0])
		if err == nil {
			log.Println("parsedCount", parsedCount)
			count = parsedCount
		}
	}

	route := transport.Route{
		Start: query["start"][0],
		End:   query["end"][0],
		Time:  float32(time),
		Count: count,
	}
	var graph transport.Graph
	json.NewDecoder(r.Body).Decode(&graph)
	routes := services.FindAllPaths(route, graph, float32(time), count)
	json.NewEncoder(w).Encode(routes)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/route/shortest", findShortestRoute).Methods("POST")
	router.HandleFunc("/route/all", findAllRoutes).Methods("POST")

	fmt.Println("Listening on port 3002")
	log.Fatal(http.ListenAndServe(":3002", router))
}
