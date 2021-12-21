package day21

import (
	"fmt"
	"log"

	"github.com/wbean1/AoC/utils"
)

type State struct {
	player1Position, player2Position int
	player1Score, player2Score       int
	dice, numRolls, rollValue        int
	turn                             int // 0 = player1's turn, 1 = player2's turn
}

func Run() {
	s := State{player1Position: 1, player2Position: 6}
	for !s.HasSomeoneWon() {
		s.TakeTurn()
	}
	losingScore, err := s.GetLosingScore()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part1: losing score * numRolls = %d\n", losingScore*s.numRolls)
	diracStates := []State{State{player1Position: 1, player2Position: 6}}
	diracStatesComplete, diracStatesIncomplete := SplitComplete(diracStates)
	for len(diracStatesIncomplete) > 0 {
		fmt.Printf("completeStates: %d, incompleteStates: %d\n", len(diracStatesComplete), len(diracStatesIncomplete))
		for _, d := range diracStatesIncomplete {
			diracStatesComplete = append(diracStatesComplete, TakeDiracTurn(d)...)
		}
		diracStatesComplete, diracStatesIncomplete = SplitComplete(diracStatesComplete)
	}
	player1Wins, player2Wins := CountWins(diracStatesComplete)
	fmt.Printf("part2: player1 wins: %d, player2 wins: %d, max: %d\n", player1Wins, player2Wins, utils.Max(player1Wins, player2Wins))
}

func CountWins(s []State) (int, int) {
	var p1wins, p2wins int
	for _, d := range s {
		if d.player1Score >= 1000 {
			p1wins++
		} else if d.player2Score >= 1000 {
			p2wins++
		} else {
			log.Fatal("shouldn't get here ever")
		}
	}
	return p1wins, p2wins
}

func TakeDiracTurn(s State) []State {
	states := []State{}
	for i := 1; i <= 3; i++ { // roll dice 3 times
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				s.rollValue = i + j + k
				states = append(states, s)
			}
		}
	}
	updatedStates := []State{}
	for _, d := range states {
		if d.turn == 0 { // player1's turn
			d.player1Position += d.rollValue
			d.player1Position = d.player1Position % 10
			if d.player1Position == 0 {
				d.player1Position = 10
			}
			d.player1Score += d.player1Position
			d.turn = 1
		} else { // player2's turn
			d.player2Position += d.rollValue
			d.player2Position = d.player2Position % 10
			if d.player2Position == 0 {
				d.player2Position = 10
			}
			d.player2Score += d.player2Position
			d.turn = 0
		}
		updatedStates = append(updatedStates, d)
	}
	return updatedStates
}

func SplitComplete(s []State) ([]State, []State) {
	complete := []State{}
	incomplete := []State{}
	for _, state := range s {
		if state.player1Score >= 1000 || state.player2Score >= 1000 {
			complete = append(complete, state)
		} else {
			incomplete = append(incomplete, state)
		}
	}
	return complete, incomplete
}

func (s *State) GetLosingScore() (int, error) {
	if s.player1Score >= 1000 {
		return s.player2Score, nil
	}
	if s.player2Score >= 1000 {
		return s.player1Score, nil
	}
	return 0, fmt.Errorf("no one has won yet?  shouldn't be here")
}

func (s *State) HasSomeoneWon() bool {
	if s.player1Score >= 1000 || s.player2Score >= 1000 {
		return true
	}
	return false
}

func (s *State) TakeTurn() {
	rollValue := 0
	for i := 1; i <= 3; i++ { // roll dice 3 times
		s.dice++
		s.numRolls++
		rollValue += s.dice
	}
	if s.turn == 0 { // player1's turn
		s.player1Position += rollValue
		s.player1Position = s.player1Position % 10
		if s.player1Position == 0 {
			s.player1Position = 10
		}
		s.player1Score += s.player1Position
		s.turn = 1
	} else { // player2's turn
		s.player2Position += rollValue
		s.player2Position = s.player2Position % 10
		if s.player2Position == 0 {
			s.player2Position = 10
		}
		s.player2Score += s.player2Position
		s.turn = 0
	}
}
