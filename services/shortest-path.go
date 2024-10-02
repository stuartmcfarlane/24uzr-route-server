package services

import (
	"24uzr-route-server/transport"
	"fmt"

	"github.com/twmb/algoimpl/go/graph"
)

func FindShortestPath(routeIn transport.Route, graphIn transport.Graph) transport.Route {
	fmt.Println(">FindShortestPath")

	workingGraph := graph.New(graph.Undirected)

	// Build the graph
	nodes := make(map[string]graph.Node, 0)
	for _, edge := range graphIn.Edges {
		startId := edge.Start
		endId := edge.End
		_, prs := nodes[startId]
		if !prs {
			n := workingGraph.MakeNode()
			nodes[startId] = n
			*n.Value = startId
		}
		_, prs = nodes[edge.End]
		if !prs {
			n := workingGraph.MakeNode()
			nodes[endId] = n
			*n.Value = endId
		}
		workingGraph.MakeEdgeWeight(nodes[startId], nodes[endId], int(edge.Metres*1000))
	}
	startNode := nodes[routeIn.Start]
	endNode := nodes[routeIn.End]

	// get the shortest path to each node
	paths := workingGraph.DijkstraSearch(startNode)

	// find the path ending at endNode
	var foundPath graph.Path
	for _, path := range paths {
		if len(path.Path) == 0 {
			continue
		}
		lastEdge := path.Path[len(path.Path)-1]
		if lastEdge.End == endNode {
			foundPath = path
			break
		}
	}

	// build array of id's to return
	var pathOut []string
	pathOut = append(pathOut, (*startNode.Value).(string))
	for _, edge := range foundPath.Path {
		pathOut = append(pathOut, (*edge.End.Value).(string))
	}

	// build route for transport
	shortestRoute := transport.Route{
		Start: routeIn.Start,
		End:   routeIn.End,
		Path:  transport.Path{Metres: float32(len(pathOut)), Nodes: pathOut},
	}
	fmt.Println("<FindShortestPath")
	return shortestRoute
}
