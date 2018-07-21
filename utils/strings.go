package utils

import "strings"

// FixBackQuotes Replace \%\*\*\*\% to \`\`\` and \%\*\% to \`
func FixBackQuotes(content string) string {
	content = strings.Replace(content, "%***%", "```", -1)
	content = strings.Replace(content, "%*%", "`", -1)

	return content
}
