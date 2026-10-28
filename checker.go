package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	comm := ""
	if len(os.Args) != 2 {
		return
	}
	liste := os.Args[1]
	var liste1 []int
	var liste2 []int
	listee := strings.Split(liste, " ")
	for _, c := range listee {
		n, err := strconv.Atoi(c)
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		}
		liste1 = append(liste1, n)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := scanner.Text()
		fmt.Println(instruction)
		comm += instruction + "\n"
	}
	comm = comm[:len(comm)-1]
	com := strings.Split(comm, "\n")
	fmt.Println(liste1)
	fmt.Println(liste2)
	for _, c := range com {
		Commande(string(c), &liste1, &liste2)
	}
	fmt.Println(liste1)
	fmt.Println(liste2)
	if isSorted(liste1) && len(liste2) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func Commande(mot string, liste1, liste2 *[]int) {
	switch mot {
	case "pa":
		if len(*liste2) == 0 {
			fmt.Println("Error")
			return
		}
		*liste1 = append([]int{(*liste2)[0]}, *liste1...)
		*liste2 = (*liste2)[1:]
	case "pb":
		if len(*liste1) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		*liste2 = append([]int{(*liste1)[0]}, *liste2...)
		*liste1 = (*liste1)[1:]
	case "sa":
		if len(*liste1) < 2 {
			fmt.Println("Error")
			os.Exit(0)
		}
		(*liste1)[0], (*liste1)[1] = (*liste1)[1], (*liste1)[0]
	case "sb":
		if len(*liste2) < 2 {
			fmt.Println("Error")
			os.Exit(0)
		}
		(*liste2)[0], (*liste2)[1] = (*liste2)[1], (*liste2)[0]
	case "ss":
		if len(*liste2) < 2 || len(*liste1) < 2 {
			fmt.Println("Error")
			os.Exit(0)
		}
		(*liste1)[0], (*liste1)[1] = (*liste1)[1], (*liste1)[0]
		(*liste2)[0], (*liste2)[1] = (*liste2)[1], (*liste2)[0]
	case "ra":
		first := (*liste1)[0]
		*liste1 = append((*liste1)[1:], first)
	case "rb":
		first := (*liste2)[0]
		*liste2 = append((*liste2)[1:], first)
	case "rr":
		first := (*liste1)[0]
		*liste1 = append((*liste1)[1:], first)
		first = (*liste2)[0]
		*liste2 = append((*liste2)[1:], first)
	case "rra":
		if len(*liste1) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		last := (*liste1)[len(*liste1)-1]
		*liste1 = append([]int{last}, (*liste1)[:len(*liste1)-1]...)
	case "rrb":
		if len(*liste2) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		last := (*liste2)[len(*liste2)-1]
		*liste2 = append([]int{last}, (*liste2)[:len(*liste2)-1]...)
	case "rrr":
		if len(*liste1) == 0 || len(*liste2) == 0 {
			fmt.Println("Error")
			os.Exit(0)
		}
		last := (*liste1)[len(*liste1)-1]
		*liste1 = append([]int{last}, (*liste1)[:len(*liste1)-1]...)
		last = (*liste2)[len(*liste2)-1]
		*liste2 = append([]int{last}, (*liste2)[:len(*liste2)-1]...)
	default:
		fmt.Println("Error")
		os.Exit(0)
	}
	return
}
