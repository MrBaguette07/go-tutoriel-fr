package main

import (
	"fmt"
	"math"
)

// Calcul de l'aire d'un cercle avec un rayon donné
func main() {
	// Rayon du cercle
	r := 5.0

	// Calcul de l'aire du cercle avec la formule A = π * r^2
	result := math.Pi * r * r

	// Affichage du résultat avec 2 décimales
	fmt.Printf("%.2f\n", result)
}

// Exercice 3 : Semi-ok
