package euchre

// Player represents a single player.
type Player string

// Team represents a two person team.
type Team struct {
	Name    string
	Players []Player
}
