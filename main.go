package main

import (
	"net/http"
	"github.com/BlackJack/api"
	"github.com/nleskiw/goplaycards/deck"
	"github.com/gin-gonic/gin"

)

func main() {
	wallet := 100.00
	var d deck.Deck
	d.Initialize()
	d.Shuffle()
	//defer dbmap.Db.Close()

	var playerHand []deck.Card
	var dealerHand []deck.Card
	var bet int
	for wallet > 5.0 {

		router := gin.Default()
		router.POST("/blackjack21/start", func(c *gin.Context) {
			p, v, message, b, err := api.StartBet(&wallet, &d, c.PostForm("bet_amount"))
			if err != nil {
				c.JSON(400, gin.H{"result": err})
			}
			playerHand = p
			dealerHand = v
			bet = b
    	c.JSON(http.StatusOK, gin.H{"message": message})})
		router.POST("/blackjack21/action", func(c *gin.Context) {
			p, message, err := api.Action(&wallet, &d, playerHand, dealerHand, bet, c.PostForm("action"))
			if err != nil {
				c.JSON(400, gin.H{"result": err})
			}
			playerHand = p
    	c.JSON(http.StatusOK, gin.H{"message": message})})
		router.POST("/blackjack21/addFund", func(c *gin.Context) {
			message, err := api.AddFund(&wallet, &d, c.PostForm("amount"))
			if err != nil {
				c.JSON(400, gin.H{"result": err})
			}
    	c.JSON(http.StatusOK, gin.H{"message": message})})
		router.Run(":8000")
	}
}
