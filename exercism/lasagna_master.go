package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    // Initialize total preparation time
    var totalPrepTime int
    if time == 0{
        totalPrepTime = len(layers) *2
    }else{
        totalPrepTime = len(layers) * time
    }
    return totalPrepTime
}

// Quantities calculates the amount of ingredients needed based on the layers.
func Quantities(layers []string) (int, float64) {
    var sauceCount float64 = 0
    var noodlesCount int = 0

    // Iterate through all layers
    for _, layer := range layers {
        if layer == "sauce" {
            sauceCount++
        }
        if layer == "noodles" {
            noodlesCount++
        }
    }

    // Calculate quantities
    sauceAmount := sauceCount * 0.2 // Each sauce layer requires 0.2 liters
    noodleAmount := noodlesCount * 50 // Each noodle layer requires 50 grams

    return noodleAmount, sauceAmount
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string){
    if len(myList) > 0 && myList[len(myList)-1] == "?" {
        // Replace "?" with the last ingredient from friendsList
        myList[len(myList)-1] = friendsList[len(friendsList)-1]
    }
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    // Calculate the scaling factor
    scalingFactor := float64(portions) / 2.0

    // Create a new slice to hold the scaled quantities
    scaledQuantities := make([]float64, len(quantities))

    // Scale each quantity and store it in the new slice
    for i, amount := range quantities {
        scaledQuantities[i] = amount * scalingFactor
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