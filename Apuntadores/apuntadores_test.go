package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_cambiarValor(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1",
			args: args{
				a: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cambiarValor(tt.args.a)
		})
	}
}
