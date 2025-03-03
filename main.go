package main

import (
	"fmt"

	"github.com/betamemo/go-cinema/movie"
	"github.com/betamemo/go-cinema/tickets"
)

func main() {
	movieName := movie.FindName("tt4154796")
	fmt.Println(movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
