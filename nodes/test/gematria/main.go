package main

import (
	"fmt"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

func main() {
	input := "i am ozymandias king of kings"
	gematrias := core.CalculateAllGematrias(input)
	for _, g := range gematrias {
		fmt.Println("----------------")
		fmt.Println("name:", g.ShowName)
		fmt.Println("sum: ", g.Sum)
	}
	fmt.Println("----------------")
	format := core.FormatGematria(gematrias)

	fmt.Println("format:", format)
	fmt.Println("----------------")

	//now lets try to order it in objectives?
	sums := core.DecodeFGematrias(format)
	fmt.Println(sums)
}
