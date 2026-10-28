package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	indexB    int
	targetA   int
	costA     int
	costB     int
	totalCost int
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]
	mot := strings.Fields(arg)
	var liste []int
	var listef []int
	for _, c := range mot {
		n, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("Error")
			return
		}
		liste = append(liste, n)
	}
	//doublon
	for i := 0; i < len(liste); i++ {
		for j := i + 1; j < len(liste); j++ {
			if liste[i] == liste[j] {
				fmt.Println("Error")
				return
			}
		}
	}
	//ranger
	for i := 0; i < len(liste); i++ {
		rank := 0

		for j := 0; j < len(liste); j++ {
			if liste[j] < liste[i] {
				rank++
			}
		}

		listef = append(listef, rank)
	}
	if isSorted(listef) {
		return
	}
	var liste2 []int
	var moves []string

	turk(&listef, &liste2, &moves)

	fmt.Println(strings.Join(moves, " "))
}

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}

	return true
}

// Retourne le nombre de rotations nécessaires.
// Un coût positif signifie ra/rb, un coût négatif rra/rrb.
func rotationCost(index, size int) int {
	if index <= size/2 {
		return index
	}
	return index - size
}

// Trouve où insérer x dans A pour conserver l'ordre circulaire.
func findTarget(a []int, x int) int {
	if len(a) == 0 {
		return 0
	}

	target := -1
	for i, value := range a {
		if value > x {
			if target == -1 || value < a[target] {
				target = i
			}
		}
	}

	// Si x est supérieur à tous les éléments,
	// il sera placé avant le plus petit élément.
	if target == -1 {
		for i := range a {
			if a[i] < a[targetMin(a)] {
				target = i
			}
		}
	}

	return target
}

func targetMin(a []int) int {
	min := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[min] {
			min = i
		}
	}
	return min
}

// Calcule le coût total d'un déplacement de B vers A.
func calculateCost(a, b []int, indexB int) Move {
	targetA := findTarget(a, b[indexB])

	costA := rotationCost(targetA, len(a))
	costB := rotationCost(indexB, len(b))

	total := 0

	// Les rotations dans le même sens sont combinées.
	if costA > 0 && costB > 0 {
		if costA > costB {
			total = costA
		} else {
			total = costB
		}
	} else if costA < 0 && costB < 0 {
		if -costA > -costB {
			total = -costA
		} else {
			total = -costB
		}
	} else {
		total = abs(costA) + abs(costB)
	}

	return Move{
		indexB:    indexB,
		targetA:   targetA,
		costA:     costA,
		costB:     costB,
		totalCost: total,
	}
}

func bestMove(a, b []int) Move {
	best := calculateCost(a, b, 0)

	for i := 1; i < len(b); i++ {
		move := calculateCost(a, b, i)

		if move.totalCost < best.totalCost {
			best = move
		}
	}

	return best
}

func executeMove(a, b *[]int, moves *[]string, m Move) {
	costA := m.costA
	costB := m.costB

	// Rotations simultanées vers le haut.
	for costA > 0 && costB > 0 {
		commande("rr", a, b, moves)
		costA--
		costB--
	}

	// Rotations simultanées vers le bas.
	for costA < 0 && costB < 0 {
		commande("rrr", a, b, moves)
		costA++
		costB++
	}

	// Terminer les rotations de A.
	for costA > 0 {
		commande("ra", a, b, moves)
		costA--
	}
	for costA < 0 {
		commande("rra", a, b, moves)
		costA++
	}

	// Terminer les rotations de B.
	for costB > 0 {
		commande("rb", a, b, moves)
		costB--
	}
	for costB < 0 {
		commande("rrb", a, b, moves)
		costB++
	}

	// Placer le sommet de B dans A.
	commande("pa", a, b, moves)
}

func sortThree(a, b *[]int, moves *[]string) {
	if isSorted(*a) {
		return
	}

	if len(*a) != 3 {
		return
	}

	x, y, z := (*a)[0], (*a)[1], (*a)[2]

	switch {
	case x > y && y < z && x < z:
		commande("sa", a, b, moves)

	case x > y && y > z:
		commande("sa", a, b, moves)
		commande("rra", a, b, moves)

	case x > y && y < z && x > z:
		commande("ra", a, b, moves)

	case x < y && y > z && x < z:
		commande("sa", a, b, moves)
		commande("ra", a, b, moves)

	case x < y && y > z && x > z:
		commande("rra", a, b, moves)
	}
}

func turk(a, b *[]int, moves *[]string) {
	if len(*a) <= 3 {
		sortThree(a, b, moves)
		return
	}

	// Garder trois éléments dans A.
	for len(*a) > 3 {
		commande("pb", a, b, moves)
	}

	sortThree(a, b, moves)

	// Réinsérer les éléments de B dans A.
	for len(*b) > 0 {
		m := bestMove(*a, *b)
		executeMove(a, b, moves, m)
	}

	// Remettre le plus petit rang au sommet de A.
	minIndex := targetMin(*a)
	cost := rotationCost(minIndex, len(*a))

	for cost > 0 {
		commande("ra", a, b, moves)
		cost--
	}
	for cost < 0 {
		commande("rra", a, b, moves)
		cost++
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func commande(mot string, liste1, liste2 *[]int, lmot *[]string) {
	switch mot {
	case "pa":
		if len(*liste2) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		*liste1 = append([]int{(*liste2)[0]}, *liste1...)
		*liste2 = (*liste2)[1:]
		*lmot = append(*lmot, "pa")
	case "pb":
		if len(*liste1) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		*liste2 = append([]int{(*liste1)[0]}, *liste2...)
		*liste1 = (*liste1)[1:]
		*lmot = append(*lmot, "pb")
	case "sa":
		if len(*liste1) < 2 {
			fmt.Println("Error")
			os.Exit(0)
		}
		(*liste1)[0], (*liste1)[1] = (*liste1)[1], (*liste1)[0]
		*lmot = append(*lmot, "sa")
	case "sb":
		if len(*liste2) < 2 {
			fmt.Println("Error")
			os.Exit(0)
		}
		(*liste2)[0], (*liste2)[1] = (*liste2)[1], (*liste2)[0]
		*lmot = append(*lmot, "sb")
	case "ss":
		if len(*liste2) < 2 || len(*liste1) < 2 {
			fmt.Println("Error")
			os.Exit(0)
		}
		(*liste1)[0], (*liste1)[1] = (*liste1)[1], (*liste1)[0]
		(*liste2)[0], (*liste2)[1] = (*liste2)[1], (*liste2)[0]
		*lmot = append(*lmot, "ss")
	case "ra":
		if len(*liste1) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		first := (*liste1)[0]
		*liste1 = append((*liste1)[1:], first)
		*lmot = append(*lmot, "ra")
	case "rb":
		if len(*liste2) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		first := (*liste2)[0]
		*liste2 = append((*liste2)[1:], first)
		*lmot = append(*lmot, "rb")
	case "rr":
		if len(*liste1) == 0 || len(*liste2) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		first := (*liste1)[0]
		*liste1 = append((*liste1)[1:], first)
		first = (*liste2)[0]
		*liste2 = append((*liste2)[1:], first)
		*lmot = append(*lmot, "rr")
	case "rra":
		if len(*liste1) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		last := (*liste1)[len(*liste1)-1]
		*liste1 = append([]int{last}, (*liste1)[:len(*liste1)-1]...)
		*lmot = append(*lmot, "rra")
	case "rrb":
		if len(*liste2) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		last := (*liste2)[len(*liste2)-1]
		*liste2 = append([]int{last}, (*liste2)[:len(*liste2)-1]...)
		*lmot = append(*lmot, "rrb")
	case "rrr":
		if len(*liste1) == 0 || len(*liste2) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		last := (*liste1)[len(*liste1)-1]
		lastt := (*liste2)[len(*liste2)-1]
		*liste1 = append([]int{last}, (*liste1)[:len(*liste1)-1]...)
		*liste2 = append([]int{lastt}, (*liste2)[:len(*liste2)-1]...)
		*lmot = append(*lmot, "rrr")
	default:
		fmt.Println("Error")
		os.Exit(0)
	}
	return
}
