package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	readFile, err := os.Open("File.txt") //ouverture du fichier
	if err != nil {                      //si erreur
		log.Fatal(err) //on affiche l'erreur
	}

	fileScanner := bufio.NewScanner(readFile) //on scan le fichier
	fileScanner.Split(bufio.ScanLines)        //on scan ligne par ligne
	var lines []string                        //on crée un tableau de string
	for fileScanner.Scan() {                  //on scan le fichier
		lines = append(lines, fileScanner.Text()) //on ajoute chaque ligne au tableau
	}
	readFile.Close() //on ferme le fichier

	fmt.Println(lines[0])            //on affiche la première ligne du fichier
	fmt.Println(lines[len(lines)-1]) //on affiche la dernière ligne du fichier
	for index, mots := range lines { //on parcours le tableau
		if mots == "before" { //si on trouve le mot "before"
			v := lines[index+1]     //on récupère la ligne suivante
			a, _ := strconv.Atoi(v) //on convertit la ligne suivante en int
			fmt.Println(lines[a-1]) //on affiche la ligne correspondante

		}
		if mots == "now " { //si on trouve le mot "now"
			b := lines[index-1]         //on récupère la ligne précédente
			c := int(b[1]) / len(lines) //on calcule la valeur de la ligne précédente
			fmt.Println(lines[c-1])     //on affiche la ligne correspondante
		}

	}

	rand.Seed(time.Now().UnixNano()) // genere un nombre aleatoire via un changement de seed en fonction du temps
	fmt.Println(rand.Intn(100))      // genere un nombre aleatoire entre 0 et 100
}
