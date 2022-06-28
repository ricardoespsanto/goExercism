package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	sauce := 0.0
	noodles := 0

	for _, i := range layers {
		if i == "sauce" {
			sauce += 0.2
		} else if i == "noodles" {
			noodles++
		}
	}

	return noodles * 50, sauce
}

func AddSecretIngredient(friendsIng []string, myList []string) {
	myList[len(myList)-1] = friendsIng[len(friendsIng)-1]
}

func ScaleRecipe(quantitiesForTwo []float64, portions int) []float64 {
	var quantitiesForX []float64
	for _, i := range quantitiesForTwo {
		quantitiesForX = append(quantitiesForX, i*(float64(portions)/2))
	}

	return quantitiesForX
}
