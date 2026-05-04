class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hash = {}
        for key, val in enumerate(nums):
            if val not in hash:
                hash[val] = []
            hash[val].append(key)
        print(hash)
        for key, val in hash.items():
            n = target - key
            if n in hash:
                if n == key:
                    if len(val) > 1:
                        return [val[0], val[1]]
                    else:
                        continue
                return [val[0], hash[n][0]]