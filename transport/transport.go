package transport

type Path struct {
    Weight float32
    Nodes []string `json:path`
}
type Route struct {
    Start string `json:start`
    End string `json:start`
    Path Path `json:path`
}

type Routes struct {
    Start string `json:start`
    End string `json:start`
    Paths []Path `json:paths`
}

type Edge struct {
    Start string `json:start`
    End string `json:end`
    Weight float32 `json:weight`
    WeightSE float32 `json:weightSE`
    WeightES float32 `json:weightES`
}
type Graph struct {
    Edges []Edge `json:edges`
}

