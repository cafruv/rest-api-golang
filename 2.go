package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("fruit-count.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dict := make(map[string]string)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		s := strings.Split(scanner.Text(), ",")
		//fmt.Println(s[0])
		dict[s[0]+"\n"] = s[1]
		//fmt.Println(s[1])
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	for {
		text, _ := reader.ReadString('\n')
		if text == "exit\n" {
			os.Exit(0)
		}
		v, ok := dict[text]
		if ok {
			fmt.Println(v)
		} else {
			fmt.Println(0)
		}
	}

}
