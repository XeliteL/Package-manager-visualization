package dots

import "strings"

// Удаление/замена несовместимых с Graphviz символов
func SanitizeName(name string) string {
	replacer := strings.NewReplacer(
		"<", "",
		">", "",
		",", "",
		"\"", "",
		"=", "_",
	)
	return replacer.Replace(name)
}
