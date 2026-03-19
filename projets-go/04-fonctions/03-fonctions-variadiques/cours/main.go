package main

import (
	"fmt"
	"time"
)

// Fonction avec paramètres normaux ET variadiques
func saluerPlusieurs(salutation string, noms ...string) {
	if len(noms) == 0 {
		fmt.Println("Personne à saluer !")
		return
	}

	fmt.Printf("%s ", salutation)
	for i, nom := range noms {
		if i == len(noms)-1 {
			fmt.Printf("%s !\n", nom)
		} else if i == len(noms)-2 {
			fmt.Printf("%s et ", nom)
		} else {
			fmt.Printf("%s, ", nom)
		}
	}
}

func main() {
	saluerPlusieurs("Bonjour")
	saluerPlusieurs("Bonjour", "Alice")
	saluerPlusieurs("Bonjour", "Alice", "Bob")
	saluerPlusieurs("Bonjour", "Alice", "Bob", "Charlie", "Diana")

	afficherTout("Hello", 42, 3.14, true, []int{1, 2, 3})

	log("INFO", "Application démarrée")
	log("DEBUG", "Valeur de x:", 42)
	log("ERROR", "Erreur dans", "la fonction", "calculer", "avec code", 500)

	// Avec des variables
	utilisateur := "Alice"
	score := 95.5
	log("INFO", "Utilisateur", utilisateur, "a obtenu le score", score)
}

func afficherTout(elements ...interface{}) {
	fmt.Printf("J'ai reçu %d élément(s):\n", len(elements))

	for i, element := range elements {
		fmt.Printf("  [%d] %v (type: %T)\n", i, element, element)
	}
}

func log(niveau string, messages ...interface{}) {
	// Créer un timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Afficher le niveau et le timestamp
	fmt.Printf("[%s] %s: ", timestamp, niveau)

	// Afficher tous les messages
	fmt.Println(messages...)
}

// En gros les fonctions variadiques permettent de passer un nombre variable d'arguments à une fonction, ce qui est très pratique pour des opérations comme la somme, la concaténation, etc.
