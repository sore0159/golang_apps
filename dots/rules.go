package main

func (d *Dots) Score() [2]int {
	score := [2]int{}
	for _, b := range d.Boxes {
		if b[2] == 1 {
			score[0] += 1
		} else if b[2] == 2 {
			score[1] += 1
		}
	}
	return score
}

func (d *Dots) AddLine(ln [2]int) (boxes byte, valid bool) {
	if !d.ValidLine(ln) {
		return 0, false
	}
	if !d.FreeLine(ln) {
		return 0, false
	}
	n := [3]int{ln[0], ln[1], 0}
	x1, y1, x2, y2 := d.LineBoxes(n)
	b1, b2 := [3]int{x1, y1, 0}, [3]int{x2, y2, 0}
	if d.P1Turn {
		n[2] = 1
		b1[2] = 1
		b2[2] = 1
	} else {
		n[2] = 2
		b1[2] = 2
		b2[2] = 2
	}
	if x1 == 0 {
		b1 = [3]int{-1, -1, -1}
	}
	if x2 == 0 {
		b2 = [3]int{-1, -1, -1}
	}
	// Box Check //
	var c1, c2 int
	for _, test := range d.Lines {
		x1, y1, x2, y2 := d.LineBoxes(test)
		if (x1 == b1[0] && y1 == b1[1]) || (x2 == b1[0] && y2 == b1[1]) {
			c1 += 1
		}
		if (x1 == b2[0] && y1 == b2[1]) || (x2 == b2[0] && y2 == b2[1]) {
			c2 += 1
		}
	}
	if c1 == 3 {
		d.Boxes = append(d.Boxes, b1)
		boxes += 1
	}
	if c2 == 3 {
		d.Boxes = append(d.Boxes, b2)
		boxes += 1
	}
	//           //
	d.Lines = append(d.Lines, n)
	return boxes, true
}
