package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	data := `{
		"Thinly Sliced Peeled Apples": "6 Cups",
		"sugar": "3/4 cup",
		"flour": "2 tablespoons",
		"cinamon": "1/4 teaspoon",
		"nutmeg": "1/8 tablespoon",
		"lemon juice": "1 tablespoon"
	}`

	recipeMap := make(map[string]string)

	err := json.Unmarshal([]byte(data), &recipeMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonData, err := json.MarshalIndent(recipeMap, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("recipe.json", jsonData, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Recipe saved to recipe.json")
}
