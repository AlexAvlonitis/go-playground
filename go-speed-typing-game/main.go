// Speed Typing command line game in golang.
// https://github.com/alexavlonitis

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func readFile() string {
	content, err := ioutil.ReadFile("sentences.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileLastRow := 485
	startRange, endRange := randomRange(1, fileLastRow)
	fmt.Println(startRange, endRange)
	s := strings.Join(strings.Split(string(content), "\n")[startRange:endRange], "")

	return s
}

func randomRange(min int, max int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	startRange := rand.Intn(max-min) + min
	endRange := startRange + 2
	return startRange, endRange
}

func createQueue(file string) []string {
	s := strings.Split(file, "")

	return s
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printIntroText() {
	clearScreen()
	fmt.Println("Press ESC to quit...")
	fmt.Println("== Type the text below ==")
	fmt.Println()
}

func keyLogging(file string, logger *Logger) {
	queue := createQueue(file)
	currentChar := queue[0]

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(file)

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if currentChar == " " || currentChar == "\n" {
			currentChar = "\x00"
		}

		if string(char) == string(currentChar) {
			logger.incrementChars()
			queue = queue[1:]
			file = strings.Join(queue, "")

			if len(queue) < 1 {
				clearScreen()
				break
			}
			currentChar = queue[0]
		} else {
			logger.incrementErrors()
		}

		clearScreen()

		if key == keyboard.KeyEsc {
			break
		}
	}
}

func main() {
	file := readFile()
	logger := newLogger(0, 0)
	logger.setTotalChars(len(file))

	printIntroText()
	logger.startTimer()
	keyLogging(file, logger)
	logger.endTimer()

	logger.printStats()
}
