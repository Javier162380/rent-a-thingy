package engine

import (
	"reflect"
	"testing"
)

func TestNewEngine(t *testing.T) {
	type args struct {
		engineType string
	}
	tests := []struct {
		name string
		args args
		want EngineBuilder
	}{{
		name: "NewFundaEngine",
		args: args{engineType: "funda"},
		want: &funda{baseUrl: FUNDA_BASE_URL},
	}, {
		name: "NewParariusEngine",
		args: args{engineType: "pararius"},
		want: &pararius{baseUrl: PARARIUS_BASE_URL},
	}, {
		name: "NewIdealistaEngine",
		args: args{engineType: "idealista"},
		want: &idealista{baseURL: IDEALISTA_BASE_URL},
	}, {
		name: "NewRandomEngine",
		args: args{engineType: "random"},
		want: nil,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEngine(tt.args.engineType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}
