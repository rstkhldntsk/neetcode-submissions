func generateParenthesis(n int) []string {
	result := make([]string, 0)

	var gen func(prefix string, open, closed int)
	gen = func(prefix string, open, closed int) {
		if open == n && closed == n {
			result = append(result, prefix)
			return
		}
		if open < n {
			gen(prefix+"(", open+1, closed)
		}
		if closed < open {
			gen(prefix+")", open, closed+1)
		}
	}
	gen("", 0 ,0)

	return result
}
