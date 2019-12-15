package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// EXIT flag
	EXIT = "exit"
	// DOMAIN enter a domain name as a prefix to the urls
	DOMAIN = "http://test.com"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand := newRandomiser()
	stor := newStorage()

	showMenu()

	i := ""
	for i != EXIT {
		fmt.Print("Enter option: ")
		option, _ := reader.ReadString('\n')

		i = strings.TrimRight(option, "\n")

		if i == "1" {
			fmt.Print("Enter url to shorten: ")
			url, _ := reader.ReadString('\n')
			i = strings.TrimRight(url, "\n")

			shortURL := rand.create(6)
			fmt.Println(stor.add(shortURL, url))
		}

		if i == "2" {
			fmt.Print("Enter short URL: ")
			url, _ := reader.ReadString('\n')
			tURL := strings.TrimRight(url, "\n")

			fmt.Println(stor.read(tURL))
		}
	}
}

func showMenu() {
	fmt.Println("Type exit, to exit the program")
	fmt.Println("Type 1, to shorten a URL")
	fmt.Println("Type 2, to retrieve a URL")
}
