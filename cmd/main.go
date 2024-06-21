package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sdblg/vrp/pkg/models"
	"github.com/sdblg/vrp/pkg/utils"
)

var Version string

func main() {
	if Version == "" {
		Version = "local"
	}
	begin := time.Now()
	defer fmt.Println(time.Since(begin))

	fmt.Println("Current version is", Version)
	// TODO make it argument using flag package
	fileName := "./test-data/TrainingProblems/problem20.txt"
	out := make(chan string, 5)
	go utils.ReadFile(fileName, out)

	grids := make([]*models.Load, 0)
	for s := range out {
		if l := MakeLoad(s); l != nil {
			grids = append(grids, l)
			fmt.Printf("%s\n", l)
		}
	}	
}

func MakeLoad(s string) *models.Load {
	arr := strings.Split(s, " ")
	if arr[0] == "loadNumber" {
		return nil
	}
	b := strings.Split(arr[1], ",")
	bx, _ := strconv.ParseFloat(b[0][1:], 64)	
	by, _ := strconv.ParseFloat(b[1][:len(b[1])-1], 64)

	e := strings.Split(arr[2], ",")
	ex, _ := strconv.ParseFloat(e[0][1:], 64)	
	ey, _ := strconv.ParseFloat(e[1][:len(e[1])-1], 64)

	l := &models.Load{
		LoadNumber: arr[0],
		Point: models.Point{
			Bx: bx,
			By: by,

			Ex: ex,
			Ey: ey,
		},
	}

	l.CalculateDistances()

	return l
}
