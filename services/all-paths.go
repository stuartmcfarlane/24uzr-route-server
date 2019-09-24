package services

import (
    "log"
    "sort"
    "24uzr-route-server/transport"
)

type adjacencyMatrix struct {
    nodeNames map[int]string
    nodeIdx map[string]int
    link [][]int
    weight [][]float32
}

type weightedPath struct {
    weight float32
    path []int
}

func FindAllPaths(route transport.Route, graph transport.Graph) transport.Routes {
    log.Println(">FindAllPaths")

    m := makeAdjacencyMatrix(graph)

    s := m.nodeIdx[route.Start];
    e := m.nodeIdx[route.End];

    p := weightedPath{weight: 0, path: []int{s}}
    paths := findPaths(p, s, e, m)

    sort.Slice(paths, func(i, j int) bool {
        return paths[i].weight > paths[j].weight
    })

    page := paths[:10]
    pathsOut := make([]transport.Path, len(page))
    for i, path := range page {
        p := make([]string, len(path.path))
        for j, n := range path.path {
            p[j] = m.nodeNames[n]
        }
        pathsOut[i] = transport.Path{Weight: path.weight, Nodes: p}
    }
    routes := transport.Routes{Start: route.Start, End: route.End, Paths: pathsOut}
    log.Println("<FindAllPaths", len(routes.Paths))
    return routes
}

func removeEdge(m adjacencyMatrix, s int, e int) adjacencyMatrix {
    if (m.link[s][e] > 0) {
        m.link[s][e] = m.link[s][e] - 1;
        m.link[e][s] = m.link[e][s] - 1;
    }
    return m;
}
func putEdge(m adjacencyMatrix, s int, e int) adjacencyMatrix {
    m.link[s][e] = m.link[s][e] + 1;
    m.link[e][s] = m.link[e][s] + 1;
    return m;
}

func findPaths( path weightedPath, start int, end int, matrix adjacencyMatrix) []weightedPath {
    if start == end {
        return []weightedPath{path}
    }

    if len(path.path) > 14 {
        return make([]weightedPath, 0)
    }

    children := getChildren(start, &matrix)

    paths := make([]weightedPath, 0)
    for _, c := range children {
        m := removeEdge(matrix, start, c)
        np := weightedPath{weight: path.weight, path:make([]int, len(path.path))}
        copy(np.path, path.path)
        np.path = append(np.path, c)
        np.weight = np.weight + matrix.weight[start][c]
        pp := findPaths(np, c, end, m)
        for _, p := range pp {
            paths = append(paths, p)
        }
        matrix = putEdge(m, start, c)
    }
    return paths
}

func getChildren(n int, m *adjacencyMatrix) []int {
    row := m.link[n]
    children := make([]int, 0)
    for c, edge := range row {
        if edge > 0 {
            children = append(children, c)
        }
    }
    return children
}

func getChildrenByWeight(n int, m *adjacencyMatrix) []int {
    children := getChildrenByWeight(n, m)
    sort.Slice(children, func(i, j int) bool {
        return m.weight[n][i] > m.weight[n][j]
    })
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
    link := make([][]int, nc)
    for i = 0; i < nc; i++ {
        link[i] = make([]int, nc)
    }
    for _, e := range g.Edges {
        si := ni[e.Start]
        ei := ni[e.End]
        link[si][ei] = 1
        link[ei][si] = 1
    }
    weight := make([][]float32, nc)
    for i = 0; i < nc; i++ {
        weight[i] = make([]float32, nc)
    }
    for _, e := range g.Edges {
        si := ni[e.Start]
        ei := ni[e.End]
        weight[si][ei] = e.WeightSE
        weight[ei][si] = e.WeightES
    }

    return adjacencyMatrix{nodeNames: nn, nodeIdx: ni, link: link, weight: weight}
}
