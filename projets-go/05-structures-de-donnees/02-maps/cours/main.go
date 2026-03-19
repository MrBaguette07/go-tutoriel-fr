package main

func main() {
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	maps := make(map[string]int)

	maps["Alice"] = 30
	maps["Bob"] = 25
	maps["Charlie"] = 35

	println("Âge de Alice:", ages["Alice"])
	println("Âge de Bob:", maps["Bob"])

}

// OrdreDonnees struct pour préserver l'ordre d'insertion des clés
type OrdreDonnees struct {
	ordre   []string
	donnees map[string]interface{}
}

// Ajouter une clé-valeur tout en préservant l'ordre d'insertion
func (od *OrdreDonnees) Ajouter(cle string, valeur interface{}) {
	if _, existe := od.donnees[cle]; !existe {
		od.ordre = append(od.ordre, cle)
	}
	od.donnees[cle] = valeur
}
