package main

import (
	"fmt"
)

func main() {
	escreverMenuPrincipal()
	var opcaoEscolhida int
	_, err := fmt.Scanf("%d", &opcaoEscolhida)
	if err != nil {
		fmt.Println(err.Error())
	}
	selecionaOpcaoEscolhida(opcaoEscolhida)
}

func escreverMenuPrincipal() {
	fmt.Println("\n =================================================================================== ")
	fmt.Println(" ========================= Entendendo Algoritmos em Golang ========================= ")
	fmt.Println(" =================================================================================== ")
	fmt.Println("    === Qual opção gostaria de executar?")
	fmt.Println("\n    >>> (2) - Brincando com listas")
	fmt.Println("    >>> (1) - Voltar")
	fmt.Println("    >>> (0) - Sair")
	fmt.Println(" =================================================================================== ")
	fmt.Print("    >>> Opção: ")
}

func selecionaOpcaoEscolhida(opcaoEscolhida int) {
	switch opcaoEscolhida {
	case 0:
		fmt.Println("    >>> Saindo...")
	case 1:
		fmt.Println("    >>> Voltando...")
	case 2:
		fmt.Println("    >>> Brincando com listas!")
	default:
		fmt.Println("    >>> Opção invalida!")
	}
}
