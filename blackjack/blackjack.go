package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	totalValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	if card1 == "ace" && card2 == "ace" {
		return "P"
	} else if totalValue == 11 {
		return "H"
	} else if totalValue > 16 && totalValue < 21 {
		return "S"
	} else if totalValue > 11 && totalValue < 17 {
		if ParseCard(dealerCard) >= 7 {
			return "H"
		} else {
			return "S"
		}
	} else if totalValue == 21 && dealerValue != 11 && dealerValue != 10 {
		return "W"
	} else if totalValue == 21 {
		return "S"
	}
	return "H"
}
