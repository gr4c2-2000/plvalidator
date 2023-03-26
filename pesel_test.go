package plvalidator

import "testing"

func TestPesel(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty string", args{""}, false},
		{"proper", args{"66111146882"}, true},
		{"inccorect sum", args{"66111146883"}, false},
		{"alfabet", args{"cdcdcdcdcdc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pesel(tt.args.id); got != tt.want {
				t.Errorf("Pesel() = %v, want %v", got, tt.want)
			}
		})
	}
}
