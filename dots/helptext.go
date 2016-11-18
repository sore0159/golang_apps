package main

import "fmt"

func DisplayUsage() {
	fmt.Println(`USAGE:
 dots -n[ew] [SIZE 3+]  Creates a new dots.json file of size SIZE
 dots -m[ove] [X] [Y]   Plays the next line at X,Y if able
                        Indicates if boxes were captured and whose turn is next
                        Warns on invalid moves
 dots -i[nfo]           Prints current game information:
                           game size, whose turn it is, current score
 dots -d[raw]           Creates dots.png file of current game state
 dots -s[erve]          Starts a http server at :8080 to serve dots.png
 dots -v[ersion]        Prints version number
 dots -h[elp]           Prints out a description of Dots & Boxes
`)
}

func DisplayHelp() {
	fmt.Println(`HELP:
    dots -m [SPACE COUNT] [DOT COUNT]
For horizontal lines, the Space Count represents the column the line is in: for vertical lines, it represents the row.  Count spaces between dots starting at "1" for the space between the top-left dot and the dot to it's right.  Count clockwise until you reach the column/row you want.

The Dot Count represents how far into that row/column to go to place your move, starting at "1".  For horizontal lines this is a count starting from the top and for vertical lines this is a count starting from the left.

    [SPACE COUNT]         [DOT COUNT]
       1     2           1     2     3
    x --- x --- x        x --- x --- x 1
  3 |     |     |        |     |     |
    x --- x --- x        x --- x --- x 2
  4 |     |     |        |     |     |
    x --- x --- x        x --- x --- x 3

EXAMPLE: dots -m 2 3  (on a size 3 board)
   x     x     x
                
   x     x     x
                
   x     x --- x

`)
}
