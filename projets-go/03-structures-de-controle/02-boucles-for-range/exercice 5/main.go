// package main qui génère les n premières lignes du triangle de Pascal
package main

import "fmt"

func main() {
	// Définir le nombre de lignes du triangle de Pascal à générer
	n := 10

	// Générer le triangle de Pascal et afficher les résultats
	triangle := genererTrianglePascal(n)
	for _, ligne := range triangle {
		println(ligne)
	}

}

func genererTrianglePascal(n int) []string {
	// Initialiser une slice pour stocker les lignes du triangle de Pascal
	triangle := make([]string, n)
	sliceTriangle := make([][]int, n)

	for i := 0; i < n; i++ {
		// Initialiser la ligne actuelle du triangle de Pascal
		sliceTriangle[i] = make([]int, i+1)
		sliceTriangle[i][0] = 1
		sliceTriangle[i][i] = 1

		// Calculer les éléments de la ligne actuelle en utilisant les éléments de la ligne précédente
		for j := 1; j < i; j++ {
			sliceTriangle[i][j] = sliceTriangle[i-1][j-1] + sliceTriangle[i-1][j]
		}

		// Convertir la ligne actuelle du triangle de Pascal en une string et la stocker dans la slice triangle
		ligne := ""
		for _, element := range sliceTriangle[i] {
			ligne += fmt.Sprintf("%d ", element)
		}

		triangle[i] = ligne
	}
	return triangle
}

// Exercice 5 : Smi-ok (19)
