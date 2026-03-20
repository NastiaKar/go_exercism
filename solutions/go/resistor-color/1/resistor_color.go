package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	colors := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
    return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colors := Colors()
    for i, _ := range colors {
        if colors[i] == color {
            return i
        }
    }
    return 0
}
