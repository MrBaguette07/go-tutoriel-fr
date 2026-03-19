// package main qui créer une fonction qui :
// Prend un slice de strings
// Retourne un nouveau slice sans les doublons
// Préserve l'ordre de la première occurrence
package main

func main() {
	strs := []string{"apple", "banana", "apple", "orange", "banana", "grape"}

	// Appel de la fonction pour supprimer les doublons
	result := supprimerDoublons(strs)

	// Affichage du résultat
	for _, str := range result {
		println(str)
	}
}

// Fonction qui supprime les doublons d'un slice de strings
func supprimerDoublons(strs []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, str := range strs {
		if !seen[str] {
			seen[str] = true
			result = append(result, str)
		}
	}
	return result
}

// Exercice 3 : OK (20)
