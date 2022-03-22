package main

func main() {
	// var randomValues interface{}
	// _ = randomValues

	// randomValues = "Jalan Sudirman"
	// randomValues = 20
	// randomValues = true
	// randomValues = []string{"raditeo", "warma"}

	// var v interface{}

	// v = 20
	// v = v * 9

	// var v interface{}

	// v = 20
	// if value, ok := v.(int); ok == true {
	// 	v = value * 9
	// }

	// fmt.Println(v)

	rs := []interface{}{1, "Raditeo", true, 2, "Warma", true}

	rm := map[string]interface{}{
		"Name":   "Raditeo",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
