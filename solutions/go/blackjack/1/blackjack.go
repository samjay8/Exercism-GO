package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	default:
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var result string

	val1 := ParseCard(card1)
	val2 := ParseCard(card2)
	val3 := ParseCard(dealerCard)

	score := val1 + val2

	if score <= 11 {
		result = "H"
	}
	if score == 22 {
		result = "P"
	}
	if score == 21 {
		switch dealerCard {
		case "ten", "jack", "queen", "king", "ace":
			result = "S"
		default:
			result = "W"

		}
	}
	if score >= 17 && score <= 20 {
		result = "S"
	}
	if score >= 12 && score <= 16 {
		if val3 >= 7 {
			result = "H"
		} else {
			result = "S"
		}
	}
	return result
}
