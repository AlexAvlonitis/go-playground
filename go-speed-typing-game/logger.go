package main

import (
	"fmt"
	"time"
)

type Logger struct {
	totalChars  int
	numOfChars  int
	numOfErrors int
	start       time.Time
	end         time.Time
}

func newLogger(numOfChars int, numOfErrors int) *Logger {
	return &Logger{
		numOfChars:  numOfChars,
		numOfErrors: numOfErrors,
	}
}

func (l *Logger) incrementChars() {
	l.numOfChars += 1
}

func (l *Logger) incrementErrors() {
	l.numOfErrors += 1
}

func (l *Logger) setTotalChars(chars int) {
	l.totalChars = chars
}

func (l *Logger) startTimer() {
	l.start = time.Now()
}

func (l *Logger) endTimer() {
	l.end = time.Now()
}

func (l *Logger) printStats() {
	finishedTime := l.end.Sub(l.start)
	fmt.Println("Total characters:", l.totalChars)

	fmt.Println(
		"Time to finish (seconds):",
		fmt.Sprintf("%.2f", finishedTime.Seconds()))

	fmt.Println("Number of errors:", l.numOfErrors)

	speed := float64(l.numOfChars) / (finishedTime.Seconds() / 60)
	fmt.Println("Characters per minute:", fmt.Sprintf("%.2f", speed))
}
