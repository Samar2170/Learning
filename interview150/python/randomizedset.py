import random
class RandomizedSet:

    def __init__(self):
        self.dataset = {}
        self.datalist = []


    def insert(self, val: int) -> bool:
        if val not in self.dataset:
            self.datalist.append(val)
            self.dataset[val] = len(self.datalist)-1
            return True
        return False

    def remove(self, val: int) -> bool:
        if val in self.dataset:
            indx = self.dataset[val]
            self.dataset[self.datalist[-1]] = indx
            if indx!=len(self.datalist)-1:
                self.datalist[indx] = self.datalist[-1]
            del self.dataset[val]
            self.datalist.pop() 
            return True
        return False

    def getRandom(self) -> int:
        return random.choice(self.datalist)

if __name__ == "__main__":
    # ["RandomizedSet","insert","insert","remove","insert","remove","getRandom"]
    # [[],[0],[1],[0],[2],[1],[]]
    rs = RandomizedSet()
    print(rs.insert(0))
    print(rs.insert(1))
    print(rs.remove(0))
    print(rs.insert(2))
    print(rs.remove(1))
    print(rs.getRandom())