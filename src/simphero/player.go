package player

type Player struct {
	var name string
	var health int
	var abilities [][]interface{}
	//[["fireball",23, .6],["lightning",17, .8]]
	var isAi bool
	var agility int
}
func (player *Player) Attack(opponentHealth int) int, bool{
	if player.isAi {
		move := random.Choice(abilities)
		if  rand.Float64() <= move[3] {

			return opponentHealth -= move[2], true
		}
		return opponentHealth, false
	}
	Pmove := GameController.ChooseMove(player.abilities)
	if rand.Float64 <= Pmove[3] {
		return opponentHealth -= move[2], true
	}
	return opponentHealth, false
}
type GameController struct {
	WhoGoesFirst string
}

func (controller *GameController) ChooseMove(abilities [][]interface{}) {
	x := 1
	for _, movePair := range abilities {
		fmt.Printf("[%v]%v (%v damage, %v % accuracy)",x,movePair[0], movePair[1], movePair[2] * 100)
		x += 1
	}
	response := strconv.Atoi(input.Prompt("Enter number of desired move: > "))
	return abilities[response]

}
func (controller *GameController) PrintTotals(ability []interface{}, health int, hasHit bool) {
	if hasHit {
		fmt.Println(ability[1] health)
	}


//*Player is an instance of pointer
	
func (player
player
-health
-abilities
-func attack
-ai bool

ability
-name
-damage

controller
-func StartGame
-func WhoGoesFirst
-func PrintTotals
*/
