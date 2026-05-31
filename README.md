# Quadchecker

A program that identifies which quad function produced a given output and displays its name and dimensions.

## Usage

```bash
# Build the program
go build -o quadchecker main.go

# Test with quadA 3 3
./quadA 3 3 | go run .
# Output: [quadA] [3] [3]

# Test with quadC 1 1 (matches multiple functions)
./quadC 1 1 | go run .
# Output: [quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]

# Test with invalid input
echo 0 0 | go run .
# Output: Not a quad function
```

## Quad Functions

- **quadA**: `o` corners, `|` vertical, `-` horizontal
- **quadB**: `/` and `\` corners, `*` vertical and horizontal
- **quadC**: `A` top, `C` bottom, `B` vertical and horizontal
- **quadD**: `A` left, `C` right, `B` vertical and horizontal
- **quadE**: `A` top-left/bottom-right, `C` top-right/bottom-left, `B` vertical and horizontal
