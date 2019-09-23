package services

import (
    "log"
    "24uzr-route-server/transport"
)

type adjacencyMatrix struct {
    nodeNames map[int]string
    nodeIdx map[string]int
    matrix [][]int
}

func FindAllPaths(route transport.Route, graph transport.Graph) transport.Routes {
    log.Println(">FindAllPaths")

    m := makeAdjacencyMatrix(graph)

    s := m.nodeIdx[route.Start];
    e := m.nodeIdx[route.End];

    p := []int{s}
    paths := findPaths(p, s, e, m)

    pathsOut := make([][]string, len(paths))
    for i, path := range paths {
        p := make([]string, len(path))
        for j, n := range path {
            p[j] = m.nodeNames[n]
        }
        pathsOut[i] = p
    }
    routes := transport.Routes{Start: route.Start, End: route.End, Paths: pathsOut}
    log.Println("<FindAllPaths", len(routes.Paths))
    return routes
}

func removeEdge(m adjacencyMatrix, s int, e int) adjacencyMatrix {
    if (m.matrix[s][e] > 0) {
        m.matrix[s][e] = m.matrix[s][e] - 1;
        m.matrix[e][s] = m.matrix[e][s] - 1;
    }
    return m;
}
func putEdge(m adjacencyMatrix, s int, e int) adjacencyMatrix {
    m.matrix[s][e] = m.matrix[s][e] + 1;
    m.matrix[e][s] = m.matrix[e][s] + 1;
    return m;
}
func findPaths( path []int, start int, end int, matrix adjacencyMatrix) [][]int {
    if start == end {
        return [][]int{path}
    }

    if len(path) > 12 {
        return make([][]int, 0)
    }

    children := getChildren(start, &matrix)

    paths := make([][]int, 0)
    for _, c := range children {
        m := removeEdge(matrix, start, c)
        np := make([]int, len(path))
        copy(np, path)
        np = append(np, c)
        pp := findPaths(np, c, end, m)
        for _, p := range pp {
            paths = append(paths, p)
        }
        matrix = putEdge(m, start, c)
    }
    return paths
}

func getChildren(n int, m *adjacencyMatrix) []int {
    row := m.matrix[n]
    children := make([]int, 0)
    for c, edge := range row {
        if edge > 0 {
            children = append(children, c)
        }
    }
    return children
}

func makeAdjacencyMatrix(g transport.Graph) adjacencyMatrix {

    nn := make(map[int]string);
    ni := make(map[string]int);
    i := int(0)
    // put the nodes in nodeName (nn) and nodeIndex (ni) maps
    for _, e := range g.Edges {
        for _, n := range []string{ e.Start, e.End } {
            _, prs := ni[n]
            if ! prs {
                ni[n] = i
                nn[i] = n
                i = i + 1
            }
        }
    }
    nc := len(nn)
    m := make([][]int, nc)
    for i = 0; i < nc; i++ {
        m[i] = make([]int, nc)
    }
    for _, e := range g.Edges {
        si := ni[e.Start]
        ei := ni[e.End]
        m[si][ei] = 1
        m[ei][si] = 1
    }

    return adjacencyMatrix{nodeNames: nn, nodeIdx: ni, matrix: m}
}
