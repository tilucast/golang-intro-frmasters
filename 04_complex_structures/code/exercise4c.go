// package main

// import "fmt"

// func calculateAverage(scores [5]float64) float64 {
// 	var average float64

// 	for _, v := range scores {
// 		average += v
// 	}

// 	return average / float64(len(scores))
// }

// func findAnimalOnMap(animalType string, actualMap map[string]string) bool{
// 	_, ok := actualMap[animalType]
// 	return ok
// }

// func appendShitToSlice(groceryList []string, newGroceries ...string) []string{
// 	newArray := append(groceryList, newGroceries...)
// 	fmt.Println(newArray)

// 	return newArray
// }

// func main() {
// 	scores := [5]float64{5, 2.2, 8, 9.3, 6.4}
// 	fmt.Println(calculateAverage(scores))

// 	pets := map[string]string{
// 		"cat": "penico",
// 		"dog": "jegue",
// 		"donkey": "peste",
// 	}

// 	fmt.Println(findAnimalOnMap("cat", pets))

// 	groceries := []string{"candy","othercandy","yetanothercandy"}

// 	result := appendShitToSlice(groceries, "chocolate","bagel","m&m's")
// 	fmt.Println("calling appendtoshit on groceries yields to", result)
// }