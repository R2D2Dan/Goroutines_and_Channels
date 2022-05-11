package main

import (
	"log"
)

func main() {

	Mine := [5]string{"ore", "earth", "earth", "ore", "ore"}
	MetalChanel := make(chan string, 5)
	IngotChanel := make(chan string, 5)
	Done := make(chan int)

	go Find(&Mine, &MetalChanel)
	go Melt(&MetalChanel, &IngotChanel)
	go Ingot(&IngotChanel, &Done)

	<-Done

}

func Find(Mine *[5]string, MetalChanel *chan string) {
	for _, m := range Mine {
		if m == "ore" {
			log.Printf("Find: ore")
			*MetalChanel <- "ore"
		}
	}
	close(*MetalChanel)
	log.Println("Fine: Done")

}

func Melt(MetalChanel *chan string, IngotChanel *chan string) {

	for m := range *MetalChanel {
		if m == "ore" {
			log.Println("Melt: ore -> metal")
			*IngotChanel <- "metal"
		}
	}
	close(*IngotChanel)
	log.Println("Melt: Done")

}

func Ingot(IngotChanel *chan string, Done *chan int) {

	for m := range *IngotChanel {
		if m == "metal" {
			log.Println("Ingot: metal -> ingot")
		}
	}
	log.Println("Ingot: Done")

	*Done <- 0

}
