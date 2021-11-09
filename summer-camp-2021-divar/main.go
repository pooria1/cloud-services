package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	Saturday = 1
	Sunday   = 2
	Monday   = 3
)

type person struct {
	left  int
	right int

	name string
}

type desk struct {
	persons []person
	option  string
}

type floor struct {
	deskLists []desk
	floorType string
}

func main() {
	var (
		f                 int
		optionsPrice      []int
		n                 int
		specialFloorPrice int
	)

	var a string
	var b string

	fmt.Scanf("%s %s", &a, &b)

	len(f)

	fmt.Scanf("%d", f)
	for i := 0; i < f; i++ {
		var d int
		fmt.Scanf("%d", d)
		optionsPrice = append(optionsPrice, d)
	}
	fmt.Scanf("%d %d", &n, &specialFloorPrice)
	floors := make([]floor, n)
	floorsCap := make([]int, n)
	for i := 0; i < n; i++ {
		var ft string
		fmt.Scanf("%d %s", &floorsCap[i], &ft)
		var dl []desk
		for j := 0; j < floorsCap[i]; j++ {
			var o string
			fmt.Scanf("%s", &o)
			dl = append(dl, desk{
				option: o,
			})
		}
		floors[i] = floor{
			deskLists: dl,
			floorType: ft,
		}
	}

	for {
		var (
			command      [5]string
			l            int
			r            int
			tableType    string
			gotOrReserve string
		)
		fmt.Scanf("%s", &command[0])
		if command[0] == "end" {
			return
		}
		fmt.Scanf("%s %s %s %s", &command[1], &command[2], &command[3], &command[4])
		if command[1] == "reserve_desk" {
			l, _ = strconv.Atoi(command[3])
			r, _ = strconv.Atoi(command[4])
			tableType = "special"
			gotOrReserve = "reserved"
		} else {
			l, _ = strconv.Atoi(command[0])
			r, _ = strconv.Atoi(command[4])
			tableType = command[3]
			gotOrReserve = "got"
		}
		newPerson := person{
			left:  l,
			name:  command[2],
			right: l + r,
		}

		found := false
		for floorIndex := range floors {
			for deskIndex := range floors[floorIndex].deskLists {
				if tableType != floors[floorIndex].floorType {
					break
				}
				if len(floors[floorIndex].deskLists[deskIndex].persons) == 0 {
					fp := 0
					if tableType == "special" {
						fp = specialFloorPrice
					}
					fmt.Printf(
						"%s %s desk %d-%d for %d",
						newPerson.name,
						gotOrReserve,
						floorIndex+1,
						deskIndex+1,
						newPerson.getRequestPrice(fp, "a", optionsPrice),
					)
					floors[floorIndex].deskLists[deskIndex].persons = append(floors[floorIndex].deskLists[deskIndex].persons, newPerson)
					found = true
					break
				}

				tadakholFlag := false
				for personIndex := range floors[floorIndex].deskLists[deskIndex].persons {
					if tadakhol(newPerson, floors[floorIndex].deskLists[deskIndex].persons[personIndex]) {
						tadakholFlag = true
						break
					}
				}
				if !tadakholFlag {
					floors[floorIndex].deskLists[deskIndex].persons = append(floors[floorIndex].deskLists[deskIndex].persons, newPerson)
					fmt.Printf("%s %s desk %d-%d", newPerson.name, gotOrReserve, floorIndex+1, deskIndex+1)
					if tableType == "special" {
						fmt.Println(" for", specialFloorPrice)
					} else {
						fmt.Println()
					}
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			fmt.Println("no desk available")
		}
	}
}

func tadakhol(person1, person2 person) bool {
	if person1.right <= person2.left {
		return false
	}
	if person1.left >= person2.right {
		return false
	}
	return true

}

func (p person) getDuration() int {
	return p.right - p.left
}

func (p person) getRequestPrice(floorPrice int, option string, allPrice []int) int {
	optionPrice := 0
	requestedOps := strings.Split(option, "")
	for i := 0; i < len(requestedOps); i++ {
		if requestedOps[i] == "1" {
			optionPrice += allPrice[i]
		}
	}
	return floorPrice + optionPrice*p.getDuration()
}
