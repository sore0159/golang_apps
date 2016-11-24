package main

import (
	"context"
	"errors"
	"net/http"
	"os/exec"
)

type Player struct {
	Playing Playing
	Queue   Queue
}

func MakePlayerData() *Player {
	p := &Player{
	//
	}
	return p
}

type Playing struct {
	C      context.Context
	D      *SongData
	Cancel context.CancelFunc
}
type SongData struct {
	Path  string
	Tags  []string
	Count int
}

func PlaySong(path string) (*Playing, error) {
	d := &SongData{Path: path}
	path = "http://robin/music/" + path
	r, err := http.Head(path)
	if err != nil {
		return nil, err
	}
	switch r.StatusCode {
	case 404:
		return nil, errors.New("HTTP 404 File Not Found")
	case 200:
	default:
		return nil, errors.New("HTTP ERROR: " + r.Status)
	}

	c, cf := context.WithCancel(context.Background())
	err = exec.CommandContext(c, "mplayer", path).Start()
	if err != nil {
		return nil, err
	}
	p := &Playing{
		C:      c,
		D:      d,
		Cancel: cf,
	}
	return p, nil
}
