## INTRO 
Graph is a non linear data structure consisting of nodes/vertices and edges.
Graph is a composed set of vertices/nodes and edges/links. 
Denoted as G = (V, E) where V is set of vertices and E is set of edges.

Graphs are of two types:
1. Directed Graphs: Graphs with directed edges.
2. Undirected Graphs: Graphs with undirected edges.

Graphs can be represented in two ways:
1. Adjacency Matrix: A matrix of size VxV where V is number of vertices.
2. Adjacency List: A list of lists where each list represents a vertex and its adjacent vertices.

## Questions to ask while solving graph theory probs
1. Is the graph directed or undirected?
2. Is the graph weighted or unweighted?
3. Is the graph connected or disconnected?
4. Use adjancency matrix or adjancency list or Edge list?

### PROBLEMS
1. Shortest Path, given weighted path find shortest path between two vertices.
2. Does a path exist between two vertices?
3. Negative cycles in a graph. (a trap in shortest path, negative weight cycle) currency arbitrage.
4. Strongly connected components in a graph. used in social networks.
5. Travelling salesman problem. (NP hard)
6. Bridges in a graph. (cut edge) useful in network design. identify weak points.
7. Articulation points in a graph. (cut vertex) useful in network design. identify weak points.
8. Minimum spanning tree. (MST) used in network design.
9. Network flow problems. (max flow, min cut) used in network design.


### TYPES
1. Directed Graphs : Graphs with directed edges.
2. Undirected Graphs : Graphs with undirected edges.
3. Weighted Graphs : Graphs with weighted edges.
4. Unweighted Graphs : Graphs with unweighted edges.
5. Connected Graphs : Graphs with a path between every pair of vertices.
6. Disconnected Graphs : Graphs with no path between every pair of vertices.
7. Cyclic Graphs : Graphs with a cycle.
8. Acyclic Graphs : Graphs with no cycle.
9. Bipartite Graphs : Graphs with vertices divided into two disjoint sets.
10. Complete Graphs : Graphs with every pair of vertices connected by an edge.
11. Sparse Graphs : Graphs with few edges.
12. Dense Graphs : Graphs with many edges.
13. Planar Graphs : Graphs that can be drawn on a plane without edges crossing.
14. Non-Planar Graphs : Graphs that cannot be drawn on a plane without edges crossing.
15. Regular Graphs : Graphs with every vertex having the same degree.
16. Irregular Graphs : Graphs with vertices having different degrees.

### REPRESENTATION
1. Adjacency Matrix : A matrix of size VxV where V is number of vertices.
2. Adjacency List : A list of lists where each list represents a vertex and its adjacent vertices.
3. Edge List : A list of tuples where each tuple represents an edge.

### TRAVERSAL
1. Breadth First Search (BFS) : Traverses the graph level by level.
2. Depth First Search (DFS) : Traverses the graph by going as far as possible along a branch and backtracking.

### ALGORITHMS
1. Dijkstra's Algorithm : Finds the shortest path between two vertices in a weighted graph.
2. Bellman-Ford Algorithm : Finds the shortest path between two vertices in a weighted graph with negative weights.
3. Floyd-Warshall Algorithm : Finds the shortest path between all pairs of vertices in a weighted graph.
4. Kruskal's Algorithm : Finds the minimum spanning tree of a connected, undirected graph.
5. Prim's Algorithm : Finds the minimum spanning tree of a connected, undirected graph.
6. Tarjan's Algorithm : Finds the strongly connected components of a directed graph.
7. Kosaraju's Algorithm : Finds the strongly connected components of a directed graph.
8. Travelling Salesman Problem : Finds the shortest possible route that visits every vertex exactly once and returns to the starting vertex.
9. Ford-Fulkerson Algorithm : Finds the maximum flow in a network.
10. Edmonds-Karp Algorithm : Finds the maximum flow in a network.
11. Hopcroft-Karp Algorithm : Finds the maximum cardinality matching in a bipartite graph.
12. Hungarian Algorithm : Finds the maximum weight matching in a bipartite graph.
13. Topological Sorting : Arranges the vertices in a directed acyclic graph in a linear order.
14. Bridges in a Graph : Finds the bridges in an undirected graph.
15. Articulation Points in a Graph : Finds the articulation points in an undirected graph.
16. Negative Cycles in a Graph : Finds the negative cycles in a weighted graph.

### APPLICATIONS
1. Social Networks : Representing relationships between people.
2. Maps : Representing routes between locations.
3. Computer Networks : Representing connections between devices.
4. Recommendation Systems : Representing relationships between items.
5. Bioinformatics : Representing relationships between proteins.
6. Scheduling Problems : Representing tasks and their dependencies.
7. Linguistics : Representing relationships between words.
8. Chemistry : Representing molecular structures.
9. Physics : Representing interactions between particles.
10. Electrical Engineering : Representing circuits.

### PROPERTIES
1. Degree : The number of edges incident to a vertex.
2. In-Degree : The number of edges entering a vertex in a directed graph.
3. Out-Degree : The number of edges leaving a vertex in a directed graph.
4. Weight : The value assigned to an edge in a weighted graph.
5. Path : A sequence of vertices connected by edges.
6. Cycle : A path that starts and ends at the same vertex.
7. Connectedness : The property of a graph having a path between every pair of vertices.
8. Acyclic : The property of a graph having no cycles.
9. Bipartite : The property of a graph having vertices divided into two disjoint sets.
10. Complete : The property of a graph having every pair of vertices connected by an edge.

