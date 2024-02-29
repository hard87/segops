package main

import (
	"fmt"
)

func main() {
	
	var opcaoEscolhida int = 1

	for opcaoEscolhida != 0 {
		escreverMenuPrincipal()
		_, err := fmt.Scanln(&opcaoEscolhida)
		
	    if err != nil {
			opcaoEscolhida = 0
			fmt.Println(err.Error())
			
		}else {
			selecionaOpcaoEscolhida(opcaoEscolhida)
		}
	}
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
		fmt.Println("    >>> Não é uma opção valida, digite uma das opções abaixo!")
	}
}
