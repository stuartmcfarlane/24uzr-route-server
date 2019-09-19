#24uzr route server

##Rules

 - Each leg may be travelled at most twice.
 - The time to traverse the path must be no less that 23 hours and no more than 25 hours
 - The most desirable path is the longest in distance

##Proposed algorithm

 1. Generate a graph with 2 edges per leg
 1. Take all spanning trees of the graph [reference](https://link.springer.com/article/10.1007/s40747-018-0079-7)
 1. Discard paths longer than 25hrs not ending at the end node
 1. Discard paths shorter than 23hrs
 1. Perform a DFS for the longest path with a heuristic backtracking scheme