package main

import (
	"fmt"
	"sort"
	"types"
	"writefile"
)

var choiche int

/*Biblioteca in formato testuale*/
func menu() {
	fmt.Println("\n1: Inserisci")
	fmt.Println("2: Cerca")
	fmt.Println("3: Modifica")
	fmt.Println("4: Cancella")
	fmt.Println("5: Elenco")
	fmt.Println("0: Esci")
	_, err := fmt.Scan(&choiche)

	fmt.Println(choiche)
	if err == nil {
		switch choiche {
		case 1:
			fmt.Println("Inserisci")
			inserisci()
		case 2:
			fmt.Println("Cerca")
			cerca()
		case 3:
			fmt.Println("Modifica")
		case 4:
			fmt.Println("Cancella")
		case 5:
			fmt.Println("Elenco")
			elencoMap()
		case 0:
			fmt.Println("Esci")
		default:
			panic("Scelta non valida")
		}
	} else {
		panic(err)
	}

}

func main() {
	menu()
	for choiche != 0 {
		menu()
	}
}

func inserisci() {
	var titolo, autore string
	b := &types.Book{}
	fmt.Println("\nTitolo")
	fmt.Scan(&titolo)
	fmt.Println("Autore")
	fmt.Scan(&autore)
	b.Titolo = titolo
	b.Autore = autore
	writefile.WriteBookToFile(b)
	fmt.Println("Inserimento Completato\n\n")
}

func elenco() {
	writefile.ElencoLibri()
}

func elencoMap() {
	books := writefile.ElencoLibri2()

	var keys []int
	for k := range books {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for key := range keys {
		fmt.Println("ID:", key, "\nTitolo:", books[key].Titolo, "\nAutore:", books[key].Autore)
	}

}

func cerca() {
	var chiave string
	fmt.Println("Titolo oppure Autore")
	fmt.Scanf("\n%s", &chiave)
	fmt.Println()
	books := writefile.ElencoLibri2()

	var keys []int
	for k := range books {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var booksV []*types.Book
	for key := range keys {
		if books[key].Titolo == chiave || books[key].Autore == chiave {
			booksV = append(booksV, books[key])

		}
	}
	for i := 0; i < len(booksV); i++ {
		fmt.Println(booksV[i])
	}

}
