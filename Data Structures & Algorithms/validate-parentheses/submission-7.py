class Solution:
	def isOpenBracket(self, c: str) -> bool:
		if c == '{' or c == '(' or c == '[':
			return True
		return False
	def isCorrectPair(self, c1: str, c2: str) -> bool:
		match c1:
			case '{':
				if c2 == '}':
					return True
			case '(':
				if c2 == ')':
					return True
			case '[':
				if c2 == ']':
					return True
			case _:
				return False
		return False
	
	def isValid(self, s: str) -> bool:
		stack = []

		for c in s:
			if self.isOpenBracket(c):
				stack.append(c)
			else:
				if not stack:
					return False
				top = stack.pop()
				if not self.isCorrectPair(top, c):
					return False

		return len(stack) == 0