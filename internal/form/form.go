package form

import (
	"fmt"

	"rent-a-thingy/internal/engine"
	"rent-a-thingy/internal/models"
)

func genericForm() (models.RequestMetadata, error) {

	cityName, err := argSelectorWithOptions("Select City", models.Cities, "City not available", true)

	if err != nil {
		return models.RequestMetadata{}, err
	}

	cityMetadata, err := models.GetCityMetadata(cityName)

	if err != nil {
		return models.RequestMetadata{}, err
	}

	engineName, err := argSelectorWithOptions("Select engine", cityMetadata.Engines, "Site not found", true)

	if err != nil {
		return models.RequestMetadata{}, err
	}

	engineMetadata, err := models.GetEngineMetadata(engineName, cityName)

	if err != nil {
		return models.RequestMetadata{}, err
	}

	zipCodeOrDistricts, err := argSelectorWithOptions("Select ZipCode or Districts", engineMetadata.ZipCodesOrDistricts, "ZipCode or Districts not found", true)

	if err != nil {
		return models.RequestMetadata{}, err
	}

	sortCategory, err := argSelectorWithOptions("Sort Information by", engineMetadata.SortBy, "Not allowed value", true)

	if err != nil {
		return models.RequestMetadata{}, err
	}

	maxPrice, err := argSelector("Select max price", "Not allowed value")

	if err != nil {
		return models.RequestMetadata{}, err
	}

	minPrice, err := argSelector("Select min price", "Not allowed value")

	if err != nil {
		return models.RequestMetadata{}, err
	}

	genericMetadata := models.RequestMetadata{
		City:               cityName,
		Engine:             engineName,
		ZipCodeOrDistricts: zipCodeOrDistricts,
		SortCategory:       sortCategory,
		MaxPrice:           maxPrice,
		MinPrice:           minPrice,
	}

	return genericMetadata, nil

}

func ExecuteForm() {

	requestMetadata, err := genericForm()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	formInstance := NewEngineForm(requestMetadata.Engine)
	enrichMetadata, err := formInstance.MutateRequestMetadata(requestMetadata)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	engineInstance := engine.NewEngine(requestMetadata.Engine)

	url := engineInstance.BuildUrl(enrichMetadata)

	openBrowser(url)
}
