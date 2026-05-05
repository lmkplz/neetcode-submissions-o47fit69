class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max = 0
        curPrice = prices[0]

        for price in prices:
            if curPrice < price:
                margin = price - curPrice
                if max < margin:
                    max = margin
            else:
                curPrice = price

        return max