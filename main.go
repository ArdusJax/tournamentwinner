// Package main
package main

import "fmt"

// HomeTeamWon tracks who the home team is
const HomeTeamWon = 1

// TournamentWinner is a function that determines the winner for a tournament
func TournamentWinner(competitions [][]string, results []int) string {
	m := make(map[string]int)
	winner := ""
	for i := 0; i < len(competitions); i++ {
		if results[i] == HomeTeamWon {
			winner = competitions[i][0]
		} else {
			winner = competitions[i][1]
		}
		m[winner] += 3
	}
	points := 0
	for k, v := range m {
		if v > points {
			points = v
			winner = k
		}
	}

	fmt.Println(m)
	return winner
}

// TournamentWinnerWithoutConst is a function to determine the overall winner of competitions without the use of the const
func TournamentWinnerWithoutConst(competitions [][]string, results []int) string {
	scores := map[string]int{}
	winner := ""
	for i := range competitions {
		roundWinner := competitions[i][1-results[i]]
		scores[roundWinner] += 3

		if scores[roundWinner] > scores[winner] {
			winner = roundWinner
		}
	}
	return winner
}

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 1, 1}

	answer := TournamentWinner(competitions, results)
	answer2 := TournamentWinnerWithoutConst(competitions, results)

	fmt.Println(answer)
	fmt.Println(answer2)
}
