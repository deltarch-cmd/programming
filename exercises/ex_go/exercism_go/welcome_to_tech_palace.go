package main
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	borderWelcome := border + "\n" + welcomeMsg + "\n" + border
	return borderWelcome
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleaned := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(cleaned)
}
