class Node:
    def __init__(self, key, val):
        self.key = key
        self.val = val
        self.next = None
        self.prev = None

class LRUCache:
    def __init__(self, capacity: int) -> None:
        self.cap = capacity
        self.cache = {}
        self.left, self.right = Node(0,0), Node(0,0)
        self.left.next, self.right.prev = self.right,self.left

    def remove(self, node):
        prev,next=node.prev, node.next
        prev.next = next
        next.prev = prev


    def insert(self, node: Node):
        #insert node at right
        prev,next = self.right.prev, self.right
        prev.next = node
        next.prev = node
        node.next, node.prev = next, prev

    def get(self, key:int) -> int:
        if key in self.cache:
            self.remove(self.cache[key])
            self.insert(self.cache[key])
            return self.cache[key].val
        return -1
    
    def put(self, key:int, val:int) -> None:
        if key in self.cache:
            self.remove(self.cache[key])
        self.cache[key] = Node(key, val)
        self.insert(self.cache[key])

        if len(self.cache)>self.cap:
            lru = self.left.next
            self.remove(lru)
            del self.cache[key]