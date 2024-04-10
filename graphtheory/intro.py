
am_ab = [
    [0,4,1,9],
    [3,0,6,11],
    [4,1,0,2],
    [6,5,2,0]
]

al_ab = [
    [('B', 4), ('C', 1), ('D', 9)],
    [('A', 3), ('C', 6), ('D', 11)],
    [('A', 4), ('B', 1), ('D', 2)],
    [('A', 6), ('B', 5), ('C', 2)]
]

class Graph:
    def __init__(self, vertices):
        self.vertices = vertices
        self.graph = {}
        for vertex in vertices:
            self.graph[vertex] = []

    def add_edge(self, u, v):
        self.graph[u].append(v)
        self.graph[v].append(u)

    def print_graph(self):
        for vertex in self.vertices:
            print(vertex, '->', ' -> '.join([str(i) for i in self.graph[vertex]]))

    def bfs(self, start):
        visited = [False] * len(self.vertices)
        queue = []
        queue.append(start)
        visited[start] = True
        while queue:
            start = queue.pop(0)
            print(start, end = ' ')
            for i in self.graph[start]:
                if visited[i] == False:
                    queue.append(i)
                    visited[i] = True

    def dfs(self, start):
        visited = [False] * len(self.vertices)
        self.dfs_util(start, visited)

    def dfs_util(self, start, visited):
        visited[start] = True
        print(start, end = ' ')
        for i in self.graph[start]:
            if visited[i] == False:
                self.dfs_util(i, visited)