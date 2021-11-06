package form

import "rent-a-thingy/internal/models"

type EngineForm interface {
	MutateRequestMetadata(models.RequestMetadata) (models.RequestMetadata, error)
}

type fundaForm struct {
}

type ParariusForm struct {
}

type IdealistaForm struct {
}

func (f *fundaForm) MutateRequestMetadata(metadata models.RequestMetadata) (models.RequestMetadata, error) {

	params := make(map[string]string)

	distanceRadius := []string{"+0km", "+1km", "+2km", "+5km", "+10km", "+15km", "+30km", "+50km", "+100km"}
	distance, err := argSelectorWithOptions("Select Radius From the Zip/Code", distanceRadius, "Not allowed value", true)

	if err != nil {
		return metadata, err
	}

	params["distance"] = distance
	metadata.CustomParams = params

	return metadata, nil
}

func (p *ParariusForm) MutateRequestMetadata(metadata models.RequestMetadata) (models.RequestMetadata, error) {
	return metadata, nil
}

func (i *IdealistaForm) MutateRequestMetadata(metadata models.RequestMetadata) (models.RequestMetadata, error) {
	return metadata, nil
}

func NewEngineForm(engine string) EngineForm {
	switch engine {
	case "funda":
		return &fundaForm{}
	case "pararius":
		return &ParariusForm{}
	case "idealista":
		return &IdealistaForm{}
	default:
		return nil
	}
}
