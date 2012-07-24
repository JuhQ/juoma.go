/*
 * This is a simple drinking game. Juoma means "a drink" in Finnish.
 * Rules are simple:
 *   - List players
 *   - List available drinks
 *   - The person with the highest number has to drink the drink provided by the game,
 *      player can decide portion size
 *   - If you puke or pass out, you lose
 *   - If a drink runs out, it must be removed from the drink list
 *   - Last one who's not thrown up or passed out, IS THE WINNER!
 *   - Do not die
 *   - Have fun
 *
 *  Juha Tauriainen @juha_tauriainen juha@bin.fi
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// list players here
	players := []string{"Juha", "Pekka", "Timo", "Liisa", "Maija"}

	// list drinks here
	drinks := []string{"Beer", "Koskenkorva", "Whiskey", "Milk"}

	winnerStr := ""
	winner := 0

	rand.Seed(time.Now().UnixNano())
	for i := range players {
		roll := rand.Intn(100)
		string := fmt.Sprintf("%s rolls %d for drink %s", players[i], roll, drinks[rand.Intn(len(drinks))])

		if roll > winner {
			winner = roll
			winnerStr = string
		}
		fmt.Println(string)
	}

	fmt.Println("WINNER:", winnerStr)
}
