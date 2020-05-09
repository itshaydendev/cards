package logger

import (
	"fmt"

	"github.com/itshaydendev/cards/pkg/colors"
)

// Info is a utility function to display informational text to the console
func Info(message string) {
	fmt.Println(
		colors.ColorText(colors.CyanFG, "info") +
			colors.ColorText(colors.DarkGrayFG, " => ") +
			message,
	)
}

// Error is a utility function to display error text to the console
func Error(message string) {
	fmt.Println(
		colors.ColorText(colors.RedFG, "error") +
			colors.ColorText(colors.DarkGrayFG, " => ") +
			message,
	)
}

// Warn is a utility function to display warning text to the console
func Warn(message string) {
	fmt.Println(
		colors.ColorText(colors.YellowFG, "info") +
			colors.ColorText(colors.DarkGrayFG, " => ") +
			message,
	)
}
