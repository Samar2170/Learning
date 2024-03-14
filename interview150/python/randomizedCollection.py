
class Holder:
    def __init__(self, val, indx):
        self.val = val
        self.freq = 1
        self.indices = [indx]
    
    def add(self, indx):
        self.indices.append(indx)
        self.freq+=1

    def remove(self):
        self.indices.pop()
        self.freq-=1

    def get_last_index(self):
        return self.indices[-1]

class RandomizedCollection:

    def __init__(self):
        self.collection={}
        self.datalist = []

    def insert(self, val: int) -> bool:
        if val in self.collection:
            self.collection[val].add(len(self.datalist))
            self.datalist.append(val)
            return False
        self.collection[val]=Holder(val, len(self.datalist))
        self.datalist.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val in self.collection:
            if self.collection[val].freq==1:
                indx=self.collection[val].get_last_index()
                del self.collection[val]
                if indx!=len(self.datalist)-1:
                    self.collection
            else:
                self.collection[val].remove()
            return True
        return False

    def getRandom(self) -> int:
        pass        
