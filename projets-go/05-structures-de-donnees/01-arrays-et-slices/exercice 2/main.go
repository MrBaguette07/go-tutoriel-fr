// package main qui créer une fonction qui prend un slice d'entiers et retourne :
// Le plus petit élément
// Le plus grand élément
// La somme de tous les éléments
package main

func main() {
	nums := []int{3, 1, 4, 1, 5, 9}

	min, max, sum := analyserSlice(nums)

	println("Min:", min)
	println("Max:", max)
	println("Sum:", sum)
}

// Fonction qui analyse un slice d'entiers et retourne le min, max et la somme
func analyserSlice(nums []int) (int, int, int) {
	if len(nums) == 0 {
		return 0, 0, 0
	}

	min, max, sum := nums[0], nums[0], 0

	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}

		sum += num
	}

	return min, max, sum
}

// Exercice 2 : OK (20)
