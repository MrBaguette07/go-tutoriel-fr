// package main qui créer un programme qui :
// Crée une map pour stocker les notes d'étudiants (nom → note)
// Ajoute 5 étudiants avec leurs notes
// Calcule et affiche la moyenne des notes
// Trouve et affiche l'étudiant avec la meilleure note
package main

func main() {
	notes := make(map[string]float64)
	notes["Alice"] = 85.5
	notes["Bob"] = 92.0
	notes["Charlie"] = 78.0
	notes["David"] = 88.5
	notes["Eve"] = 91.0

	moyenne, meilleurEtudiant := analyserNotes(notes)
	println("Moyenne des notes:", moyenne)
	println("Meilleur étudiant:", meilleurEtudiant)
}

func analyserNotes(notes map[string]float64) (float64, string) {
	if len(notes) == 0 {
		return 0, ""
	}
	var somme float64
	var meilleurEtudiant string
	var meilleureNote float64

	// Parcours de la map pour calculer la somme et trouver la meilleure note
	for etudiant, note := range notes {
		somme += note
		if note > meilleureNote {
			meilleureNote = note
			meilleurEtudiant = etudiant
		}
	}

	// Calcul de la moyenne
	moyenne := somme / float64(len(notes))
	return moyenne, meilleurEtudiant
}

// Exercice 1 : OK (20)
