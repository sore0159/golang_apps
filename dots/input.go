package main

//      line[0]                line[1]
//
//       1     2            1     2     3
//    x --- x --- x       1 x --- x --- x
//  3 |     |     |         |     |     |
//    x --- x --- x       2 x --- x --- x
//  4 |     |     |         |     |     |
//    x --- x --- x       3 x --- x --- x

func (d *Dots) ValidLine(ln [2]int) (valid bool) {
	return !(ln[0] < 1 || ln[0] > (d.Size-1)*2 || ln[1] < 1 || ln[1] > d.Size)
}

func (d *Dots) FreeLine(ln [2]int) (free bool) {
	for _, test := range d.Lines {
		if ln[0] == test[0] && ln[1] == test[1] {
			return false
		}
	}
	return true
}

func (d *Dots) Valid() bool {
	if d.Size < 3 {
		return false
	}
	for i, ln := range d.Lines {
		test := [2]int{ln[0], ln[1]}
		if !d.ValidLine(test) {
			return false
		}
		for _, ln2 := range d.Lines[i+1:] {
			if ln2[0] == ln[0] && ln[1] == ln2[1] {
				return false
			}
		}
	}
	return true
}

func (d *Dots) LineDots(ln [3]int) (x1, y1, x2, y2 int) {
	if ln[0] < d.Size {
		x1, x2 = ln[0], ln[0]+1
		y1, y2 = ln[1], ln[1]
	} else {
		x1, x2 = ln[1], ln[1]
		y1, y2 = ln[0]+1-d.Size, ln[0]+2-d.Size
	}
	return
}

func (d *Dots) LineBoxes(ln [3]int) (x1, y1, x2, y2 int) {
	if ln[0] < d.Size {
		x1, x2 = ln[0], ln[0]
		y1, y2 = ln[1]-1, ln[1]
	} else {
		y1, y2 = ln[0]+1-d.Size, ln[0]+1-d.Size
		x1, x2 = ln[1]-1, ln[1]
	}
	if x1 < 1 || y1 < 1 || x1 >= d.Size || y1 >= d.Size {
		x1, y1 = 0, 0
	}
	if x2 < 1 || y2 < 1 || x2 >= d.Size || y2 >= d.Size {
		x2, y2 = 0, 0
	}
	return
}
