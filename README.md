# 24uzr route server

## Getting started

```bash
cd ${GOPATH}
. .bash_profile
cd src/24uzr-route-server
go build
./24uzr-route-server
```

## Sample post data

## Rules

- Each leg may be traveled at most twice.
- The time to traverse the path must be no less that 23 hours and no more than 25 hours
- The most desirable path is the longest in distance

## Proposed algorithm

Exhaustive search

    1. Create an adjacency matrix of the nodes, where M[n,m] = 2 for each Enm 0 otherwise
    1.

## Rejected algorithms

1.  Generate a graph with 2 edges per leg
1.  Take all spanning trees of the graph
1.  Discard paths longer than 25hrs not ending at the end node
1.  Discard paths shorter than 23hrs
1.  Perform a DFS for the longest path with a heuristic backtracking scheme

Spanning trees are not what we are looking for. We need paths.

## Combinatorial Optimization

## References

1.  [yourbasic/graph](https://godoc.org/github.com/yourbasic/graph)
2.  [Graph spanning algorithms](https://link.springer.com/article/10.1007/s40747-018-0079-7)
3.  [Combinatorial Optimization](https://en.wikipedia.org/wiki/Combinatorial_optimization)
