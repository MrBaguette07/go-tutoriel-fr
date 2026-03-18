// package main qui calcule la factorielle d'un nombre (n! = 1 × 2 × 3 × ... × n).
package main

import "fmt"

func main() {
	// Définir un nombre pour lequel calculer la factorielle
	nombre := 5
	// Calculer la factorielle et afficher le résultat
	resultat := factorielle(nombre)
	if resultat != -1 {
		fmt.Printf("La factorielle de %d est %d.\n", nombre, resultat)
	} else {
		fmt.Printf("La factorielle n'est pas définie pour les nombres négatifs (%d).\n", nombre)
	}
}

func factorielle(n int) int {
	if n < 0 {
		return -1 // La factorielle n'est pas définie pour les nombres négatifs
	}

	// La factorielle de 0 et 1 est 1
	if n == 0 || n == 1 {
		return 1 // La factorielle de 0 et 1 est 1
	}

	// Initialiser le résultat à 1 (car 0! = 1)
	resultat := 1

	// Utiliser une boucle for pour calculer la factorielle
	for i := 2; i <= n; i++ {
		resultat *= i
	}
	return resultat
}

// Exercice 2 : Smi-ok (18)
