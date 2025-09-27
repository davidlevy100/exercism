// Package lasagna provides functions for calculating preparation
// time, ingredient needs, and scaling lasagna recipes.
package lasagna

// PreparationTime returns the total prep time given layers and time per layer.
// Defaults to 2 minutes if timePerLayer is 0.
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return timePerLayer * len(layers)
}

// Quantities returns the grams of noodles and liters of sauce
// needed for the given layers (50g noodles, 0.2l sauce per layer).
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient replaces the last item in myList with the
// last item from friendsList.
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe returns scaled ingredient quantities for the desired
// portions, based on a 2-portion base.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	result := make([]float64, len(quantities))
	factor := float64(portions) / 2.0
	for i, q := range quantities {
		result[i] = q * factor
	}
	return result
}
