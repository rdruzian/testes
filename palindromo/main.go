/*
5 - Desafio Construir um script para identificar se o texto é palíndromo ou não: (remover acentos e espaços)
R: É um palíndromo;
R: É um palíndromo;
R: É um palíndromo;
R: Não é um palíndromo;
R: Não é um palíndromo;
R: Não é um palíndromo;
R: É um palíndromo;
R: É um palíndromo;
R: É um palíndromo;
*/

package main

import (
	"fmt"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	palindromoCheck := false
	frases := []string{"Ana",
		"A cara rajada da jararaca",
		"A cara rajad4a4 da jararaca",
		"A cara rajad45 da jararaca",
		"A cara rajada 45 da jararaca",
		"A cara rajada 55 da jararaca",
		"Aí, Lima falou: \"Olá, família!\".",
		"Amo Omã. Se Roma me tem amores, amo Omã!",
		"Me vê se a panela da moça é de aço, Madalena Paes, e vem.",
	}

	for _, v := range frases {
		fraseTratado := trataFrase(v)
		k := len(fraseTratado) - 1

		for i := 0; i < len(fraseTratado)-1; i++ {
			if i == k && fraseTratado[i] == fraseTratado[k] {
				break
			}
			if fraseTratado[i] == fraseTratado[k] {
				palindromoCheck = true
			} else {
				palindromoCheck = false
				break
			}
			k--
		}

		if palindromoCheck {
			fmt.Println("A frase é palindromo: ", v)
		} else {
			fmt.Println("A frase não é palindromo: ", v)
		}
	}
}

func trataFrase(frase string) string {
	frase = strings.ToLower(frase)
	frase = strings.Replace(frase, " ", "", -1)
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	frase, _, _ = transform.String(t, frase)
	frase = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(frase, "")
	return frase
}
