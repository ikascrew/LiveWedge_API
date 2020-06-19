package main

import (
	"fmt"

	"github.com/cerevo/LiveWedge_API/libvsw/v2"
)

func main() {
	items := libvsw.Find()
	fmt.Printf("Found %d LiveWedge(s).\n", len(items))
	for _, v := range items {
		fmt.Printf("%s\t%s\n", v.Address, v.DisplayName)
	}
}
