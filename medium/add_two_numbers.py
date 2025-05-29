class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        
    def _append(self, value):
        current = self
        while current.next:
            current = current.next
            
        current.next = ListNode(value)
        
    def append_list(self, vals: list[int]):
        for val in vals:
            self._append(val)
    
    def __repr__(self)-> str:
        output = ''
        current = self
        while current:
            output += str(current.val)
            current = current.next 
            
        return output

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: Optional[ListNode]
        :type l2: Optional[ListNode]
        :rtype: Optional[ListNode]
        """
        
        dummy_head = ListNode(0)
        current = dummy_head
        
        carry = 0
        
        # Continue loop while any of the lists has nodes or there's a carry
        while l1 or l2 or carry:
            # Get values (or 0 if node is None)
            x = l1.val if l1 else 0
            y = l2.val if l2 else 0
            
            # Calculate sum and carry
            total = x + y + carry
            carry = total // 10
            
            # dd new node to output
            current.next = ListNode(total % 10)
            current = current.next
            
            # advance pointers
            if l1:
                l1 = l1.next
            if l2:
                l2 = l2.next
        
        return dummy_head.next
        
if __name__ == "__main__":
    l1 = ListNode(2)
    l1.append_list([4, 3])
    
    l2 = ListNode(5)
    l2.append_list([6, 4])
    
    x = Solution()
    result = x.addTwoNumbers(l1, l2)
    print(result)