package engine

import (
	"fmt"
	"rent-a-thingy/internal/models"
	"strings"
)

var IDEALISTA_BASE_URL = "https://www.idealista.com/alquiler-viviendas/"

type idealista struct {
	baseURL string
}

func translateIdealistaSortCategory(sortTerm string) string {

	idealistaSortPrefix := "?ordenado-por="
	switch sortTerm {
	case "Relevancia":
		return ""
	case "Baratos":
		return fmt.Sprintf("%s%s", idealistaSortPrefix, "precios-asc")
	case "Caros":
		return fmt.Sprintf("%s%s", idealistaSortPrefix, "precios-desc")
	case "Recientes":
		return fmt.Sprintf("%s%s", idealistaSortPrefix, "fecha-desc")
	case "Mansiones":
		return fmt.Sprintf("%s%s", idealistaSortPrefix, "area-desc")
	case "Zulos":
		return fmt.Sprintf("%s%s", idealistaSortPrefix, "area-asc")
	default:
		return ""
	}
}

func translateIdealistaDistrict(districtName string) string {
	singleDistrictName := strings.ToLower(districtName)
	return strings.ReplaceAll(singleDistrictName, " ", "-")
}

func (i *idealista) BuildUrl(metadata models.RequestMetadata) string {
	locationUrl := i.baseURL + strings.ToLower(metadata.City)

	if formatDistrict := translateIdealistaDistrict(metadata.ZipCodeOrDistricts); formatDistrict != "" {
		locationUrl += "/" + formatDistrict + "/"
	}

	if formatSortCategory := translateIdealistaSortCategory(metadata.SortCategory); formatSortCategory != "" {
		locationUrl += formatSortCategory
	}

	return locationUrl
}

func NewIdealistaEngine() EngineBuilder {
	return &idealista{
		baseURL: IDEALISTA_BASE_URL,
	}
}
