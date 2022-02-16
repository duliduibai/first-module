package math

func Add(params []float32)float32{
	var total float32
	for _, param := range params {
		total += param
	}
	return total
}
