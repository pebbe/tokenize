package alpino

import "strings"

func Escape(s string, fullbrackets bool, first bool) string {
	lines := strings.Split(s, "\n")
	found1 := false
	for i, line := range lines {
		tokens := strings.Fields(line)
		found2 := false
		for j, token := range tokens {
			if token == "[" || token == "]" || (fullbrackets && (token == "\\[" || token == "\\]")) {
				tokens[j] = "\\" + token
				found2 = true
			}
		}
		if found2 {
			lines[i] = strings.Join(tokens, " ")
			found1 = true
		}
		if first {
			if strings.HasPrefix(lines[i], "%") || strings.HasPrefix(lines[i], "|") {
				lines[i] = "|" + lines[i]
				found1 = true
			}
		}
	}
	if found1 {
		return strings.Join(lines, "\n")
	}
	return s
}
