package form

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"

	"rent-a-thingy/internal/engine"
	"rent-a-thingy/internal/models"

	"github.com/manifoldco/promptui"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func argSelector(msg string, errorMessage string) (string, error) {

	promptValue := promptui.Prompt{
		Label: msg,
	}

	result, err := promptValue.Run()

	if err != nil {
		return "", errors.New(errorMessage)
	}

	return result, nil

}

func argSelectorWithOptions(msg string, options []string, errorMessage string, useoptions bool) (string, error) {

	if useoptions {
		promptList := promptui.Select{
			Label: msg,
			Items: options,
		}

		_, result, err := promptList.Run()

		if err != nil {
			return "", errors.New(errorMessage)
		}

		return result, nil

	} else {
		promptValue := promptui.Prompt{
			Label: msg,
		}

		result, err := promptValue.Run()

		if err != nil {
			return "", errors.New(errorMessage)
		}

		if contains(options, result) {
			return result, nil
		} else {
			return "", errors.New(fmt.Sprintf("Option %s : %s it is not available", msg, result))
		}

	}

}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("sensible-browser", url)
	default:
		return fmt.Errorf("Failed to open browser. Platform %s not supported.", runtime.GOOS)
	}

	return cmd.Start()
}

func ExecuteForm() {

	cityName, err := argSelectorWithOptions("Select City", models.Cities, "City not available", true)

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}

	cityMetadata, err := models.GetCityMetadata(cityName)

	if err != nil {
		fmt.Printf("%s", err)
	}

	engineName, err := argSelectorWithOptions("Select site", cityMetadata.Engines, "Site not found", true)

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

	distance, err := argSelector("Select distance", "Not allowed value")

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

	requestMetadata := models.RequestMetadata{
		City:               cityName,
		Engine:             engineName,
		ZipCodeOrDistricts: zipCodeOrDistricts,
		SortCategory:       sortCategory,
		Distance:           distance,
		MaxPrice:           maxPrice,
		MinPrice:           minPrice,
	}

	engineInstance := engine.NewEngine(engineName)

	url := engineInstance.BuildUrl(requestMetadata)

	openBrowser(url)
}
