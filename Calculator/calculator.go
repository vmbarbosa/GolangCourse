package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type estructura struct{}

func (estructura) operacion(entrada string) int {
	valores := strings.Split(entrada, " ")
	op1, _ := strconv.Atoi(valores[0])
	op2, _ := strconv.Atoi(valores[2])
	switch valores[1] {
	case "+":
		return op1 + op2
	case "-":
		return op1 + op2
	case "/":
		return op1 + op2
	case "*":
		return op1 + op2
	default:
		fmt.Println("El siguiente operador no se encuentra soportado: ", valores[1])
		return 0
	}
}

func leer() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println("Digite la operación que desea realizar")
	estruc := estructura{}
	entrada := leer()
	fmt.Println("Resultado: ")
	fmt.Println(estruc.operacion(entrada))
}
