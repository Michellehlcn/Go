
import (
	"fmt"
	"strconv"
	"strings"
)

func calculate(input1, input2 string) float64 {
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
}
