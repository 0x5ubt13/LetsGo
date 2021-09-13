package cards

import (
	"fmt"
)

func main() {
	// Check whether the file exists:
	/* if _, err := os.Stat("./my_cards"); err == nil {
		cards := newDeckFromFile("my_cards")
		fmt.Println("Exists")
	} else if os.IsNotExist(err) {
		cards := newDeck()
		fmt.Println("Doesn't exost")
	} else {
		fmt.Println("Fucking hell")}
	*/
	cards := newDeck()
	cards.shuffle()
	cards.print()
	fmt.Println("There are", len(cards), "cards")
	cards.saveToFile("my_cards")
}
