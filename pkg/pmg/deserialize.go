package pmg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Georges struct to hold all different Georges
type Georges struct {
	Georges []George `json:"georges"`
}

// Blang struct to hold the bling for the Georges to wear
type Blang struct {
	Blang []Bling `json:"blang"`
}

// George struct for George and all his sick accessories
type George struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Accessories []Bling
}

//Bling struct for bling that George can wear
type Bling struct {
	Noun     string   `json:"noun"`
	Location Location `json:"location"`
	Variants []BlingVariants `json:"variants"`
}

type BlingVariants struct {
	Adj string `json:"adj"`
	Image string `json:"image"`
}


// Location for the bling
type Location string

// Location enum for where the bling should be placed on George
const (
	HEAD Location = "HEAD"
	FACE Location = "FACE"
	NECK Location = "NECK"
	TORSO Location = "TORSO"
	LEGS Location = "LEGS"
	HANDS Location = "HANDS"
	FEET Location = "FEET"
)

// FetchGeorgeBlingData Parses JSON files and returns slices of the George structs and Bling structs
func FetchGeorgeBlingData(georgesPath string, blangPath string) ([]George, []Bling) {
	georgesJSON, georgesErr := os.Open(georgesPath)
	blangJSON, blangErr := os.Open(blangPath)

	if georgesErr != nil {
		fmt.Println(georgesErr)
	}

	if blangErr != nil {
		fmt.Println(blangErr)
	}

	defer georgesJSON.Close()
	defer blangJSON.Close()

	georgesBytes, _ := ioutil.ReadAll(georgesJSON)
	blangBytes, _ := ioutil.ReadAll(blangJSON)

	var georges Georges
	var blang Blang

	json.Unmarshal(georgesBytes, &georges)
	json.Unmarshal(blangBytes, &blang)

	return georges.Georges, blang.Blang
}
