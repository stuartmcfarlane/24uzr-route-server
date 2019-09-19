#24uzr route server

##Proposed algorithm

Take the minimum spanning tree of the graph
Discard paths longer then 24hrs not ending at the end node
Perform a DFS for the longest path with a heuristic backtraching scheme
