package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	var lessThanMonth []*github.Issue
	var lessThanYear []*github.Issue
	var moreThanYear []*github.Issue
	currentTime := time.Now()
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		diff := currentTime.Sub(item.CreatedAt)
		// fmt.Println(diff)
		dif := int64(diff.Hours() / 24)
		// fmt.Println("In days", dif)
		if dif < 30 {
			lessThanMonth = append(lessThanMonth, item)
		} else if dif > 30 && dif < 365 {
			lessThanYear = append(lessThanYear, item)
		} else if dif > 365 {
			moreThanYear = append(moreThanYear, item)
		}

	}
	fmt.Println("This is issues with less than 30 days")
	for _, item := range lessThanMonth {
		fmt.Printf("#%5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("This is issues with less than a year")
	for _, item := range lessThanYear {
		fmt.Printf("#%5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("This is issues with more than a year")
	for _, item := range moreThanYear {
		fmt.Printf("#%5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
