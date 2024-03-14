from typing import List
import random
import copy

class Solution:

    def __init__(self, nums: List[int]):
        self.nums=copy.deepcopy(nums)
        self.shuffled=copy.deepcopy(nums)

    def reset(self) -> List[int]:
        self.shuffled = copy.deepcopy(self.nums)
        return self.nums

    def shuffle(self) -> List[int]:
        size = len(self.shuffled)
        if size%2==0:
            middle = size/2
        else:
            middle = int((size)/2)+1
        i=0
        j=0
        while i < middle:
            while j==i:
                j=random.randint(0,size-1)
            self.shuffled[i], self.shuffled[j]=self.shuffled[j], self.shuffled[i]
            i+=1
        return self.shuffled

if __name__=="__main__":
    nums = [1,2,3]
    s = Solution(nums)
    print(s.shuffle())
    print(s.reset())
    print(s.shuffle())


# shallow copy vs deep copy 
## https://www.geeksforgeeks.org/copy-python-deep-copy-shallow-copy/
## https://www.programiz.com/python-programming/shallow-deep-copy
## https://www.geeksforgeeks.org/copy-python-deep-copy-shallow-copy/