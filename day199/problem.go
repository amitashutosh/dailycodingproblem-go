package day199

import "strings"

// BalanceParens balances unbalanced parens.
// Runs in O(N) time and O(N) extra space.
// Only accepts a string composed of '(' and ')'.
func BalanceParens(parens string) string {
	var sb strings.Builder
	opens := 0
	for _, r := range parens {
		if r == '(' {
			opens++
		} else if r == ')' && opens == 0 {
			sb.WriteRune('(')
		} else {
			opens--
		}
		sb.WriteRune(r)
	}
	for i := 0; i < opens; i++ {
		sb.WriteRune(')')
	}
	return sb.String()
}
