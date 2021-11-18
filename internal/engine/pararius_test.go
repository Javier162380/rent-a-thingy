package engine

import (
	// "reflect"
	// "rent-a-thingy/internal/models"
	"reflect"
	"rent-a-thingy/internal/models"
	"testing"
)

func Test_translateParariusSortCategory(t *testing.T) {
	type args struct {
		sortTerm string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "date",
		args: args{sortTerm: "date ↓"},
		want: "",
	}, {
		name: "empty",
		args: args{sortTerm: ""},
		want: "",
	}, {
		name: "rental price up",
		args: args{sortTerm: "rental price ↑"},
		want: "sorteer-prijs-op",
	}, {
		name: "rental price down",
		args: args{sortTerm: "rental price ↓"},
		want: "sorteer-prijs-af",
	}, {
		name: "living are down",
		args: args{sortTerm: "living area ↓"},
		want: "sorteer-woonopp-af",
	}, {
		name: "living area up",
		args: args{sortTerm: "living area ↑"},
		want: "sorteer-woonopp-op",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateParariusSortCategory(tt.args.sortTerm); got != tt.want {
				t.Errorf("translateParariusSortCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateParariusDistrict(t *testing.T) {
	type args struct {
		districtName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "District 1",
		args: args{districtName: "dostr 2"},
		want: "wijk-dostr-2",
	}, {
		name: "District 2",
		args: args{districtName: "dostr 3"},
		want: "wijk-dostr-3",
	}, {
		name: "District 3",
		args: args{districtName: "dostr4"},
		want: "wijk-dostr4",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateParariusDistrict(tt.args.districtName); got != tt.want {
				t.Errorf("translateParariusDistrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateParariusPrices(t *testing.T) {
	type args struct {
		maxPrice string
		minPrice string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{

		name: "interval 1",
		args: args{maxPrice: "100", minPrice: "50"},
		want: "50-100",
	}, {
		name: "interval 2",
		args: args{maxPrice: "100", minPrice: "100"},
		want: "100-100",
	}, {
		name: "interval 3",
		args: args{maxPrice: "100", minPrice: "200"},
		want: "100-200",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateParariusPrices(tt.args.maxPrice, tt.args.minPrice); got != tt.want {

				t.Errorf("translateParariusPrices(%v, %v) = %v, want %v", tt.args.maxPrice, tt.args.minPrice, got, tt.want)
			}
		})
	}
}

func Test_pararius_BuildUrl(t *testing.T) {
	type fields struct {
		baseUrl string
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
		name:   "test 1",
		fields: fields{baseUrl: PARARIUS_BASE_URL},
		args:   args{metadata: models.RequestMetadata{City: "Amsterdam", Engine: "parirus", ZipCodeOrDistricts: "wrok", SortCategory: "date ↓", MaxPrice: "2000", MinPrice: "4000"}},
		want:   "https://www.pararius.nl/huurwoningen/amsterdam/2000-4000/wijk-wrok",
	}, {
		name:   "test 2",
		fields: fields{baseUrl: PARARIUS_BASE_URL},
		args:   args{metadata: models.RequestMetadata{City: "Amsterdam", Engine: "parirus", ZipCodeOrDistricts: "wrok 3", SortCategory: "rental price ↑", MaxPrice: "4000", MinPrice: "1000"}},
		want:   "https://www.pararius.nl/huurwoningen/amsterdam/1000-4000/sorteer-prijs-op/wijk-wrok-3",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pararius{
				baseUrl: tt.fields.baseUrl,
			}
			if got := p.BuildUrl(tt.args.metadata); got != tt.want {
				t.Errorf("pararius.BuildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewParariusEngine(t *testing.T) {
	tests := []struct {
		name string
		want EngineBuilder
	}{{
		name: "pararius",
		want: &pararius{baseUrl: PARARIUS_BASE_URL},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParariusEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParariusEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}
