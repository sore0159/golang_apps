package main

import (
	"log"
	"os"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}
func TestTwo(t *testing.T) {
	pts := [][2]float64{
		[2]float64{130, 130},
		[2]float64{130, 400},
		[2]float64{230, 200},
		[2]float64{350, 370},
	}
	file, err := os.Create("bez_test.png")
	if err != nil {
		log.Println("Error creating png file")
		return
	}
	DrawBez(file, pts)
	log.Println("bez_test.png created")
}
