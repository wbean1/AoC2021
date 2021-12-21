package day21

import (
	"fmt"
	"log"

	"github.com/wbean1/AoC/utils"
)

type State struct {
	player1Position, player2Position int
	player1Score, player2Score       int
	dice, numRolls                   int
	turn                             int // 0 = player1's turn, 1 = player2's turn
}

type DState struct {
	player1Position, player2Position uint8
	player1Score, player2Score       uint8
	rollValue                        uint8
	universeCount                    uint64
	turn                             bool // 0 = player1's turn, 1 = player2's turn
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
	diracStates := []DState{DState{player1Position: 1, player2Position: 6, universeCount: 1}}
	diracStatesComplete, diracStatesIncomplete := SplitComplete(diracStates)
	var p1wins, p2wins uint64
	for len(diracStatesIncomplete) > 0 {
		fmt.Printf("completeStates: %d, incompleteStates: %d\n", p1wins+p2wins, len(diracStatesIncomplete))
		for _, d := range diracStatesIncomplete {
			diracStatesComplete = append(diracStatesComplete, TakeDiracTurn(d)...)
		}
		diracStatesComplete, diracStatesIncomplete = SplitComplete(diracStatesComplete)
		player1Wins, player2Wins := CountWins(diracStatesComplete)
		p1wins += player1Wins
		p2wins += player2Wins
		diracStatesComplete = []DState{}
	}
	fmt.Printf("part2: player1 wins: %d, player2 wins: %d, max: %d\n", p1wins, p2wins, utils.Max64(p1wins, p2wins))
}

func CountWins(s []DState) (uint64, uint64) {
	var p1wins, p2wins uint64
	for _, d := range s {
		if d.player1Score >= 21 {
			p1wins += d.universeCount
		} else if d.player2Score >= 21 {
			p2wins += d.universeCount
		} else {
			log.Fatal("shouldn't get here ever")
		}
	}
	return p1wins, p2wins
}

func TakeDiracTurn(s DState) []DState {
	rollMap := map[uint8]uint64{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	} // map of rollValue to number of universes that get split into this...
	states := []DState{}
	for value, count := range rollMap {
		copy := s
		copy.rollValue = value
		copy.universeCount *= count
		states = append(states, copy)
	}
	updatedStates := []DState{}
	for _, d := range states {
		if d.turn { // player1's turn
			d.player1Position += d.rollValue
			d.player1Position = d.player1Position % 10
			if d.player1Position == 0 {
				d.player1Position = 10
			}
			d.player1Score += d.player1Position
			d.turn = false
		} else { // player2's turn
			d.player2Position += d.rollValue
			d.player2Position = d.player2Position % 10
			if d.player2Position == 0 {
				d.player2Position = 10
			}
			d.player2Score += d.player2Position
			d.turn = true
		}
		updatedStates = append(updatedStates, d)
	}
	return updatedStates
}

func SplitComplete(s []DState) ([]DState, []DState) {
	complete := []DState{}
	incomplete := []DState{}
	for _, state := range s {
		if state.player1Score >= 21 || state.player2Score >= 21 {
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
