package main

import (
	"fmt"
	"github.com/A-ndrey/leetcode-client/internal/usecases"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("specify username")
	}

	stat, err := usecases.UserProblemsSolved(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	templ := "%s: %d/%d\n"

	if stat.AllTotal != 0 {
		fmt.Printf(templ, "All", stat.AllScore, stat.AllTotal)
	}
	if stat.EasyTotal != 0 {
		fmt.Printf(templ, "Easy", stat.EasyScore, stat.EasyTotal)
	}
	if stat.MediumTotal != 0 {
		fmt.Printf(templ, "Medium", stat.MediumScore, stat.MediumTotal)
	}
	if stat.HardTotal != 0 {
		fmt.Printf(templ, "Hard", stat.HardScore, stat.HardTotal)
	}
}
