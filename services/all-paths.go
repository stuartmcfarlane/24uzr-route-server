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
    log.Println("matrix", m)

    s := m.nodeIdx[route.Start];
    e := m.nodeIdx[route.End];

    p := []int{s}
    paths := findPaths(p, s, e, m)
    log.Println("paths", paths)

    pathsOut := make([][]string, len(paths))
    for i, path := range paths {
        p := make([]string, len(path))
        for j, n := range path {
            p[j] = m.nodeNames[n]
        }
        pathsOut[i] = p
    }
    routes := transport.Routes{Start: route.Start, End: route.End, Paths: pathsOut}
    log.Println("<FindAllPaths", routes)
    return routes
}

func removeEdge(m adjacencyMatrix, s int, e int) adjacencyMatrix {
    if (m.matrix[s][e] > 0) {
        m.matrix[s][e] = m.matrix[s][e] - 1;
        m.matrix[e][s] = m.matrix[e][s] - 1;
    }
    return m;
}
func findPaths( path []int, start int, end int, matrix adjacencyMatrix) [][]int {
    log.Println(">findPaths", start, end, path)
    if start == end {
        log.Println("<findPaths found", path)
        return [][]int{path}
    }

    children := getChildren(start, &matrix)
    log.Println("children", children)

    paths := make([][]int, 0)
    for _, c := range children {
        pp := findPaths(append(path, c), c, end, removeEdge(matrix, start, c))
        for _, p := range pp {
            paths = append(paths, p)
        }
    }

    if len(paths) > 0 { log.Println("<findPaths found", len(paths)) }
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
