package techpalace
import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	name := strings.ToUpper(customer)
	var result string
	result = fmt.Sprintf("Welcome to the Tech Palace, %v", name)
	return result
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	result := border + "\n" + welcomeMsg + "\n" + border
	return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	clean := strings.Map(func (r rune)rune{
				if r == 42 {
					return -1
				}
				return r	
	}, oldMsg)
	slice := strings.Fields(clean)
	return strings.Join(slice, " ")
}
