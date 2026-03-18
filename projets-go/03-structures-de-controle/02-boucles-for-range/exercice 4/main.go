// package main qui inverse une chaine de caractères
package main

func main() {
	// Définir une chaîne de caractères à inverser
	chaine := "Seapalm est vraiment une super entreprise !"

	// Inverser la chaîne de caractères et afficher le résultat
	chaineInversee := inverserChaine(chaine)
	println("Chaîne originale : ", chaine)
	println("Chaîne inversée : ", chaineInversee)

}

// inverserChaine prend une chaîne de caractères en entrée et retourne la chaîne inversée.
func inverserChaine(chaine string) string {
	// Convertir la chaîne de caractères en une slice de runes pour gérer les caractères Unicode
	runes := []rune(chaine)

	// Initialiser une slice de runes pour stocker les caractères inversés
	inversee := make([]rune, len(runes))

	// Utiliser une boucle for pour remplir la slice inversée
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		inversee[i] = runes[j]
	}

	// Convertir la slice de runes inversée en une chaîne de caractères et la retourner
	return string(inversee)
}

// Exercice 4 : Smi-ok (18)
