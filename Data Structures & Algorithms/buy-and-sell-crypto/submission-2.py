class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max = 0
        curPrice = prices[0]
        buyIndex = 0
        maxIndex = {}

        for key, price in enumerate(prices):
            if curPrice < price:
                margin = price - curPrice
                if max < margin:
                    max = margin
                    maxIndex = { buyIndex, key }
            else:
                curPrice = price
                buyIndex = key
		
        print(maxIndex)

        return max