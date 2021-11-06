package models

import "errors"

var Cities = []string{"Amsterdam", "Rotterdam", "Utrech", "Madrid"}
var fundaSortCategories = []string{"date ↓", "date ↑", "relevance", "rental price", "floor area", "availability"}
var parariusSortCategories = []string{"date ↓", "rental price ↑", "rental price ↓", "living area ↑", "living area ↓"}
var idealistaSortCategories = []string{"Relevancia", "Baratos", "Caros", "Recientes", "Mansiones", "Zulos"}

type City struct {
	Engines []string
}

type Engine struct {
	ZipCodesOrDistricts []string
	SortBy              []string
}

type RequestMetadata struct {
	City               string
	Engine             string
	SortCategory       string
	ZipCodeOrDistricts string
	Distance           string
	MaxPrice           string
	MinPrice           string
}

func cityMetadata() map[string]City {
	metaData := make(map[string]City)
	metaData["Amsterdam"] = City{Engines: []string{"pararius", "funda"}}
	metaData["Rotterdam"] = City{Engines: []string{"pararius", "funda"}}
	metaData["Utrech"] = City{Engines: []string{"pararius", "funda"}}
	metaData["Madrid"] = City{Engines: []string{"idealista"}}

	return metaData
}

func engineMetadata(city string) (map[string]Engine, error) {
	metaData := make(map[string]Engine)

	switch city {
	case "Amsterdam":
		metaData["funda"] = Engine{
			ZipCodesOrDistricts: []string{"1011", "1012", "1013", "1014", "1015", "1016", "1017", "1018", "1019", "1021", "1022", "1023", "1024", "1025", "1026", "1027", "1028", "1031", "1032", "1033", "1034", "1035", "1036", "1037", "1041", "1042", "1043", "1044", "1045", "1046", "1047", "1051", "1052", "1053", "1054", "1055", "1056", "1057", "1058", "1059", "1060", "1061", "1062", "1063", "1064", "1065", "1066", "1067", "1068", "1069", "1071", "1072", "1073", "1074", "1075", "1076", "1077", "1078", "1079", "1081", "1082", "1083", "1086", "1087", "1091", "1092", "1093", "1094", "1095", "1096", "1097", "1098", "1101", "1102", "1103", "1104", "1105", "1106", "1107", "1108", "1109"},
			SortBy:              fundaSortCategories}
		metaData["pararius"] = Engine{
			ZipCodesOrDistricts: []string{"amstel 3", "bijlmer centrum", "bijlmer oost", "bos en lommer", "buitenveldert", "centrum oost", "centrum west", "de aker", "de baarsjes", "de pijp", "driemond", "gaasperdam", "geuzenveld", "ijburg", "indische buurt", "noord oost", "noord west", "oostelijk havengebied", "osdorp", "oud noord", "oud oost", "oud west", "oud zuid", "rivierenbuurt", "sloten nieuw sloten", "sloterdijken", "slotermeer", "slotervaart", "watergraafsmeer", "westelijk havengebied", "westerpark", "zeeburgereiland", "zuidas"},
			SortBy:              parariusSortCategories}
	case "Rotterdam":
		metaData["funda"] = Engine{
			ZipCodesOrDistricts: []string{"3011", "3012", "3013", "3014", "3015", "3016", "3021", "3022", "3023", "3024", "3025", "3026", "3027", "3028", "3029", "3031", "3032", "3033", "3034", "3035", "3036", "3037", "3038", "3039", "3041", "3042", "3043", "3044", "3045", "3046", "3047", "3051", "3052", "3053", "3054", "3055", "3056", "3059", "3061", "3062", "3063", "3064", "3065", "3066", "3067", "3068", "3069", "3071", "3072", "3073", "3074", "3075", "3076", "3077", "3078", "3079", "3081", "3082", "3083", "3084", "3085", "3086", "3087", "3088", "3089", "3921"},
			SortBy:              fundaSortCategories}
		metaData["pararius"] = Engine{
			ZipCodesOrDistricts: []string{"bedrijventerrein-schieveen", "charlois", "delfshaven", "feijenoord", "hillegersberg-schiebroek", "ijsselmonde", "kralingen-crooswijk", "nieuw-mathenesse", "noord", "overschie", "prins-alexander", "rivium", "rotterdam-centrum", "rotterdam-noord-west", "spaanse-polder", "waalhaven-eemhaven"},
			SortBy:              parariusSortCategories}
	case "Utrech":
		metaData["funda"] = Engine{
			ZipCodesOrDistricts: []string{"3454", "3511", "3512", "3513", "3514", "3515", "3521", "3522", "3523", "3524", "3525", "3526", "3527", "3528", "3531", "3532", "3533", "3534", "3541", "3542", "3543", "3544", "3545", "3546", "3551", "3552", "3553", "3554", "3555", "3561", "3562", "3563", "3564", "3565", "3566", "3571", "3572", "3573", "3581", "3582", "3583", "3584", "3585"},
			SortBy:              fundaSortCategories}
		metaData["pararius"] = Engine{
			ZipCodesOrDistricts: []string{"abstede gansstraat", "binnenstad city en winkelgebied", "binnenstad woongebied", "dichterswijk rivierenwijk", "haarrijn", "het zand", "kanaleneiland", "leijdsche rijn centrum e o", "leijdsche rijn zuid", "lombok leidseweg", "lunetten", "nieuw engeland schepenbuurt", "nieuw hoograven bokkenbuurt", "oog in al welgelegen", "oud hoograven tolsteeg", "oudwijk buiten wittevrouwen", "parkwijk langerak", "pijlsweerd", "taagdreef wolgadreef", "terwijde de wetering", "transwijk", "tuindorp voordorp", "vechtzoom klopvaart", "votulast", "wilhelminapark rijnsweerd", "wittevrouwen zeeheldenbuurt", "zambesidreef tigrisdreef", "zamenhofdreef neckardreef", "zuilen noord en oost", "zuilen west"},
			SortBy:              parariusSortCategories}
	case "Madrid":
		metaData["idealista"] = Engine{
			ZipCodesOrDistricts: []string{"Centro", "Carabanchel", "Barrio de Salamanca", "Barajas", "Arganzuela", "Chamartín", "Chamberí", "Ciudad Lineal", "Fuencarral", "Hortaleza", "Latina", "Moncloa", "Moratalaz", "Puente de Vallecas", "Retiro", "San Blas", "Tetuan", "Usera", "Vicalvaro", "Villa de Vallecas", "Villaverde"},
			SortBy:              idealistaSortCategories}

	default:
		return metaData, errors.New("City not found")
	}

	return metaData, nil
}

func GetCityMetadata(name string) (City, error) {

	metaData := cityMetadata()

	if value, ok := metaData[name]; ok {
		return value, nil
	} else {
		return City{}, errors.New("City not found")
	}
}

func GetEngineMetadata(name string, city string) (Engine, error) {
	metaData, err := engineMetadata(city)

	if err != nil {
		return Engine{}, err
	}

	if value, ok := metaData[name]; ok {
		return value, nil
	} else {
		return Engine{}, errors.New("City not found")
	}
}
