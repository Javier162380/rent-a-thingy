package engine

import (
	"fmt"
	"rent-a-thingy/internal/models"
	"strconv"
	"strings"
)

type pararius struct {
	baseUrl string
}

func translateParariusSortCategory(sortTerm string) string {
	switch sortTerm {
	case "date ↓":
		return ""
	case "rental price ↑":
		return "sorteer-prijs-op"
	case "rental price ↓":
		return "sorteer-prijs-af"
	case "living area ↓":
		return "sorteer-woonopp-af"
	case "living area ↑":
		return "sorteer-woonopp-op"

	default:
		return ""
	}

}

func translateParariusDistrict(districtName string) string {
	formattedDistrictName := strings.ReplaceAll(districtName, " ", "-")

	return fmt.Sprintf("%s-%s", "wijk", formattedDistrictName)
}

func translateParariusPrices(maxPrice string, minPrice string) string {
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

func (p *pararius) BuildUrl(metadata models.RequestMetadata) string {
	locationUrl := p.baseUrl + strings.ToLower(metadata.City)

	if priceInterval := translateParariusPrices(metadata.MaxPrice, metadata.MinPrice); priceInterval != "" {
		locationUrl = locationUrl + "/" + priceInterval
	}

	if sortCategory := translateParariusSortCategory(metadata.SortCategory); sortCategory != "" {
		locationUrl += "/" + sortCategory
	}

	if formatDistrict := translateParariusDistrict(metadata.ZipCodeOrDistricts); formatDistrict != "" {
		locationUrl += "/" + formatDistrict
	}

	return locationUrl
}

func NewParariusEngine() EngineBuilder {
	return &pararius{
		baseUrl: "https://www.pararius.nl/huurwoningen/",
	}
}
