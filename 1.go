package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("fruit.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dict := make(map[string]int)

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		_, ok := dict[scanner.Text()]
		if ok {
			dict[scanner.Text()] = dict[scanner.Text()] + 1
		} else {
			dict[scanner.Text()] = 1
		}
	}

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range dict {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	f, err := os.Create("fruit-count.txt")

	defer f.Close()

	//f.WriteString(fmt.Sprintf("%s, %d\n", kv.Key, kv.Value))
	fmt.Fprintf(f, "果物名,数\n")
	for _, kv := range ss {
		fmt.Fprintf(f, "%s,%d\n", kv.Key, kv.Value)
	}
}
