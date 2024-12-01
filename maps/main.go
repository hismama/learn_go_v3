package main

import "fmt"

func main() {

	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}
	fmt.Println("The age of Henry was", am["Henry"])
	fmt.Println(am)
	fmt.Printf("%#v\n", am)
	
	an := make(map[string]int)
	an["Lucas"] = 28
	an["Steph"] = 37

	for k, v := range an {
		fmt.Println(k, v)
	}
	
	delete(an, "Steph")
	
	fmt.Println(an)
	fmt.Printf("%#v\n", an)

	v, ok := an["Georgey"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key didn't exist")
	}

	if v, ok = an["Lucas"]; !ok {
		fmt.Println("Key didn't exist")
		} else {
		fmt.Println("the value prints", v)
	}


	
}
