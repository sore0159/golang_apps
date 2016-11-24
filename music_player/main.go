package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	p := MakePlayerData()
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt)
	go func() {
		<-quitChan
		fmt.Println("\nIntercepted Control-C: shutting down.")
		p.Quit()
	}()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if err := p.parse_input(scanner.Text()); err != nil {
				fmt.Println("Execution error: ", err)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("SCAN ERROR:", err)
		}
	}
}

func (p *Player) Quit() {
	fmt.Println("Goodbye!")
	p.StopPlaying()
	os.Exit(1)
}
