// package main qui créer une fonction variadique moyenne qui calcule la moyenne de nombres flottants.
package main

func main() {
	var nombres = []float64{10.5, 20.0, 30.25, 40.75}

	if moy, valide := moyenne(nombres...); !valide {
		println("Erreur: aucun nombre fourni")
	} else {
		println("Moyenne calculée:", moy)
	}
}

// Fonction variadique qui calcule la moyenne de nombres flottants.
func moyenne(nombres ...float64) (float64, bool) {
	if len(nombres) == 0 {
		return 0, false
	}

	var somme float64
	for _, nombre := range nombres {
		somme += nombre
	}

	return somme / float64(len(nombres)), true
}

// Exercice 1 : OK (19.0)
