from typing import List
import ipdb  # noqa
class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        ipdb.set_trace()
        res = []

        def backtrack(start,comb):
            ipdb.set_trace()
            if len(comb)==k:
                res.append(comb.copy())
                return
            for i in range(start,n+1):
                comb.append(i)
                backtrack(i+1, comb)
                comb.pop()

        backtrack(1,[])
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.combine(4,2))
