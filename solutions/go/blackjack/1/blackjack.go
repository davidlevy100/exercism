// Package blackjack implements basic blackjack scoring and first-turn strategy.
package blackjack

// ParseCard returns the blackjack value of a card.
// Ace = 11, 2–9 = face value, 10/face cards = 10, others = 0.
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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the first-turn decision ("P", "W", "S", "H")
// based on two player cards and the dealer's card.
func FirstTurn(card1, card2, dealerCard string) string {
	v1, v2 := ParseCard(card1), ParseCard(card2)
	sum := v1 + v2
	dealer := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P" // split aces
	case sum == 21 && (dealer == 10 || dealer == 11):
		return "S" // blackjack vs 10 or ace → stand
	case sum == 21:
		return "W" // blackjack otherwise → win
	case sum >= 17 && sum <= 20:
		return "S" // 17–20 → stand
	case sum >= 12 && sum <= 16 && dealer >= 7:
		return "H" // 12–16 vs 7+ → hit
	case sum >= 12 && sum <= 16:
		return "S" // 12–16 vs 2–6 → stand
	case sum <= 11:
		return "H" // 11 or less → hit
	}
	return "S" // fallback
}
