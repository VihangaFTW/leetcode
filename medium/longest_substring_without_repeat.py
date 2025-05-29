class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        
        
        The sliding window defines a valid substring that has no repeating characters at current position
        
        For each position:
        
        1.we adjust the window to ensure it remains valid (no repeats)
        2. we check if the current window is larger than our previous maximum
        3. we update the maximum if needed
        
        Note: the sliding window DOES NOT keep track of the length of the current biggest substring. We need to manually keep track.
        """

        # dictionary to store the last seen index for a character
        last_seen_indices: dict[str, int]  = {}
        # start position of sliding window [start, end]
        
        start:int = 0
        # max_len tracks the max length of any valid substring we have seen so far
        max_len: int = 0

        for end in range(len(s)):
            #? character exists in current window
            if s[end] in last_seen_indices and last_seen_indices[s[end]] >= start:
                # we have seen this character before, so we need to push the start of the window to the position after previous occurence of this character to maintain the window substring criteria
                start: int = last_seen_indices[s[end]] + 1
            
            
            #? new character but still check whether we got new max length
            new_window_len: int = end- start + 1
            # keeping track of the length of the current biggest substring
            max_len = max(new_window_len, max_len)
            last_seen_indices[s[end]] = end
        
        return max_len
        


if __name__  == '__main__':
    #*    0123456789
    s  = 'abcdbcabc'
    x = Solution()
    print(x.lengthOfLongestSubstring(s))