package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("-----Welcome to my cli website----")
		fmt.Println("1 Home")
		fmt.Println("2 About")
		fmt.Println("3 projects")
		fmt.Println("4 contact")
		fmt.Println("-----------------------------------")
		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("Home page")
			fmt.Println("This cli website built on golang.")
		case "2":
			fmt.Println("About me")
			fmt.Println("My name is adam ghanmi, i have 14 years oldm i like coding with html,css,javasript and more.")
		case "3":
			fmt.Println("Projects")
			fmt.Println("This some project for me on github link https://github.com/Ghanmidev2012?tab=repositories")
		case "4":
			fmt.Println("Contact ")
			fmt.Println("you can contact with me on https://github.com/Ghanmidev2012 or in dev.to https://dev.to/ghanmidev2012")
		}
	}
}