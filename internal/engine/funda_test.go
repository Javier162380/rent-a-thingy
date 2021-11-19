package engine

import (
	"reflect"
	"rent-a-thingy/internal/models"
	"testing"
)

func Test_translateFundaSortCategory(t *testing.T) {
	type args struct {
		sortTerm string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "date asc",
		args: args{sortTerm: "date ↑"},
		want: "sorteer-datum-af",
	}, {
		name: "date desc",
		args: args{sortTerm: "date ↓"},
		want: "sorteer-datum-op",
	}, {
		name: "relevance",
		args: args{sortTerm: "relevance"},
		want: "",
	}, {
		name: "rental price",
		args: args{sortTerm: "rental price"},
		want: "sorteer-huurprijs-op",
	}, {
		name: "floor area",
		args: args{sortTerm: "floor area"},
		want: "sorteer-woonopp-af",
	}, {
		name: "availability",
		args: args{sortTerm: "availability"},
		want: "sorteer-beschikbaarheid-op",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateFundaSortCategory(tt.args.sortTerm); got != tt.want {
				t.Errorf("translateFundaSortCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateFundaDistance(t *testing.T) {
	type args struct {
		distance string
	}
	tests := []struct {
		name string
		args args
		want string
	}{{
		name: "distance",
		args: args{distance: "distance"},
		want: "distance",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateFundaDistance(tt.args.distance); got != tt.want {
				t.Errorf("translateFundaDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_translateFundaPrices(t *testing.T) {
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
			if got := translateFundaPrices(tt.args.maxPrice, tt.args.minPrice); got != tt.want {
				t.Errorf("translateFundaPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_funda_BuildUrl(t *testing.T) {
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
		fields: fields{baseUrl: FUNDA_BASE_URL},
		args:   args{metadata: models.RequestMetadata{City: "Amsterdam", Engine: "funda", ZipCodeOrDistricts: "1054", MaxPrice: "100", MinPrice: "50", CustomParams: map[string]string{"distance": "+5km"}}},
		want:   "https://www.funda.nl/en/huur//1054/50-100/+5km",
	}, {
		name:   "test 2",
		fields: fields{baseUrl: FUNDA_BASE_URL},
		args:   args{metadata: models.RequestMetadata{City: "Rotterdam", Engine: "funda", ZipCodeOrDistricts: "1154", SortCategory: "rental price", MaxPrice: "", MinPrice: "", CustomParams: map[string]string{"distance": "+5km"}}},
		want:   "https://www.funda.nl/en/huur//1154/+5km/sorteer-huurprijs-op",
	}, {
		name:   "test 3",
		fields: fields{baseUrl: FUNDA_BASE_URL},
		args:   args{metadata: models.RequestMetadata{City: "Rotterdam", Engine: "funda", ZipCodeOrDistricts: "1154", SortCategory: "date ↓", MaxPrice: "1000", MinPrice: "10", CustomParams: map[string]string{"distance": "+5km"}}},
		want:   "https://www.funda.nl/en/huur//1154/10-1000/+5km/sorteer-datum-op",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &funda{
				baseUrl: tt.fields.baseUrl,
			}
			if got := f.BuildUrl(tt.args.metadata); got != tt.want {
				t.Errorf("funda.BuildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFundaEngine(t *testing.T) {
	tests := []struct {
		name string
		want EngineBuilder
	}{{
		name: "funda",
		want: &funda{baseUrl: FUNDA_BASE_URL},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFundaEngine(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFundaEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}
