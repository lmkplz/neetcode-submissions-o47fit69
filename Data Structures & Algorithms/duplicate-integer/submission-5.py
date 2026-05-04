class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        hash = {}
        for num in nums:
            if hash.get(num, 0) > 0:
                return True
            else:
                hash[num] = hash.get(num, 0) + 1
        return False