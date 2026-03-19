// package main qui créer un système pour construire des objets étape par étape en utilisant des closures.
package main

import "fmt"

func main() {

	constructeur := creerConstructeur()
	objet1 := constructeur([]string{"zonix.fr", "lakestart.fr", "seapalm.fr"})
	objet2 := constructeur([]string{"lakestart.fr", "seapalm.fr", "zonix.fr"})

	println("Objet 1:", len(objet1))
	println("Objet 2:", objet2["lakestart.fr"])

	fmt.Println("------------------- CORRECTION -------------------")

	// Utilisation du builder simplifié
	nom, age, email, tel, adresse, construire := creerPersonneBuilder()

	// Construction étape par étape
	nom("Alice Dupont")
	age(28)
	email("alice@example.com")
	tel("06 12 34 56 78")
	adresse("123 Rue de la Paix, Paris")

	// Construire l'objet final
	personne := construire()
	fmt.Printf("\nPersonne créée: %s\n", personne)

	// Créer une autre personne avec le même builder
	fmt.Println("\n=== Deuxième personne ===")
	nom2, age2, email2, _, _, construire2 := creerPersonneBuilder()

	nom2("Bob Martin")
	age2(35)
	email2("bob@example.com")
	// On ne définit pas téléphone et adresse

	personne2 := construire2()
	fmt.Printf("Deuxième personne: %s\n", personne2)

}

func creerConstructeur() func([]string) map[string]string {
	objet := make(map[string]string)

	constructeur := func(proprietes []string) map[string]string {
		for _, propriete := range proprietes {
			objet[propriete] = "seapalm.fr"
		}

		return objet
	}

	return constructeur
}

////////////////////////////// Correction //////////////////////////////

// Personne struct avec plusieurs propriétés
type Personne struct {
	Nom       string
	Age       int
	Email     string
	Telephone string
	Adresse   string
}

func (p Personne) String() string {
	return fmt.Sprintf("Personne{Nom: %s, Age: %d, Email: %s, Tel: %s, Adresse: %s}",
		p.Nom, p.Age, p.Email, p.Telephone, p.Adresse)
}

// Version simplifiée plus pratique
func creerPersonneBuilder() (func(string), func(int), func(string), func(string), func(string), func() Personne) {
	p := Personne{}

	nom := func(n string) { p.Nom = n }
	age := func(a int) { p.Age = a }
	email := func(e string) { p.Email = e }
	telephone := func(t string) { p.Telephone = t }
	adresse := func(a string) { p.Adresse = a }
	construire := func() Personne { return p }

	return nom, age, email, telephone, adresse, construire
}

// Exercice 3 : Pas OK (5)
