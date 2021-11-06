package engine

import (
	"fmt"
	"rent-a-thingy/internal/models"
	"strconv"
)

type funda struct {
	baseUrl string
}

func translateFundaSortCategory(sortTerm string) string {
	switch sortTerm {
	case "date ↑":
		return "sorteer-datum-af"

	case "date ↓":
		return "sorteer-datum-op"

	case "relevance":
		return ""

	case "rental price":
		return "sorteer-huurprijs-op"

	case "floor area":
		return "sorteer-woonopp-af"

	case "availability":
		return "sorteer-beschikbaarheid-op"

	default:
		return ""
	}

}

func translateFundaDistance(distance string) string {

	return distance
}

func translateFundaPrices(maxPrice string, minPrice string) string {
	maxPriceInteger, err := strconv.Atoi(maxPrice)

	if err != nil {
		return ""
	}

	minPriceInteger, err := strconv.Atoi(minPrice)
	if err != nil {
		return ""
	}

	switch {
	case maxPriceInteger < minPriceInteger:
		return fmt.Sprintf("%d-%d", maxPriceInteger, maxPriceInteger)

	case minPriceInteger < maxPriceInteger:
		return fmt.Sprintf("%d-%d", minPriceInteger, maxPriceInteger)

	default:
		return ""
	}

}

func (f *funda) BuildUrl(metadata models.RequestMetadata) string {
	locationUrl := f.baseUrl + "/" + metadata.ZipCodeOrDistricts

	if priceInterval := translateFundaPrices(metadata.MaxPrice, metadata.MinPrice); priceInterval != "" {
		locationUrl += "/" + priceInterval
	}

	if distance := translateFundaDistance(metadata.CustomParams["distance"]); distance != "" {
		locationUrl += "/" + distance
	}

	if sortOptions := translateFundaSortCategory(metadata.SortCategory); sortOptions != "" {
		locationUrl += "/" + sortOptions
	}

	fmt.Printf("%s\n", locationUrl)

	return locationUrl
}

func NewFundaEngine() EngineBuilder {
	return &funda{
		baseUrl: "https://www.funda.nl/en/huur/",
	}
}
