package main

import "fmt"

func main() {

	pasta_boloñesa := 30
	ensalada_italiana := 20
	pizza_hawaiana := 10
	pizza_vegana := 15
	pasta_marinera := 25
	sandwich_tres_quesos := 8
	sandwich_vegetariano := 10
	ensalada_tradicional := 5

	var seguir string
	var orden int
	valor_total := 0

	fmt.Println("Bienvenido al FuFu Restaurant")
	menu_principal := ` Menú del día
		1. Pasta boloñesa
		2. Ensalada italiana
		3. Pizza hawaiana
		4. Pizza vegana
		5. Pasta marinera
		6. Sandwich tres quesos
		7. Sandwich vegetariano
		8. Ensalada tradicional`

	for {

		fmt.Println(menu_principal)
		fmt.Println("¿Qué desea ordenar?")
		_, err := fmt.Scanf("%d", &orden)
		if err != nil {
			panic(err)
		}
		if orden == 1 {
			valor_total = valor_total + pasta_boloñesa
			fmt.Println("¿Desea seguir ordenando?")
			_, err := fmt.Scanf("%s", &seguir)
			if err != nil {
				panic(err)
			}
		}
		if orden == 2 {
			valor_total = valor_total + ensalada_italiana
			fmt.Println("¿Desea seguir ordenando?")
			_, err := fmt.Scanf("%s", &seguir)
			if err != nil {
				panic(err)
			}
		}
		if orden == 3 {
			valor_total = valor_total + pizza_hawaiana
			fmt.Println("¿Desea seguir ordenando?")
			_, err := fmt.Scanf("%s", &seguir)
			if err != nil {
				panic(err)
			}
		}
		if orden == 4 {
			valor_total = valor_total + pizza_vegana
			fmt.Println("¿Desea seguir ordenando?")
			_, err := fmt.Scanf("%s", &seguir)
			if err != nil {
				panic(err)
			}
		}
		if orden == 5 {
			valor_total = valor_total + pasta_marinera
			fmt.Println("¿Desea seguir ordenando?")
			_, err := fmt.Scanf("%s", &seguir)
			if err != nil {
				panic(err)
			}
		}
		if orden == 6 {
			valor_total = valor_total + sandwich_tres_quesos
			fmt.Println("¿Desea seguir ordenando?")
			_, err := fmt.Scanf("%s", &seguir)
			if err != nil {
				panic(err)
			}
		}
		if orden == 7 {
			valor_total = valor_total + sandwich_vegetariano
			fmt.Println("¿Desea seguir ordenando?")
			_, err := fmt.Scanf("%s", &seguir)
			if err != nil {
				panic(err)
			}
		}
		if orden == 8 {
			valor_total = valor_total + ensalada_tradicional
			fmt.Println("¿Desea seguir ordenando?")
			_, err := fmt.Scanf("%s", &seguir)
			if err != nil {
				panic(err)
			}
		}
		if seguir != "SI" {
			fmt.Println("El valor a pagar es de $", valor_total)
			break
		}

	}

}
