package lasagnamaster

import (
	"fmt"
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgpreptime int) int {
	if avgpreptime == 0 {
		avgpreptime = 2
	}
	return len(layers) * avgpreptime

}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodleweight int
	var saucevolume float64
	for _, k := range layers {
		switch k {
		case "noodles":
			noodleweight += 50
		case "sauce":
			saucevolume += 0.2
		}
	}
	return noodleweight, saucevolume
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ingredient1, ingredient2 []string) {
	ingredient2[len(ingredient2)-1] = ingredient1[len(ingredient1)-1]
	fmt.Println(ingredient2)
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(qty []float64, portions int) []float64 {
	var originalqty []float64

	for i := range qty {
		originalqty = append(originalqty, qty[i]*float64(portions)/2.0)
	}
	return originalqty
}
