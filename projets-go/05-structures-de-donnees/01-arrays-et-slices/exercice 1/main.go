// package main qui créer un programme Déclare un slice de strings avec 5 prénoms:
// Affiche chaque prénom avec son index
// Ajoute deux nouveaux prénoms
// Affiche la longueur finale
package main

import "fmt"

func main() {
	prenoms := []string{"Alice", "Bob", "Charlie", "David", "Eve"}

	for index, prenom := range prenoms {
		fmt.Printf("Prénom à l'index %d: %s\n", index, prenom)
	}

	prenoms = append(prenoms, "Gauthier", "Alex")
	fmt.Printf("Longueur finale: %d\n", len(prenoms))
}

// Exercice 1 : OK (20)
