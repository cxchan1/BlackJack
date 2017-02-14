package api

import (
	"fmt"
	"strconv"

	"github.com/BlackJack/error"
	"github.com/nleskiw/goplaycards/deck"
)

// getString gets an arbitrary string from the user with a prompt.
func getString(prompt string) string {
        fmt.Print(prompt)
        var input string
        fmt.Scanln(&input)
        return input
}

func getInteger(amount string) (m int, e error) {
	integer, err := strconv.Atoi(amount)
	if err != nil {
		e = errs.Client{fmt.Sprintf("Can't convert your bet amount into an integer.")}
		return
	}
	m = integer
	return
}

func getBet(wallet *float64, amount string) (m int, e error) {
	valid := false
	for valid == false {
		valid = true
		bet, error := getInteger(amount)
		if error != nil {
			e = error
			return
		}
		if bet < 5 {
			valid = false
		}
		if float64(bet) > *wallet {
			valid = false
		}
		if bet%5 != 0 {
			valid = false
		}
		if valid == false {
			e = errs.Client{fmt.Sprintf("Invalid bet or you don't have enough moneys.")}
			return
		}
		m = bet
	}
	*wallet = *wallet - float64(m)
	return
}

// handTotal returns the numerical value of a Blackjack hand
func handTotal(hand []deck.Card) int {
	total := 0
	numberOfAces := 0
	for _, card := range hand {
		if card.Value.Name == "Ace" {
			numberOfAces = numberOfAces + 1
		} else {
			if card.Facecard() {
				total = total + 10
			} else {
				total = total + card.Value.Value
			}
		}
	}

	// If there's at least one Ace, deal with it.
	if numberOfAces > 0 {

		for numberOfAces > 1 {
			total = total + 1
			numberOfAces = numberOfAces - 1
		}
		if total+11 > 21 {
			total = total + 1
		} else {
			// If 11 doesn't cause a bust, make it worth 11
			total = total + 11
		}
	}
	return total
}

// Returns true if a hand is bust / over 21
func isBust(hand []deck.Card) bool {
	if handTotal(hand) > 21 {
		return true
	}
	return false
}

// A Blackjack is exactly one Ace and Exactly one 10, K, Q, or A
func isBlackjack(hand []deck.Card) bool {
	if len(hand) != 2 {
		return false
	}
	if hand[0].Value.Name == "Ace" {
		if hand[1].Value.Value >= 10 && hand[1].Value.Value <= 13 {
			return true
		}
	}
	if hand[1].Value.Name == "Ace" {
		if hand[0].Value.Value >= 10 && hand[0].Value.Value <= 13 {
			return true
		}
	}
	return false
}

func printPlayerHand(hand []deck.Card) string {
	m := fmt.Sprintf(" Player Hand: ")
	for _, card := range hand {
		m = m + fmt.Sprintf(" %s  ", card.ToStr())
	}
	m = m + fmt.Sprintf(" => Total Player Hand: %d, ", handTotal(hand))
	return m
}

func printDealerHand(hand []deck.Card, hideFirst bool) string {
	m := fmt.Sprintf(" Dealer Hand: ")
	if hideFirst {
		m = m + fmt.Sprintf("XX  %s  ", hand[1].ToStr()) + ", Hit or Stand?"
		return m
	} else {
		for _, card := range hand {
			m = m + fmt.Sprintf(" %s  ", card.ToStr())
		}
		m = m + fmt.Sprintf(" => Total Dealer Hand: %d ", handTotal(hand))
		return m
	}
}

func GetBalance(wallet *float64) (m float64, e error) {
	if *wallet < 5.0 {
		e = errs.Client{fmt.Sprintf("You're out of money. Please add more fund")}
		return
	}
	m = *wallet
	return
}

func StartBet(wallet *float64, d *deck.Deck, amount string) (p []deck.Card, v []deck.Card, m string, b int, e error) {
		// Minimum number of cards to play a hand from a single deck in worst case
		if d.CardsLeft() < 17 {
			d.Initialize()
			d.Shuffle()
		}
		_, err := GetBalance(wallet)
		if err != nil {
			e = err
			return
		}
		m = fmt.Sprintf("You have %.2f left in your wallet.", *wallet)
		bet, err := getBet(wallet, amount)
		if err != nil {
			e = err
			return
		}
		a := fmt.Sprintf("You bet: %d", bet)
		m = m + a

		// draw the initial hands
		playerHand, err := d.Draw(2)
		if err != nil {
			e = err
			return
		}
		dealerHand, err := d.Draw(2)
		if err != nil {
			e = err
			return
		}
		p = playerHand
		v = dealerHand
		b = bet

		// Quick check
		if isBlackjack(dealerHand) || isBlackjack(playerHand) {
			res1 := printPlayerHand(playerHand)
			m = m + res1
			res2 := printDealerHand(dealerHand, false)
			m = m + res2

			if isBlackjack(dealerHand) && isBlackjack(playerHand) {
				m = m + " -> Both the Player and the Dealer have Blackjack. Hand is a push."
				*wallet = *wallet + float64(bet)
				return
			}

			if isBlackjack(dealerHand) {
				m = m + " -> Dealer has Blackjack. Player loses this hand."
				return
			}

			if isBlackjack(playerHand) {
				winnings := float64(bet) * 2.5
				m = m + " -> Player has Blackjack. Player wins."
				*wallet = *wallet + winnings
				return
			}
		} else {
			res1 := printPlayerHand(playerHand)
			m = m + res1
			res2 := printDealerHand(dealerHand, true)
			m = m + res2
		}
		return
}

func Action(wallet *float64, d *deck.Deck, hand []deck.Card, hand2 []deck.Card, bet int, action string) (p []deck.Card, m string, e error) {
	res1 := printPlayerHand(hand)
	m = m + res1
	res2 := printDealerHand(hand2, true)
	m = m + res2
	if action == "Hit" {
		drawnCards, err := d.Draw(1)
		if err != nil {
			e = err
			return
		}
		hand = append(hand, drawnCards[0])
		p = hand
		if isBust(hand) {
			res1 = printPlayerHand(hand)
			m = m + res1 + " -> Bust"
		}
	}	else if action == "Stand" {
			for handTotal(hand2) < 17 {
				drawnCards, err := d.Draw(1)
				if err != nil {
					e = err
					return
				}
				hand2 = append(hand2, drawnCards[0])
			}
			res2 = printDealerHand(hand2, false)
			m = m + res2
			if handTotal(hand2) > 21 {
				m = m + " -> Dealer busts.  Player wins."
				*wallet = *wallet + float64(bet*2)
				return
			}

			if handTotal(hand) > handTotal(hand2) {
				m = m + " -> Player's hand beats the dealer. Player wins."
				*wallet = *wallet + float64(bet*2)
				return
			}
			if handTotal(hand) == handTotal(hand2) {
				m = m + " -> Push."
				*wallet = *wallet + float64(bet)
				return
			}
			if handTotal(hand) < handTotal(hand2) {
				m = m + " -> Dealer wins. Player loses."
				return
			}
		} else {
			e = errs.Client{fmt.Sprintf("This action is invalid, Please choose only 'Hit' or 'Stand'.")}
			return
		}
	return
}

func AddFund(wallet *float64, d *deck.Deck, amount string) (m string, e error) {
	fund, error := getInteger(amount)
	if error != nil {
		e = error
		return
	}
	*wallet = *wallet + float64(fund)
	m = fmt.Sprintf("Fund added it. You now have %.2f in your wallet.", *wallet)
	return
}
