package engine

import "rent-a-thingy/internal/models"

type EngineBuilder interface {
	BuildUrl(models.RequestMetadata) string
}

func NewEngine(engineType string) EngineBuilder {
	switch engineType {

	case "funda":
		return NewFundaEngine()
	case "pararius":
		return NewParariusEngine()
	case "idealista":
		return NewIdealistaEngine()
	default:
		return nil
	}
}
