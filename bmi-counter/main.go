package main

import "fmt"

// BMICounter is a function that returns the BMI of a person.
func BMICounter(weight float32, height float32) float32 {
	return weight / (height * height)
}

// markHigherBMI is a function that returns true if Mark has higher BMI than John.
func markHigherBMI(markBMI float32, johnBMI float32) bool {
	return markBMI > johnBMI
}

func main() {
	// § Data 1: Berat Mark 78 kg dan tinggi 1.69 m. Berat John 92 kg dan tinggi 1.95 m.
	// § Data 2: Berat Mark 95 kg dan tinggi 1.88 m. Berat John 85 kg dan tinggi 1.76 m.

	data1 := map[string]map[string]float32{
		"Mark": {
			"weight": 78,
			"height": 1.69,
		},
		"John": {
			"weight": 92,
			"height": 1.95,
		},
	}
	markBMI := BMICounter(data1["Mark"]["weight"], data1["Mark"]["height"])
	johnBMI := BMICounter(data1["John"]["weight"], data1["John"]["height"])
	isMarkHigherBMI := markHigherBMI(markBMI, johnBMI)

	fmt.Println("Data 1")
	fmt.Println("Mark BMI:", markBMI)
	fmt.Println("John BMI:", johnBMI)
	fmt.Println("Mark has higher BMI than John: ", isMarkHigherBMI)

	data2 := map[string]map[string]float32{
		"Mark": {
			"weight": 95,
			"height": 1.88,
		},
		"John": {
			"weight": 85,
			"height": 1.76,
		},
	}
	markBMI = BMICounter(data2["Mark"]["weight"], data2["Mark"]["height"])
	johnBMI = BMICounter(data2["John"]["weight"], data2["John"]["height"])
	isMarkHigherBMI = markHigherBMI(markBMI, johnBMI)

	fmt.Println("\nData 2")
	fmt.Println("Mark BMI:", markBMI)
	fmt.Println("John BMI:", johnBMI)
	fmt.Println("Mark has higher BMI than John: ", isMarkHigherBMI)
}
