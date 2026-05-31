package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// renderQuad generates the quad pattern
func renderQuad(x, y int, tlc, trc, blc, brc, vl, hl rune) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	var result strings.Builder

	for h := 1; h <= y; h++ {
		isTopRow := h == 1
		isBottomRow := h == y

		for w := 1; w <= x; w++ {
			isLeftCol := w == 1
			isRightCol := w == x

			switch {
			case isTopRow && isLeftCol:
				result.WriteRune(tlc)
			case isTopRow && isRightCol:
				result.WriteRune(trc)
			case isBottomRow && isLeftCol:
				result.WriteRune(blc)
			case isBottomRow && isRightCol:
				result.WriteRune(brc)
			case isTopRow || isBottomRow:
				result.WriteRune(hl)
			case isLeftCol || isRightCol:
				result.WriteRune(vl)
			default:
				result.WriteRune(' ')
			}
		}
		result.WriteRune('\n')
	}

	return result.String()
}

// Quad functions
func quadA(x, y int) string {
	return renderQuad(x, y, 'o', 'o', 'o', 'o', '|', '-')
}

func quadB(x, y int) string {
	return renderQuad(x, y, '/', '\\', '\\', '/', '*', '*')
}

func quadC(x, y int) string {
	return renderQuad(x, y, 'A', 'A', 'C', 'C', 'B', 'B')
}

func quadD(x, y int) string {
	return renderQuad(x, y, 'A', 'C', 'A', 'C', 'B', 'B')
}

func quadE(x, y int) string {
	return renderQuad(x, y, 'A', 'C', 'C', 'A', 'B', 'B')
}

func main() {
	// Read all input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	var inputLines []string
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if len(inputLines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Reconstruct input with newlines
	inputOutput := strings.Join(inputLines, "\n")
	if inputOutput != "" {
		inputOutput += "\n"
	}

	// Map of quad functions with their names
	quads := []struct {
		name string
		fn   func(int, int) string
	}{
		{"quadA", quadA},
		{"quadB", quadB},
		{"quadC", quadC},
		{"quadD", quadD},
		{"quadE", quadE},
	}

	var matches []string

	// Try all reasonable x, y dimensions (limit to reasonable size)
	for x := 1; x <= 50; x++ {
		for y := 1; y <= 50; y++ {
			for _, quad := range quads {
				if quad.fn(x, y) == inputOutput {
					matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad.name, x, y))
				}
			}
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Sort alphabetically
	sort.Strings(matches)

	// Output with " || " separator and newline
	fmt.Println(strings.Join(matches, " || "))
}
