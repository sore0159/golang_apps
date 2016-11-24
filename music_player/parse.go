package main

func (p *Player) parse_input(input string) (err error) {
	p.StopPlaying()
	s, err := PlaySong(input)
	if err != nil {
		return err
	}
	p.Playing = *s
	return nil
}

func (p *Player) StopPlaying() {
	if p.Playing.C != nil {
		p.Playing.Cancel()
	}
}
