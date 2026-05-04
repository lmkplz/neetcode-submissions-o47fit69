class Solution:
    def isPalindrome(self, s: str) -> bool:
        newStr = "".join([c.lower() for c in s if c.isalnum()])

        n = len(newStr) // 2
        for i in range(n):
            if newStr[i] != newStr[len(newStr)-1-i]:
                return False
        
        return True