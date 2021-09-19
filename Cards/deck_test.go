package cards

import "testing"

func testNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of Ace of Hearts, but got %v instead.", d[0])
	}

	if d[len(d) - 1] != "King of Diamonds" {
		t.Errorf("Expected last card of King of Diamonds, but got %v instead.", d[len(d) -1])
	}
	

}
