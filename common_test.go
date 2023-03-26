package plvalidator

import (
	"reflect"
	"testing"
)

func TestSplitToInt(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{"proper", args{"01234"}, []int{0, 1, 2, 3, 4}, false},
		{"error", args{"aaaaaa"}, nil, true},
		{"zero", args{"0"}, []int{0}, false},
		{"mix", args{"120aa12"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitToInt(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
