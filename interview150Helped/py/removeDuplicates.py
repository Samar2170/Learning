class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        unique = set()
        remove_indices = []
        indx = len(nums)-1
        while indx >=0:
            if nums[indx] in unique:
                remove_indices.append(indx)
            else:
                unique.add(nums[indx])
            indx=indx-1
        
        for r in remove_indices:
            nums.pop(r)
            
        return len(nums)
        
        