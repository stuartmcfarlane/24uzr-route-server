package services

import (
    "log"
    "github.com/twmb/algoimpl/go/graph"
    "24uzr-route-server/transport"
)

func FindShortestRoute(routeIn transport.Route, graphIn transport.Graph ) transport.Route {

    workingGraph := graph.New(graph.Undirected)

    nodes := make(map[string]graph.Node, 0)
    for _, edge := range graphIn.Edges {
        startId := edge.Start
        endId := edge.End
        _, prs := nodes[startId]
        if ! prs {
            n := workingGraph.MakeNode()
            nodes[startId] = n
            *n.Value = startId
            log.Println("added node", *n.Value)
        }
        _, prs = nodes[edge.End]
        if ! prs {
            n := workingGraph.MakeNode()
            nodes[endId] = n
            *n.Value = endId
            log.Println("added node", *n.Value)
        }
        workingGraph.MakeEdgeWeight(nodes[startId], nodes[endId], int(edge.Weight * 1000))
    }
    startNode := nodes[routeIn.Start]
    endNode := nodes[routeIn.End]

    log.Println("Start", startNode)
    log.Println("End", endNode)

    paths := workingGraph.DijkstraSearch(startNode)

    var foundPath graph.Path
    for key, path := range paths {
        log.Println("Path", key, path.Path)
        if len(path.Path) == 0 {
            continue
        }
        lastEdge := path.Path[len(path.Path) - 1]
        log.Println("last edge", lastEdge)
        if  lastEdge.End == endNode {
            foundPath = path
            break
        }
    }
    log.Println("found", foundPath)
    log.Printf("start node %T", *startNode.Value)
    var startId string
    startId = (*startNode.Value).(string)

    var pathOut []string
    pathOut = append(pathOut, startId)
    for _, edge := range foundPath.Path {
        log.Println("next node", *edge.End.Value)
        endId := (*edge.End.Value).(string)
        pathOut = append(pathOut, endId)
    }
    shortestRoute := transport.Route{
        Start:routeIn.Start,
        End: routeIn.End,
        Path: pathOut,
    }
    return shortestRoute
}
