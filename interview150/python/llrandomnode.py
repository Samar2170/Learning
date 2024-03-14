# Definition for singly-linked list.
from typing import Optional
import random

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:

    def __init__(self, head: Optional[ListNode]):
        self.head = head
        self.len = 0
        ph = head
        while ph.next:
            self.len+=1        
            ph=ph.next

    def getRandom(self) -> int:
        randInt = random.randint(0,self.len)
        i=0
        ph = self.head
        while i<randInt:
            ph=ph.next
            i+=1
        return ph


if __name__=="__main__":
#     "Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
# [[[1, 2, 3]], [], [], [], [], []]
    head = ListNode(1,ListNode(2,ListNode(3)))
    s = Solution(head)
    print(s.getRandom())
    print(s.getRandom())
    print(s.getRandom())
