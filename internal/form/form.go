package form

import (
	"fmt"

	"rent-a-thingy/internal/engine"
	"rent-a-thingy/internal/models"
)

func ExecuteForm() {

	cityName, err := argSelectorWithOptions("Select City", models.Cities, "City not available", true)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}

	cityMetadata, err := models.GetCityMetadata(cityName)

	if err != nil {
		fmt.Printf("%s", err)
	}

	engineName, err := argSelectorWithOptions("Select engine", cityMetadata.Engines, "Site not found", true)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	engineMetadata, err := models.GetEngineMetadata(engineName, cityName)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	zipCodeOrDistricts, err := argSelectorWithOptions("Select ZipCode or Districts", engineMetadata.ZipCodesOrDistricts, "ZipCode or Districts not found", true)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	sortCategory, err := argSelectorWithOptions("Sort Information by", engineMetadata.SortBy, "Not allowed value", true)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	maxPrice, err := argSelector("Select max price", "Not allowed value")

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	minPrice, err := argSelector("Select min price", "Not allowed value")

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	genericMetadata := models.RequestMetadata{
		City:               cityName,
		Engine:             engineName,
		ZipCodeOrDistricts: zipCodeOrDistricts,
		SortCategory:       sortCategory,
		MaxPrice:           maxPrice,
		MinPrice:           minPrice,
	}

	formInstance := NewEngineForm(engineName)
	enrichMetadata, err := formInstance.MutateRequestMetadata(genericMetadata)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	engineInstance := engine.NewEngine(engineName)

	url := engineInstance.BuildUrl(enrichMetadata)

	openBrowser(url)
}
