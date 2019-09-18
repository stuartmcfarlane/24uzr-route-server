package services

import (
    "log"

    "24uzr-route-server/transport"
)

func FindShortestRoute(route transport.Route, graph transport.Graph ) transport.Route {
    log.Println("route in", route)
    log.Println("graph in", graph)

    return route
}
