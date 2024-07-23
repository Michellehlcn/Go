
import (
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Simple calculate

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

// convert a slice of strings to a map
func convertToMap(items []string) map[string]float64 {
	result := make(map[string]float64)

	elementValue := 10 / float64(len(items))
	for _, fruit := range items {
		result[fruit] = elementValue
	}
	return result
}

// Calculate shopping cart

type cartItem struct {
	name     string
	price    float64
	quantity int
}

func calculateTotal(cart []cartItem) float64 {
	var total float64 = 0
	for _, item := range cart {
		total += item.price * float64(item.quantity)
	}
	return total
}

// Advanced calculator
func calculateAdvanced(input1, input2, operation string) float64 {

	var result float64
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
	case "+":
		result = addValues(value1, value2)
	case "-":
		result = subtractValues(value1, value2)
	case "*":
		result = multiplyValues(value1, value2)
	case "/":
		result = divideValues(value1, value2)
	default:
		message := fmt.Sprintf("Unknown operation %s", operation)
		panic(message)
	}

}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", input)
		panic(message)
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	if value2 == 0 {
		panic("Cannot divide by zero")
	}
	return value1 / value2
}

// parse JSON-formated text

type cartItem_ struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func parseJSON(jsonString string) []cartItem_ {
	var cart []cartItem_
	err := json.Unmarshal([]byte(jsonString), &cart)
	if err != nil {
		panic(err)
	}
	return cart
}

// parse json from url

func parseJSONURL(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Println(content)

	rours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Printf("%s (%s)\n", tour.Name, tour.Price)
	}
}
func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}
		tours = append(tours, tour)
	}
}

type Tour struct {
	Name, Price string
}