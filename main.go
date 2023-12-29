package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)


func action() (int, string) {
	var quantCar int
	var action int
	var password string
	var newPass int

	fmt.Print("(1) Ver senhas Salvas\n(2) Criar uma nova senha\n>> ")
	fmt.Scan(&newPass)

	if newPass == 1 {
		passwords()
	} else {
		fmt.Print("Digite a Quantidade de Caracteres: ")
		fmt.Scan(&quantCar)

		password = GeraSenha(quantCar)
		fmt.Printf("Senha Gerada foi: %v \n", password)

		fmt.Print("\nDeseja salva senha?\n(1) Sim\n(2) Sair\n>> ")
		fmt.Scan(&action)
	}

	return action, password

}
func main() {
	action, password := action()
	switch action {
	case 1:
		SalvaTxt(password)
	case 2:
		break
	}

}

func GeraSenha(quant int) string {
	charArray := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "!", "@", "#", "$", "%", "&", "*", "(", ")", "_", "+", ".", "?"}
	var password string
	count := 0
	for count < quant {
		NumAleatorio := rand.Intn(len(charArray))
		password += charArray[NumAleatorio]
		count++
	}

	return password
}

func SalvaTxt(password string) {
	// Abre ou cria o arquivo password.txt
	file, err := os.OpenFile("Passwords.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("\nErro ao abrir/criar Arquivo", err)
	}
	// para fechar o arquivo quando função terminar de ser executada
	defer file.Close()

	// pegar os textos que já estão salvos
	scanner := bufio.NewScanner(file)
	var linha string
	for scanner.Scan() {
		linha = scanner.Text()
	}

	//  escreve a senha no arquivo
	_, err = file.WriteString(linha + "\n" + password)

	if err != nil {
		fmt.Println("\nErro ao escrever salvar no arquivo", err)
	} else {
		fmt.Println("\nSua senha salva no Arquivo passwords.txt")
	}

}

func passwords() {
	file, err := os.Open("Passwords.txt")
	if err != nil {
		fmt.Println("\nErro ao abrir/criar Arquivo", err)
	}
	// para fechar o arquivo quando função terminar de ser executada
	defer file.Close()

	// pegar os textos que já estão salvos
	scanner := bufio.NewScanner(file)
	var linha []string
	for scanner.Scan() {
		linha = append(linha, scanner.Text())
	}
	for i := 1; i < len(linha); i++ {
		fmt.Printf("Senha %d: %v\n", i, linha[i])
	}

}
