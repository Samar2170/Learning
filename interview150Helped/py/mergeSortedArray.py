from typing import List
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        i,j=0,0
        while i<m or j<n:
            if j<n and i<m:
                if nums2[j] < nums1[i]:
                    nums1[i+1:m-1] = nums1[i:m]
                    nums1[i] = nums2[j]

                    j+=1
                    m+=1
                else:
                    i+=1
            else:
                if j<n:
                    nums1[i] = nums2[j]
                    j+=1
                    i+=1
                else:
                    break


        print(nums1)




if __name__ == "__main__":
    nums1 = [1,2,3,0,0,0]
    m=3
    nums2 = [2,5,6]
    n=3
    s = Solution()
    s.merge(nums1,m,nums2,n)