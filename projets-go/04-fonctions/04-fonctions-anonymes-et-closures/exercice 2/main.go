// package main qui créer un système qui permet d'enregistrer des fonctions de rappel et de les exécuter toutes ensemble.
package main

import "fmt"

func main() {
	enregistrer, executer := creerGestionnaireCallbacks()

	// Enregistrer quelques callbacks
	enregistrer(func() {
		println("zonix.fr")
	})

	enregistrer(func() {
		println("Lakestart.fr")
	})

	enregistrer(func() {
		println("Seapalm.fr")
	})

	message := "VIVE SEAPALM ATMO BY SEAPALM"
	enregistrer(func() {
		fmt.Printf("Message: %s\n", message)
	})

	// Exécuter toutes les callbacks enregistrées
	executer()

	message = "je voudrais que Seapalm soit multimillinaire"

	executer()
}

func creerGestionnaireCallbacks() (func(func()), func()) {
	callbacks := []func(){}

	// Fonction pour enregistrer une callback
	enregistrer := func(cb func()) {
		callbacks = append(callbacks, cb)
	}

	// Fonction pour exécuter toutes les callbacks enregistrées
	executer := func() {
		for _, cb := range callbacks {
			cb()
		}
	}

	return enregistrer, executer
}

// Exercice 2 : OK (20)
