package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)

	countries["mexico"] = "Ciudad de México"
	countries["argentina"] = "Buenos Aires"

	fmt.Println(countries)
	fmt.Printf("%s \n\n", countries["mexico"])

	championship := map[string]int{
		"barcelona":    39,
		"real_madrid":  38,
		"chivas":       37,
		"boca_juniors": 30,
	}

	for team, points := range championship {
		fmt.Printf("\t%s - %d\n", team, points)
	}
	fmt.Println(championship)
	delete(championship, "real_madrid")
	fmt.Println(championship)

	points, exist := championship["chivas"]
	fmt.Printf("\nPuntaje: %d, Existe: %t\n", points, exist)
}
