package main

import "fmt"

func main() {
	p1 := Player{
		playerName: "Gary Smith",
		eMail:      "GarySmith@neit.edu",
		balance:    15.00,
	}

	p2 := Player{
		playerName: "Laila Kerniech",
		eMail:      "LailaKerniech@neit.edu",
		balance:    30.00,
	}

	table := make(map[string]Player)
	table[p1.hash()] = p1
	table[p2.hash()] = p2

	fmt.Println("This is Gary's email using the hash: ", table[p1.hash()].eMail)
}

// Player contains all fields for a sample player in our hashing routine
type Player struct {
	playerName string
	eMail      string
	balance    float64
}

func (p *Player) hash() string {
	key := ""
	for i, char := range p.playerName {
		key += string(char + rune(i))
	}
	return key
}
