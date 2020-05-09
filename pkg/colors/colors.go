package colors

import "strconv"

const (
	// DarkGrayFG is the ANSI code for dark gray foregrounds
	DarkGrayFG int = 90

	// CyanFG is the ANSI code for cyan foregrounds
	CyanFG int = 36

	// RedFG is the ANSI code for red foregrounds
	RedFG int = 31

	// YellowFG is the ANSI code for orange foregrounds
	YellowFG int = 33
)

// ColorText will return text with a given ANSI code
func ColorText(code int, text string) string {
	return "\033[1;" + strconv.Itoa(code) + "m" + text + "\033[0m"
}
