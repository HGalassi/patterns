package main

import "fmt"

func main() {
	weeklyRoutine := NewDailyRoutine().SetEat("3 meals").
		SetFamilyTime(2).
		SetWork(8).
		SetSleep("7-8 hours").
		SetProgramming("2 hours").
		SetHobby(true).
		SetExercise(true).
		SetLanguageStudy(true).
		Build()

	fmt.Printf("%+v\n", weeklyRoutine)

	dailyRoutine := NewDailyRoutine().SetEat("2 meals").
		SetFamilyTime(1).
		SetSleep("6 hours").
		SetProgramming("1 hour").
		SetHobby(false).
		Build()

	fmt.Printf("%+v\n", dailyRoutine)
}
