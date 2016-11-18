package main

import (
	"fmt"
	"os"
	"strconv"
)

const VERSION = 1

func main() {
	if len(os.Args) < 2 {
		DisplayUsage()
		return
	}
	switch os.Args[1] {
	case "-h", "-help":
		DisplayHelp()
		return
	case "-n", "-new":
		if len(os.Args) != 3 {
			DisplayUsage()
			return
		}
		x, err := strconv.Atoi(os.Args[2])
		if err != nil || x < 3 {
			DisplayUsage()
			return
		}
		d := NewDots(x)
		if err := d.Save(); err != nil {
			fmt.Println("Error Creating Dotsfile: ", err)
		}
		fmt.Printf("Game size %d created, Player One's turn.\n", x)
		if err = d.Draw(); err != nil {
			fmt.Println("Error drawing dots: ", err)
		}
		fmt.Println("dots.png created")
		return
	case "-v", "-version":
		fmt.Printf("Dots & Boxes, version %d\n", VERSION)
		return
	case "-d", "-draw":
		d, err := LoadDots()
		if err != nil {
			fmt.Println("Error loading dotsfile: ", err)
			return
		}

		if err = d.Draw(); err != nil {
			fmt.Println("Error drawing dots: ", err)
		}
		fmt.Println("dots.png created")
		return
	case "-i", "-info":
		d, err := LoadDots()
		if err != nil {
			fmt.Println("Dots & Boxes: No game found!")
			return
		} else {
			score := d.Score()
			var turn string
			if d.P1Turn {
				turn = "Player One"
			} else {
				turn = "Player Two"
			}
			fmt.Printf(`Dots & Boxes:
  Game Size:%d
  Current Score: Player 1: %d, Player 2: %d
  %s's Turn
\n`, d.Size, score[0], score[1], turn)
		}
	case "-m", "-move":
		if len(os.Args) != 4 {
			DisplayUsage()
			return
		}
		x, err1 := strconv.Atoi(os.Args[2])
		y, err2 := strconv.Atoi(os.Args[3])
		if err1 != nil || err2 != nil {
			DisplayUsage()
			return
		}
		d, err := LoadDots()
		if err != nil {
			fmt.Println("Error loading dotsfile: ", err)
			return
		}
		ln := [2]int{x, y}
		if !d.ValidLine(ln) {
			DisplayHelp()
			return
		}
		if !d.FreeLine(ln) {
			fmt.Println("That line is already taken!\n")
			return
		}
		boxes, valid := d.AddLine(ln)
		if !valid {
			DisplayHelp()
			return
		}
		player, next := "Player One", "Player Two"
		if !d.P1Turn {
			player, next = next, player
		}
		switch boxes {
		case 1:
			fmt.Printf("%s took one box and gets another move.\n", player)
		case 2:
			fmt.Printf("%s took two boxes and gets another move.\n", player)
		default:
			d.P1Turn = !d.P1Turn
			fmt.Printf("%s took no boxes, it is now %s's move.\n", player, next)
		}
		if err := d.Save(); err != nil {
			fmt.Println("Error Saving Dotsfile: ", err)
		}
		if err = d.Draw(); err != nil {
			fmt.Println("Error drawing dots: ", err)
		}
		fmt.Println("dots.png created")
	case "-s", "-serve":
		ServeHTTP()
		return
	default:
		DisplayUsage()
		return
	}
}
