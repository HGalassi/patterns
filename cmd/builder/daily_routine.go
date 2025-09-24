package main

type dailyRoutine struct {
	familyTime    int
	work          int
	sleep         string
	eat           string
	programming   string
	hasHobby      bool
	exercise      bool
	languageStudy bool
}

type DailyRoutineBuilder struct {
	dailyRoutine dailyRoutine
}

func NewDailyRoutineBuilder() *DailyRoutineBuilder {
	return &DailyRoutineBuilder{dailyRoutine: dailyRoutine{}}
}

func (b *DailyRoutineBuilder) SetFamilyTime(hours int) *DailyRoutineBuilder {
	b.dailyRoutine.familyTime = hours
	return b
}
func (b *DailyRoutineBuilder) SetWork(hours int) *DailyRoutineBuilder {
	b.dailyRoutine.work = hours
	return b
}
func (b *DailyRoutineBuilder) SetSleep(hours string) *DailyRoutineBuilder {
	b.dailyRoutine.sleep = hours
	return b
}
func (b *DailyRoutineBuilder) SetEat(meals string) *DailyRoutineBuilder {
	b.dailyRoutine.eat = meals
	return b
}
func (b *DailyRoutineBuilder) SetProgramming(hours string) *DailyRoutineBuilder {
	b.dailyRoutine.programming = hours
	return b
}
func (b *DailyRoutineBuilder) SetHobby(hasHobby bool) *DailyRoutineBuilder {
	b.dailyRoutine.hasHobby = hasHobby
	return b
}
func (b *DailyRoutineBuilder) SetExercise(exercise bool) *DailyRoutineBuilder {
	b.dailyRoutine.exercise = exercise
	return b
}
func (b *DailyRoutineBuilder) SetLanguageStudy(language_study bool) *DailyRoutineBuilder {
	b.dailyRoutine.languageStudy = language_study
	return b
}
func (b *DailyRoutineBuilder) Build() dailyRoutine {
	return b.dailyRoutine
}
func NewDailyRoutine() *DailyRoutineBuilder {
	return NewDailyRoutineBuilder()
}
