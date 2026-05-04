class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        hash = {}
        for num1 in nums:
            if hash.get(num1, 0) > 0:
                return True
            else:
                hash[num1] = hash.get(num1, 0) + 1
        return False