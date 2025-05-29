class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        Prerequisite: nums1 and nums2 are sorted in ascending order
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        :time complexity: O(log(|nums1|, |nums2|))
        """

        # perform binary search on the smaller array for efficiency; time: O(log (nums1))
        if len(nums2) < len(nums1):
            nums1, nums2 = nums2, nums1
        
        x, y = len(nums1), len(nums2)
        start, end  = 0, x
        
        while start<=end:
            # calculate partitioning pivots for both arrays
            pivotX = (start + end)//2
            pivotY = (x+y+1)//2 - pivotX
            
            # compute edge values in partitions from both arrays
            # handle edge cases where a partition might not have any elements
            maxLeftX =  float('-inf') if pivotX == 0 else nums1[pivotX-1]
            minRightX = float('inf') if pivotX == x else nums1[pivotX]
            
            maxLeftY = float('-inf') if pivotY == 0 else nums2[pivotY-1]
            minRightY = float('inf') if pivotY == y else nums2[pivotY]
            
            # binary search to the left: median value closer to X's left partition
            if maxLeftX > minRightY:
                end = pivotX-1
            # binary search to the right: median value closer to X's right partition
            elif maxLeftY > minRightX:
                start = pivotX + 1
            # each partition size sum up to half of total size; so median can be calculated from the edge values of current paritioning
            else:
                # total size is even: median is average of middle two elements
                if (x+y)%2 == 0:
                    return (max(maxLeftX, maxLeftY) + min(minRightX, minRightY))/2.0
                else:
                    return max(maxLeftX, maxLeftY)


if __name__ == '__main__':
    x = Solution()
    nums1 = [1,3]
    nums2 = [2]
    print(x.findMedianSortedArrays(nums1, nums2))