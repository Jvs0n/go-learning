package main

import (
	"fmt"
	"strconv"
	"strings"
)

type House struct {
	address string
	rooms   int
}
type Person struct {
	name string
	*House
	hobbys []string
}

func main() {
	//Create Houses
	house1 := &House{address: "pointerville 1", rooms: 1}
	house2 := &House{address: "pointerville 2", rooms: 5}

	//Tom and Evi are moving to pointerville and buy house 1 (Evi) and 2 (Tom)
	person1 := &Person{name: "U Evi", House: house1, hobbys: []string{"running", "hiking"}}
	person2 := &Person{name: "Tom Cat", House: house2, hobbys: []string{"running", "hiking"}}

	persons := []*Person{person1, person2}

	//Evi picks up a new Hobby
	addHobby(person1, "flying a plane")

	//Evi and Tom switch houses
	swapHouse(person1, person2)

	//Evi doesnt like it in Toms house, so she moves back
	moveHouse(person1, house1)

	personsMap := make(map[string]*Person)
	for _, val := range persons {
		personsMap[val.name] = val
	}

	for key, value := range personsMap {
		fmt.Println(key)
		fmt.Println(value.House.address, value.House.rooms, strings.Join(value.hobbys, ", "))
	}
}

// Helpers
func information(people []*Person) {
	for _, val := range people {
		fmt.Println(val.name + ", " + val.House.address + ", " + strconv.Itoa(val.House.rooms) + ", " + strings.Join(val.hobbys, ", "))
	}
}

func addHobby(person *Person, hobby string) {
	person.hobbys = append(person.hobbys, hobby)
}

func swapHouse(person1 *Person, person2 *Person) {
	person1.House, person2.House = person2.House, person1.House
}

func moveHouse(person *Person, house *House) {
	person.House = house
}

// Day 3 - Pointers
