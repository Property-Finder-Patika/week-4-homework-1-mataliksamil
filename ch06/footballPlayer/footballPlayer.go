package main

import "fmt"

type FootballPlayer struct {
	name      string
	team      string
	goalCount int
}

func main() {
	salah := FootballPlayer{"Salah", "Liverpool", 18}
	newGoalCount := score(salah)
	salah.goalCount = newGoalCount
	fmt.Printf("%#v\n", salah)
	salah.score1()
	fmt.Printf("%#v\n", salah)
	salah.score2()
	fmt.Printf("%#v\n", salah)
}

func score(fp FootballPlayer) int {
	fp.goalCount++
	return fp.goalCount
}

func (fp FootballPlayer) score1() {
	fp.goalCount++
}

func (fp *FootballPlayer) score2() {
	fp.goalCount++
}
