package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, average int) int {
    if average == 0 {
        average = 2
    }
    return len(layers) * average
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodlesAmount := 0;
    sauceAmount := 0.0;
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodlesAmount += 50
        } else if layers[i] == "sauce" {
            sauceAmount += 0.2
        }
    }
    return noodlesAmount, sauceAmount
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, servings int) []float64{
    scaledQuantities := make([]float64, len(quantities))
    for i := 0; i < len(quantities); i++ {
        scaledQuantities[i] = quantities[i] * float64(servings) / 2.0
    }
    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
