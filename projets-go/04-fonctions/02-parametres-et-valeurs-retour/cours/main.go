package main

import "fmt"

// Cette fonction modifie un élément d'un slice, ce qui affecte le slice original car les slices sont des références à des tableaux sous-jacents.
func modifierSlice(s []int) {
	s[0] = 999 // Modifie l'élément original
	fmt.Println("Dans la fonction:", s)
}

// Cette fonction tente d'ajouter un élément à un slice, mais cela peut créer un nouveau slice si la capacité est dépassée.
func ajouterElementSlice(s []int) []int {
	s = append(s, 100) // Ceci peut créer un nouveau slice
	fmt.Println("Dans la fonction après append:", s)
	return s
}

func main() {
	// Test avec modification d'élément
	nombres := []int{1, 2, 3}
	fmt.Println("Avant modification:", nombres)
	modifierSlice(nombres)
	fmt.Println("Après modification:", nombres)

	fmt.Println("---")

	// Test avec ajout d'élément
	nombres2 := []int{1, 2, 3}
	fmt.Println("Avant append:", nombres2)
	nouveauSlice := ajouterElementSlice(nombres2)
	fmt.Println("Après append - original:", nombres2)
	fmt.Println("Après append - nouveau:", nouveauSlice)

	logiqueComplexe()
	retourFonctions()
}

// Fonction qui analyse une note et retourne plusieurs valeurs : la mention, la réussite et un commentaire.
func analyserNote(note int) (mention string, reussite bool, commentaire string) {
	// Initialisation par défaut
	reussite = false
	commentaire = "Analyse de la note"

	switch {
	case note >= 90:
		mention = "Excellent"
		reussite = true
		commentaire = "Félicitations, résultat exceptionnel !"
	case note >= 80:
		mention = "Très bien"
		reussite = true
		commentaire = "Très bon travail, continuez ainsi !"
	case note >= 70:
		mention = "Bien"
		reussite = true
		commentaire = "Bon travail, avec des efforts supplémentaires vous pouvez faire mieux."
	case note >= 60:
		mention = "Passable"
		reussite = true
		commentaire = "Résultat juste suffisant, il faut travailler davantage."
	default:
		mention = "Insuffisant"
		reussite = false
		commentaire = "Résultat insuffisant, reprise nécessaire."
	}

	return // Retour automatique des variables nommées
}

// Fonction qui utilise la logique complexe d'analyse de notes
func logiqueComplexe() {
	notes := []int{95, 85, 75, 65, 45}

	for _, note := range notes {
		mention, reussite, commentaire := analyserNote(note)
		fmt.Printf("Note: %d - %s (Réussite: %v)\n", note, mention, reussite)
		fmt.Printf("Commentaire: %s\n\n", commentaire)
	}
}

// Fonction qui retourne une fonction
func creerOperateur(operation string) func(int, int) int {
	switch operation {
	case "+":
		return func(a, b int) int { return a + b }
	case "-":
		return func(a, b int) int { return a - b }
	case "*":
		return func(a, b int) int { return a * b }
	default:
		return func(a, b int) int { return 0 }
	}
}

// Fonction qui utilise les fonctions retournées
func retourFonctions() {
	// Créer différents opérateurs
	addition := creerOperateur("+")
	soustraction := creerOperateur("-")
	multiplication := creerOperateur("*")

	// Utiliser les fonctions retournées
	fmt.Println("5 + 3 =", addition(5, 3))
	fmt.Println("5 - 3 =", soustraction(5, 3))
	fmt.Println("5 * 3 =", multiplication(5, 3))
}
