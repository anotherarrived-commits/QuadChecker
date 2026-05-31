package main

import (
	"fmt"
	"os"
	"strconv"
)

func renderQuad(x, y int, tlc, trc, blc, brc, vl, hl rune) {
	if x <= 0 || y <= 0 {
		return
	}

	for h := 1; h <= y; h++ {
		isTopRow := h == 1
		isBottomRow := h == y

		for w := 1; w <= x; w++ {
			isLeftCol := w == 1
			isRightCol := w == x

			switch {
			case isTopRow && isLeftCol:
				fmt.Print(string(tlc))
			case isTopRow && isRightCol:
				fmt.Print(string(trc))
			case isBottomRow && isLeftCol:
				fmt.Print(string(blc))
			case isBottomRow && isRightCol:
				fmt.Print(string(brc))
			case isTopRow || isBottomRow:
				fmt.Print(string(hl))
			case isLeftCol || isRightCol:
				fmt.Print(string(vl))
			default:
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func quadB(x, y int) {
	renderQuad(x, y, '/', '\\', '\\', '/', '*', '*')
}

func main() {
	if len(os.Args) < 3 {
		return
	}

	x, errX := strconv.Atoi(os.Args[1])
	y, errY := strconv.Atoi(os.Args[2])

	if errX != nil || errY != nil {
		return
	}

	quadB(x, y)
}
