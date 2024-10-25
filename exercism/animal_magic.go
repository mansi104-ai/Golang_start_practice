package chance
import "math/rand"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	n := 1+ rand.Intn(20-1+1)
    return n
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	// panic("Please implement the GenerateWandEnergy function")
    f:=  rand.Float64()
    return f*12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	// panic("Please implement the ShuffleAnimals function")
    x:= []string{"ant","beaver","cat","dog","elephant","fox","giraffe",
"hedgehog"}
    rand.Shuffle(len(x), func(i,j int){
        x[i],x[j] = x[j],x[i]
    })
    return x
                 
}