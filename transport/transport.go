package transport

type Route struct {
    Start string `json:start`
    End string `json:start`
}

type Edge struct {
    Start string `json:start`
    End string `json:end`
    Weight float32 `json:weight`
}
type Graph struct {
    Edges []Edge `json:edges`
}

