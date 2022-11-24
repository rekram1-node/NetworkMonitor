package monitor

import (
	"fmt"

	"github.com/adhocore/fast"
)

func GetSpeed() {
	noUpload := true
	res, err := fast.Measure(noUpload)

	if err != nil {
		fmt.Println("failed to measure speed:", err)
	}

	fmt.Println("Download Speed:", res.Down)
}
