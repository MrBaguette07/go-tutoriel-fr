// package main qui vérifie si un nombre est premier
package main

import "fmt"

func main() {
	// Définir un nombre à vérifier
	nombre := 8

	// Vérifier si le nombre est premier et afficher le résultat
	if EstPremier(nombre) {
		fmt.Printf("%d est un nombre premier.\n", nombre)
	} else {
		fmt.Printf("%d n'est pas un nombre premier.\n", nombre)
	}
}

// EstPremier vérifie si un nombre est premier en vérifiant les diviseurs jusqu'à la racine carrée du nombre.
func EstPremier(n int) bool {
	// Les nombres inférieurs ou égaux à 1 ne sont pas premiers
	if n <= 1 {
		return false
	}

	// Le nombre 2 est le seul nombre premier pair
	if n == 2 {
		return true
	}

	// Les nombres pairs autres que 2 ne sont pas premiers
	if n%2 == 0 {
		return false
	}

	// Vérifier les diviseurs impairs de 3 à la racine carrée de n
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// Exercice 1 : Smi-ok (18)
