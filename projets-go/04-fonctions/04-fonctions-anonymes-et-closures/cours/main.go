package main

import "fmt"

func main() {
	// Déclarer ET appeler immédiatement une fonction anonyme
	func() {
		fmt.Println("Je suis une fonction anonyme !")
	}() // Les () à la fin l'exécutent immédiatement

	// Fonction anonyme avec paramètres
	func(nom string) {
		fmt.Printf("Bonjour %s depuis une fonction anonyme !\n", nom)
	}("Alice")

	// Fonction anonyme avec retour
	resultat := func(a, b int) int {
		return a + b
	}(5, 3)

	fmt.Printf("5 + 3 = %d\n", resultat)

	// Créer un compteur
	compteur1 := creerCompteur()

	// Chaque appel incrémente et retourne la valeur
	fmt.Println("Compteur1:", compteur1()) // 1
	fmt.Println("Compteur1:", compteur1()) // 2
	fmt.Println("Compteur1:", compteur1()) // 3

	// Créer un second compteur indépendant
	compteur2 := creerCompteur()
	fmt.Println("Compteur2:", compteur2()) // 1 (indépendant du premier)
	fmt.Println("Compteur1:", compteur1()) // 4 (continue sa séquence)
	fmt.Println("Compteur2:", compteur2()) // 2

	// Créer différents multiplicateurs
	doubler := creerMultiplicateur(2)
	tripler := creerMultiplicateur(3)
	parDix := creerMultiplicateur(10)

	// Utiliser les closures
	fmt.Printf("Doubler 5: %d\n", doubler(5)) // 10
	fmt.Printf("Tripler 4: %d\n", tripler(4)) // 12
	fmt.Printf("5 × 10: %d\n", parDix(5))     // 50

	// On peut les utiliser avec des variables
	nombres := []int{1, 2, 3, 4, 5}
	fmt.Print("Doublés: ")
	for _, n := range nombres {
		fmt.Printf("%d ", doubler(n))
	}
	fmt.Println()

	fmt.Println("--------------------------------")
	// Générateur de nombres pairs
	pairs := creerGenerateurSuite(0, 2)
	fmt.Print("Nombres pairs: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", pairs())
	}
	fmt.Println()

	// Générateur de multiples de 5
	multiples5 := creerGenerateurSuite(5, 5)
	fmt.Print("Multiples de 5: ")
	for i := 0; i < 4; i++ {
		fmt.Printf("%d ", multiples5())
	}
	fmt.Println()

	// Générateur de nombres impairs en partant de 1
	impairs := creerGenerateurSuite(1, 2)
	fmt.Print("Nombres impairs: ")
	for i := 0; i < 6; i++ {
		fmt.Printf("%d ", impairs())
	}
	fmt.Println()
}

// Fonction qui crée un compteur utilisant une closure pour maintenir l'état entre les appels.
func creerCompteur() func() int {
	compte := 0 // Variable locale à creerCompteur

	// Retourner une fonction anonyme qui "capture" compte
	return func() int {
		compte++ // Modifie la variable capturée
		return compte
	}
}

func creerMultiplicateur(facteur int) func(int) int {
	// facteur est capturé par la closure
	return func(nombre int) int {
		return nombre * facteur
	}
}

func creerGenerateurSuite(debut, increment int) func() int {
	valeur := debut - increment // On commence avant pour que le premier appel donne 'debut'

	return func() int {
		valeur += increment
		return valeur
	}
}
