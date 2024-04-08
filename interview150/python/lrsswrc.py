def lengthOfLongestSubstring(input):
    l,r=0,0 
    cmap={}
    maxLen=0
    while r<len(input):
        if input[r] in cmap:
            del cmap[input[r]]
            l+=1
        else:
            cmap[input[r]]=1
            r+=1
        cs=input[l:r]
        if len(cs)>maxLen:
            maxLen=len(cs)
    return maxLen

def test_lengthOfLongestSubstring():
    # assert lengthOfLongestSubstring("abcabcbb")==3
    # assert lengthOfLongestSubstring("bbbbb")==1
    # # assert lengthOfLongestSubstring("pwwkew")==3
    # assert lengthOfLongestSubstring(" ")==1
    # assert lengthOfLongestSubstring("au")==2
    # assert lengthOfLongestSubstring("aab")==2
    # assert lengthOfLongestSubstring("dvdf")==3
    # assert lengthOfLongestSubstring("abba")==2
    print(lengthOfLongestSubstring("pwwkew"))

test_lengthOfLongestSubstring()

