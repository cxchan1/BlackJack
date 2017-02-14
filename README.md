# BlackJack
*Simple RESTful HTTP API that uses JSON messages for Backend Blackjack game*

### Note: There are currently problems with the current version, so please reports any bugs to cxchan@ualberta.ca

## Ingredients

- [Ubuntu 14.04 "trusty" LTS 64bit base image](http://www.ubuntu.com/)
- [Go(lang) 1.7.0 or less](http://golang.org/)
- [Vim](http://www.vim.org/)
- [github.com/nleskiw/goplaycards/deck](https://github.com/nleskiw/goplaycards), providing a standard deck of playing cards (52 cards / 4 Suits / 2 through Ace)
- [github.com/gin-gonic/gin](https://gin-gonic.github.io/gin/), Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin..

## Optional ingredients

- [github.com/kardianos/govendor](https://github.com/kardianos/govendor), Go vendor tool that works with the standard vendor file.

## Setup and Usage

#### Clone the github repository:

```bash
git clone git@github.com:cxchan1/BlackJack.git
cd BlakcJack
go run main.go
```

### Start the Game
```
POST /blackjack21/start
```

Params:
* `bet_amount: { string }`

Example:

`curl`
```
curl localhost:8000/blackjack21/start -d 'bet_amount=10'
```

`response`
```json
{
  "message": "You have 100.00 left in your wallet.You bet: 10 Player Hand:   8♦    J♥   => Total Player Hand: 18,  Dealer Hand: XX   3♦  , Hit or Stand?"
}
```

### Hit or Stand
```
POST /blackjack21/action
```

Params:
* `action: { string }`

Example:

`curl`
```
curl localhost:8000/blackjack21/action -d 'action=Hit'

or

curl localhost:8000/blackjack21/action -d 'action=Stand'
```

`response`
```json
{
  "message": " Player Hand:   8♦    J♥    A♥   => Total Player Hand: 19,  Dealer Hand: XX   3♦  , Hit or Stand? Player Hand:   8♦    J♥    A♥    K♣   => Total Player Hand: 29,  -> Bust"
}

or

{
  "message": " Player Hand:   4♥    5♦   => Total Player Hand: 9,  Dealer Hand: XX   3♥  , Hit or Stand? Dealer Hand:   8♣    3♥    J♦   => Total Dealer Hand: 21  -> Dealer wins. Player loses."
}
```

#### Add Fund (just in case if you are running out of money)
```
POST /blackjack21/addFund
```

Params:
* `amount: { string }`

Example:

`curl`
```
curl localhost:8000/blackjack21/addFund -d 'amount=100'
```

`response`
```json
{
  "message": "Fund added it. You now have 180.00 in your wallet."
}
```

## How to play
- Start The game
- choose Hit or Stand as your action -> Hit can be calling mutliple of time until you lose or you win.
- Watch the outcome
- Repeat
- Occasionly if you run out of fund, don't forget to add more.


## Known issues and To DO list

- Param for the Hit or Stand apis have to be exacly "Hit" or "Stand". Otherwise It will crash the game.
- Add split and double down methods.
- Need to do more Unit Testing 
