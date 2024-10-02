package transport

type Path struct {
	Metres  float32  `json:metres`
	Seconds float32  `json:seconds`
	Nodes   []string `json:nodes`
}
type Route struct {
	Start string  `json:start`
	End   string  `json:start`
	Time  float32 `json:time`
	Count int     `json:time`
	Path  Path    `json:path`
}

type Routes struct {
	Start string `json:start`
	End   string `json:start`
	Paths []Path `json:paths`
}

type Edge struct {
	Start             string  `json:start`
	End               string  `json:end`
	Metres            float32 `json:metres`
	MetresPerSecondSE float32 `json:metresPerSecondSE`
	MetresPerSecondES float32 `json:metresPerSecondES`
}
type Graph struct {
	Edges []Edge `json:edges`
}
