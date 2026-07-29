package command

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var _ = RegisterCommand(DiscLoadCommand{})

type DiscLoadCommand struct{}

func (c DiscLoadCommand) Run(args []string) error {
	var weight float64 = 0
	var propsize float64 = 0
	for index, arg := range args {
		if strings.Contains(arg, "-prop") || strings.Contains(arg, "-p") {
			p, err := strconv.ParseFloat(args[index+1], 64)
			if err != nil {
				fmt.Println("[error] error parsing argument supplied to -prop. Must be a number")
				return err
			}
			propsize = p
		}
		if strings.Contains(arg, "-weight") || strings.Contains(arg, "-w") {
			w, err := strconv.ParseFloat(args[index+1], 64)
			if err != nil {
				fmt.Println("[error] error parsing argument supplied to -prop. Must be a number")
				return err
			}
			weight = w
		}
	}

	load := CalculateLoad(weight, propsize)
	fmt.Printf("[done ] disk loading is : %f\n", load)

	return nil
}

func CalculateLoad(weight float64, props float64) float64 {
	var result float64 = 0
	var weightkg = weight * 0.001 // convert to kg
	var propsm = props * 0.0254   // convert to m

	result = weightkg / ((math.Pow(propsm/2, 2) * math.Pi) * 4)

	return result
	/*
		assume default vars are in inches and grams
		do we allow other units? if so how
		can we set default units in conf later?

		calculate disk load here

			weight - needs to be in KG
			diameter - needs to be in M (LUT for common prop sizes to M)

			load = weight / (((diameter / pi) * (diameter / pi) )x4)
			OR
			load = weight / (((diameter / pi) ^2 ) x4)
	*/
}

func (c DiscLoadCommand) Help() {

}

func (c DiscLoadCommand) Name() string {
	return "discload"
}
