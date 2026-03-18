// package main quie trouve le plus grand nombre dans un slice d'entiers
package main

import "fmt"

func main() {
	// Définir un slice d'entiers
	var nombres = []int{3, 7, 2, 9, 5}

	// Trouver le plus grand nombre dans le slice et afficher le résultat
	plusGrand := trouverPlusGrand(nombres)
	fmt.Printf("Le plus grand nombre est : %d\n", plusGrand)

}

// trouverPlusGrand prend un slice d'entiers en entrée et retourne le plus grand nombre trouvé dans ce slice.
func trouverPlusGrand(nombres []int) int {
	if len(nombres) == 0 {
		return 0 // Retourne 0 si le slice est vide
	}

	// Initialiser le plus grand nombre avec le premier élément du slice
	plusGrand := nombres[0]

	// Parcourir le slice pour trouver le plus grand nombre
	for _, nombre := range nombres {
		if nombre > plusGrand {
			plusGrand = nombre
		}
	}
	return plusGrand
}

// Exercice 3 : Smi-ok (18)
