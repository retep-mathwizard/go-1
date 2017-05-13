package ai

import ()

type Race struct {
	Name    string
	Health  int
	Gold    int
	About   string
	Level   int
	Shield  int
	Attacks []Move
}

type Move struct {
	Name     string
	Info     string
	Damage   int
	Accuracy int
}

func GenStats(opponent *load.Race) Race {
	name = "Bowser Jr."

}
