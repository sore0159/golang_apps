package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Dots struct {
	Size   int
	Lines  [][3]int
	Boxes  [][3]int
	P1Turn bool
}

func NewDots(x int) *Dots {
	return &Dots{
		Size:   x,
		P1Turn: true,
		Lines:  [][3]int{},
	}
}

func (d *Dots) Save() error {
	if !d.Valid() {
		return errors.New("Invalid Dotsfile for saving")
	}
	file, err := os.Create("dots.json")
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(d)
}

func LoadDots() (*Dots, error) {
	file, err := os.Open("dots.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := &Dots{}
	if err := json.NewDecoder(file).Decode(&d); err != nil {
		return nil, err
	}
	if d.Valid() {
		return d, nil
	}
	return nil, errors.New("Invalid dotsfile loaded from dots.json")
}
