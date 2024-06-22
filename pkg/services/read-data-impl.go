package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sdblg/vrp/pkg/models"
	"github.com/sdblg/vrp/pkg/utils"
)

func (rd *service) Do() {
	// reading data from file line by line and drop into channel for calculating distances
	loads, err := readData(rd.cfg.FileName, rd.cfg.ChannelSize)
	utils.PanicIfErr(err)

	// processing the loads
	err = processData(loads)
	utils.PanicIfErr(err)

	printLoads(loads)
}

func printLoads(loads []*models.Load) {
	for _, l := range loads {
		if l.Joined {
			continue
		}
		fmt.Println(l.LoadNumbers)
	}
}

func readData(fname string, chanSize int) ([]*models.Load, error) {
	out := make(chan string, chanSize)
	go utils.ReadFile(fname, out)

	grids := make([]*models.Load, 0)
	for s := range out {
		if l := makeLoad(s); l != nil {
			grids = append(grids, l)
			// fmt.Printf("%s\n", l)
		}
	}
	return grids, nil
}

func makeLoad(s string) *models.Load {
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
		LoadNumbers: []string{arr[0]},
		Point: models.Point{
			Bx: bx,
			By: by,

			Ex: ex,
			Ey: ey,
		},
	}

	l.InitializeDistances()

	return l
}
