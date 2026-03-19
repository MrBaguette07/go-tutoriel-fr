package main

import "fmt"

func main() {
	animaux := [...]string{"chat", "chien", "lapin", "hamster", "oiseau"}

	fmt.Println("Tableau d'animaux:", animaux)

	animaux[1] = "tortue"
	fmt.Println("Après modification:", animaux)

	fmt.Println(len(animaux))

	for i := 0; i < len(animaux); i++ {
		fmt.Printf("Animal à l'index %d: %s\n", i, animaux[i])
	}

	for index, animal := range animaux {
		fmt.Printf("Animal à l'index %d: %s\n", index, animal)
	}

	for _, animal := range animaux {
		fmt.Printf("Animal: %s\n", animal)
	}

	petitArray := [...]int{1, 2, 3}
	var grandArray []int = []int{1, 2, 3, 4, 5}

	fmt.Printf("petitArray : %v\n", petitArray)
	fmt.Printf("grandArray : %v\n", grandArray)

	slice := make([]string, 3, 5)

	fmt.Println("Longueur", len(slice))
	fmt.Println("Capacité", cap(slice))
	fmt.Println("Slice", slice)

	grandArray = append(grandArray, 6, 7, 8)
	fmt.Printf("grandArray après append: %v\n", grandArray)

	array := [5]int{1, 2, 3, 4, 5}
	slice1 := array[1:4]
	fmt.Printf("Slice1: %v\n", slice1)

	array[2] = 99
	fmt.Printf("Array après modification: %v\n", array)
	fmt.Printf("Slice1 après modification de l'array: %v\n", slice1)

	// Pour copier des slices :

	original := []string{"zonix.fr", "lakestart.fr", "seapalm.fr"}

	copie := make([]string, len(original))
	copy(copie, original)

	original[0] = "modifié.fr"

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copie: %v\n", copie)
}
