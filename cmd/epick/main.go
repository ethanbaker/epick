package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"gitlab.com/ethanbakerdev/epick"
)

func main() {
	var testing bool
	testingUsage := "Run epick in testing mode. No GUI will be shown, only functions tested."
	flag.BoolVar(&testing, "testing", false, testingUsage)
	flag.BoolVar(&testing, "t", false, "(shorthand)")

	flag.Parse()
	args := flag.Args()

	e, err := epick.Start(testing)
	if err != nil {
		log.Fatal(err)
	}

	if testing {
		return
	}

	args = append(args, "")
	option := strings.ToLower(args[0])

	switch {
	case option == "emoji" || option == "e":
		fmt.Println(e.Emoji)

	case option == "description" || option == "desc":
		fmt.Println(e.Description)

	case option == "json":
		s, err := json.MarshalIndent(e, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(s))

	default:
		fmt.Println(e.Emoji)
	}

}
