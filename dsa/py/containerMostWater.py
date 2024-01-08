class Solution:
    def maxArea(self, height) -> int:
        n=len(height)
        i,j=0,n-1
        maxArea=0
        while i<j:
            areaHere = min(height[i], height[j]) * (j-i)
            if areaHere > maxArea:
                maxArea = areaHere
            if height[i]>height[j]:
                j-=1
            else:
                i+=1
                
        return maxArea
            
        
    
if __name__ =='__main__':
    n=[1,8,6,2,5,4,8,3,7]
    s=Solution()
    print(s.maxArea(n))
    