class Solution:
    def myPow(self, x: float, n: int) -> float:
        ans=1
        if n>0:
            for i in range(n):
                ans=ans*x
        elif n<0:
            dem=1
            for i in range(-n):
                dem=dem*x
            ans=1/dem
        else:
            ans=1
        return ans



if __name__=='__main__':
    sol = Solution()
    print(sol.myPow(2,5))
    print(sol.myPow(2,-2))
    print(sol.myPow(2,0))
    print(sol.myPow(0.00001,2147483647))