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
		if args[0] == "-i" || args[0] == "-ideal" {
			IdealNumbers()
			return nil
		}
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

func IdealNumbers() {
	fmt.Println("my suggested loads for given prop sizes, results in 12 - 12.5 disc load:")
	fmt.Println("this is all based on how my 5 inch feels with 4.9 sbang props at 605.5 g ")
	fmt.Println("my upper threshold is 12.5 disc load, all numbers are the highest round gram that results in under 12.5")
	fmt.Println("3.0 inch | low : 219 : high 228 g")
	fmt.Println("    load | 12.005612 : 12.498993 ")
	fmt.Println("3.5 inch | low : 298 : high 310 g")
	fmt.Println("    load | 12.002256 : 12.485568 ")
	fmt.Println("4.9 inch | low : 584 : high 608 g")
	fmt.Println("    load | 12.000612 : 12.493787 ")
	fmt.Println("5.0 inch | low : 609 : high 633 g")
	fmt.Println("    load | 12.018769 : 12.492415 ")
}

func (c DiscLoadCommand) Help() {
	fmt.Println("description:")
	fmt.Println("\tcalculates disc loading of a for a drone with a given all up weight in grams and prop size in inches")
	fmt.Println("\tI like a disc load of 12-12.5 for freestyle")
	fmt.Println("usage:")
	fmt.Println("\tdiscload [args] ")
	fmt.Println("\t[-p | -props]\t(float) \t: prop size in inches")
	fmt.Println("\t[-w | -weight]\t(float) \t: weight in grams")
	fmt.Println("\t[-i | -ideal]\t(flag)  \t: show my ideal ranges")
	fmt.Println("examples")
	fmt.Println("\tdiscload -p 3 -w 230")
	fmt.Println("\tdiscload -p 3.5 -w 300")
	fmt.Println("\tdiscload -prop 4.9 -weight 605.5")
}

func (c DiscLoadCommand) Name() string {
	return "discload"
}
