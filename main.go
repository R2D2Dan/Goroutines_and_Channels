package main

import (
	"log"
)

func main() {

	Mine := [5]string{"ore", "earth", "earth", "ore", "ore"}
	MetalChanel := make(chan string, 5)
	IngotChanel := make(chan string, 5)
	Done := make(chan int)

	go Find(Mine, MetalChanel)
	go Melt(MetalChanel, IngotChanel)
	go Ingot(IngotChanel, Done)

	<-Done
}

func Find(Mine [5]string, MetalChanel chan string) {
	for _, m := range Mine {
		if m == "ore" {
			log.Printf("Find: ore")
			MetalChanel <- "ore"
		}
	}
	log.Println("Fine: Done")
	close(MetalChanel)
}

func Melt(MetalChanel chan string, IngotChanel chan string) {

	for m := range MetalChanel {
		if m == "ore" {
			log.Println("Melt: ore -> metal")
			IngotChanel <- "metal"
		}
	}
	log.Println("Melt: Done")
	close(IngotChanel)
}

func Ingot(IngotChanel chan string, Done chan int) {

	for m := range IngotChanel {
		if m == "metal" {
			log.Println("Ingot: metal -> ingot")
		}
	}
	log.Println("Ingot: Done")
	Done <- 0

}
