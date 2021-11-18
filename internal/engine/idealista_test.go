package engine

import (
	"reflect"
	"rent-a-thingy/internal/models"
	"testing"
)

func Test_translateIdealistaSortCategory(t *testing.T) {
	type args struct {
		sortTerm string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "Random Sort",
		args: args{sortTerm: "Random"},
		want: "",
	}, {
		name: "Relevancia",
		args: args{sortTerm: "Relevancia"},
		want: "",
	}, {
		name: "Baratos",
		args: args{sortTerm: "Baratos"},
		want: "?ordenado-por=precios-asc",
	}, {
		name: "Caros",
		args: args{sortTerm: "Caros"},
		want: "?ordenado-por=precios-desc",
	}, {
		name: "Recientes",
		args: args{sortTerm: "Recientes"},
		want: "?ordenado-por=fecha-desc",
	}, {
		name: "Mansiones",
		args: args{sortTerm: "Mansiones"},
		want: "?ordenado-por=area-desc",
	}, {
		name: "Zulos",
		args: args{sortTerm: "Zulos"},
		want: "?ordenado-por=area-asc",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateIdealistaSortCategory(tt.args.sortTerm); got != tt.want {
				t.Errorf("translateIdealistaSortCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateIdealistaDistrict(t *testing.T) {
	type args struct {
		districtName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "Almagro",
		args: args{districtName: "Almagro"},
		want: "almagro",
	}, {
		name: "Barrio de Salamanca",
		args: args{districtName: "Barrio de Salamanca"},
		want: "barrio-de-salamanca",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateIdealistaDistrict(tt.args.districtName); got != tt.want {
				t.Errorf("translateIdealistaDistrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idealista_BuildUrl(t *testing.T) {
	type fields struct {
		baseURL string
	}
	type args struct {
		metadata models.RequestMetadata
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{{
		name:   "BuildUrl1",
		fields: fields{baseURL: IDEALISTA_BASE_URL},
		args:   args{metadata: models.RequestMetadata{City: "Madrid", Engine: "idealista", ZipCodeOrDistricts: "Barrio de Salamanca", SortCategory: "Baratos"}},
		want:   "https://www.idealista.com/alquiler-viviendas/madrid/barrio-de-salamanca/?ordenado-por=precios-asc",
	}, {
		name:   "BuildUrl2",
		fields: fields{baseURL: IDEALISTA_BASE_URL},
		args:   args{metadata: models.RequestMetadata{City: "Madrid", Engine: "idealista", ZipCodeOrDistricts: "Almagro", SortCategory: "Zulos"}},
		want:   "https://www.idealista.com/alquiler-viviendas/madrid/almagro/?ordenado-por=area-asc",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &idealista{
				baseURL: tt.fields.baseURL,
			}
			if got := i.BuildUrl(tt.args.metadata); got != tt.want {
				t.Errorf("idealista.BuildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIdealistaEngine(t *testing.T) {
	tests := []struct {
		name string
		want EngineBuilder
	}{{
		name: "NewIdealistaEngine",
		want: &idealista{baseURL: IDEALISTA_BASE_URL},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdealistaEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdealistaEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}
