class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        i,j=0,0
        result = 0
        charmap=set()
        while j<len(s):
            while s[j] in charmap:
                charmap.remove(s[i])
                i+=1
            charmap.add(s[j])
            result = max(result, j-i+1)
            j+=1
            
        return result