package main

import (
	"bufio"   //Leitura de input do usuário
	"fmt"     //Formatação e impressão
	"os"      // Interação com o sistema operacional (arquivos, etc.)
	"strings" // Manipulação de strings
)

func main() {

	//Imprime uma mensagem pedindo input do usuário.
	fmt.Println("Digite um texto (ou pressione Enter para ler de um arquivo 'texto.txt'): ")

	scanner := bufio.NewScanner(os.Stdin) //Prepara para ler do terminal (entrada padrão).
	scanner.Scan()                        //Aguarda o usuário digitar algo e pressionar Enter.
	input := scanner.Text()               //Pega o texto digitado e armazena em input.

	var text string

	//Se o usuário pressionar Enter sem digitar nada (input == ""), o programa tenta ler do arquivo texto.txt.
	if input == "" {
		fileContent, err := os.ReadFile("texto.txt") //os.ReadFile("texto.txt"): Lê o conteúdo do arquivo.
		if err != nil {
			fmt.Println("Erro ao ler o arquivo: ", err)
			return
		}
		text = string(fileContent)
	} else {
		text = input //Se o usuário digitou algo, usa esse texto (text = input).
	}

	words := strings.Fields(text) //Divide o texto em palavras (separadas por espaços, tabs ou quebras de linha).

	wordCount := make(map[string]int) //Cria um mapa (dicionário) para contar quantas vezes cada palavra aparece.
	//chave string : palavra repetida
	//valor int : contar a palavra repetida

	//Percorre cada palavra:
	for _, word := range words {

		//strings.ToLower(word): Converte para minúsculas (para evitar diferenciação entre "Go" e "go").
		wordCount[strings.ToLower(word)]++ //wordCount[palavra]++: Incrementa a contagem da palavra no mapa.

	}

	fmt.Printf("Total de palavras: %d\n", len(words))   //len(words): Número total de palavras.
	fmt.Printf("Palavras únicas: %d\n", len(wordCount)) //Número de palavras únicas (tamanho do mapa).

	maxWord := ""
	maxCount := 0

	//Percorre o mapa (wordCount) e compara as contagens.
	for word, count := range wordCount {

		//Se uma palavra tiver contagem maior que maxCount, atualiza maxWord e maxCount.
		if count > maxCount {
			maxWord = word
			maxCount = count
		}
	}

	fmt.Printf("Palavra mais repetida: '%s' (%d vezes)\n", maxWord, maxCount)

}
