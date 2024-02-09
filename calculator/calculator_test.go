package calculator

import (
	"testing"
)

func TestMyCalculator_Add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		m    TheBestCalculator
		args args
		want int
	}{
		{
			name: "Test Add 1, 1 expect 2",
			m:    NewTheBestCalculator("My Calculator"),
			args: args{1, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("MyCalculator.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCalculator_GetCalculatorName(t *testing.T) {
	tests := []struct {
		name string
		m    TheBestCalculator
		want string
	}{
		{
			name: "Test GetCalculatorName",
			m:    NewTheBestCalculator("ABC Cals"),
			want: "ABC Cals",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.GetCalculatorName(); got != tt.want {
				t.Errorf("MyCalculator.GetCalculatorName() = %v, want %v", got, tt.want)
			}
		})
	}
}
