package main

import "fmt"

type Arquivo struct {
	nome       string
	tamanho    float64
	caracteres int
	palavras   int
	linhas     int
}

func (arq *Arquivo) TamanhoMedioDePalavra() float64 {
	return float64(arq.caracteres) / float64(arq.palavras)
}

func (arq *Arquivo) MediaDePalavraPorLinha() float64 {
	return float64(arq.palavras) / float64(arq.linhas)
}

func main() {
	arquivo := Arquivo{"artigo.txt", 12.68, 12986, 1862, 220}

	fmt.Printf("Média de palavras por linha: %.2f\n", arquivo.MediaDePalavraPorLinha())
	fmt.Printf("Tamanho médio de palavra: %.2f\n", arquivo.TamanhoMedioDePalavra())
}
