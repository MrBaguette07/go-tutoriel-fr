// package main qui Créer une fonction qui analyse un slice d'entiers et retourne la moyenne, le minimum, le maximum et la somme.
package main

import "fmt"

func main() {
	nombres := []int{10, 20, 30, 40, 50}
	moyenne, minimum, maximum, somme := analyserSlice(nombres)
	fmt.Printf("Moyenne: %.2f, Minimum: %d, Maximum: %d, Somme: %d\n", moyenne, minimum, maximum, somme)
}

// Fonction qui retourne la moyenne de deux entiers.
func moyenne(a, b int) float64 {
	return float64(a+b) / 2
}

// Fonction qui retourne le minimum de deux entiers.
func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Fonction qui retourne le maximum de deux entiers.
func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Fonction qui calcule la somme de deux entiers.
func somme(a, b int) int {
	return a + b
}

// Fonction qui analyse un slice d'entiers et retourne la moyenne, le minimum, le maximum et la somme.
func analyserSlice(nombres []int) (float64, int, int, int) {
	// Utiliser les fonctions définies pour calculer les valeurs
	var total int
	min := nombres[0]
	max := nombres[0]
	for _, nombre := range nombres {
		total = somme(total, nombre)
		min = minimum(min, nombre)
		max = maximum(max, nombre)
	}
	moy := moyenne(total, len(nombres))
	return moy, min, max, total
}

// Exercice 3 : OK (20.0)
