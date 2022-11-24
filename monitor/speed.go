package monitor

import (
	"fmt"

	"github.com/adhocore/fast"
)

func GetSpeed() {
	noUpload := false

	// Prints the output right away:
	// fast.Run(noUpload)
	res, err := fast.Measure(noUpload)

	if err != nil {
		fmt.Println("failed to measure speed:", err)
	}

	fmt.Println(res.Down)
}
