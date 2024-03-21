package main

import "fmt"

// Team is a struct that contains the name of the team and the scores of the team.
type Team struct {
	name   string
	scores []float32
}

// GetAverage is a method of Team struct that returns the average score of the team.
func (t Team) GetAverage() (average float32) {
	var count float32
	for _, score := range t.scores {
		count += score
	}
	average = count / float32(len(t.scores))
	return
}

// Match is a struct that contains two teams.
type Match struct {
	teams [2]Team
}

// CheckWinner is a method of Match struct that returns the winner of the match.
// The winner is determined by the average score of each team.
// If the average score of team 1 is equal to team 2, then it's a draw.
// Score average is considered valid if it's greater than or equal to 100.
func (m Match) CheckWinner() string {
	team1Average := m.teams[0].GetAverage()
	team2Average := m.teams[1].GetAverage()

	// Checking the winner
	if team1Average > team2Average { // Check is Team 1 is Win
		if validate(team1Average) {
			return fmt.Sprintf("Pemenang\t: %s\nScore Average\t: %.2f", m.teams[0].name, m.teams[0].GetAverage())
		}

	} else if team1Average < team2Average { // Check is Team 2 is Win
		if validate(team2Average) {
			return fmt.Sprintf("Pemenang\t: %s\nScore Average\t: %.2f", m.teams[1].name, m.teams[1].GetAverage())
		}

	} else { // Check is Draw
		if validate(team1Average) && validate(team2Average) {
			return "Draw"
		}
	}

	return "No Winner"
}

// validate is a function that checks if the average score is valid (greter thean equal 100).
func validate(avgScore float32) bool {
	return avgScore >= 100
}

func main() {
	data1 := Match{
		teams: [2]Team{
			{
				name:   "Lumba-lumba",
				scores: []float32{96, 108, 89},
			},
			{
				name:   "Koala",
				scores: []float32{88, 91, 110},
			},
		},
	}
	fmt.Println("Data 1")
	fmt.Println(data1.CheckWinner())

	dataBonus1 := Match{
		teams: [2]Team{
			{
				name:   "Lumba-lumba",
				scores: []float32{97, 112, 101},
			},
			{
				name:   "Koala",
				scores: []float32{109, 95, 123},
			},
		},
	}
	fmt.Println("\nData Bonus 1")
	fmt.Println(dataBonus1.CheckWinner())

	dataBonus2 := Match{
		teams: [2]Team{
			{
				name:   "Lumba-lumba",
				scores: []float32{97, 112, 101},
			},
			{
				name:   "Koala",
				scores: []float32{109, 95, 106},
			},
		},
	}
	fmt.Println("\nData Bonus 2")
	fmt.Println(dataBonus2.CheckWinner())

}
