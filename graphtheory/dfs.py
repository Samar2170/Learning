from collections import defaultdict

class Graph:
    def __init__(self) -> None:
        self.graph=defaultdict(list)
    
    def addEdge(self,u,v):
        self.graph[u].append(v)
    
    def DFSUtil(self,v,visited):
        visited.add(v)
        print(v,end=' ')
        for neighbour in self.graph[v]:
            if neighbour not in visited:
                self.DFSUtil(neighbour,visited)

    def DFS(self,v):
        visited=set()
        self.DFSUtil(v,visited)

    def print(self):
        for i in self.graph:
            print(i,'->',' -> '.join([str(j) for j in self.graph[i]]))

g=Graph()
g.addEdge(0,1)
g.addEdge(0,2)
g.addEdge(1,2)
g.print()
g.DFS(2)
g.DFS(0)
g.DFS(1)
