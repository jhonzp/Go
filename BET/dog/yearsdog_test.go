// Package dog allows us to more fully understand dogs.

package dog

import (
	"fmt"
	"testing"
)

func TestYearsTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "jhon", args: args{n: 4}, want: 28},
		{name: "juan", args: args{n: 3}, want: 21},
		{name: "Carl", args: args{n: 2}, want: 14},
		{name: "Otro", args: args{n: 1}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YearsTwo(tt.args.n); got != tt.want {
				t.Errorf("YearsTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(2))
	//Output: 14
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(30)
	}

}

func TestYears(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{name: "jhon", args: 4, want: 28},
		{name: "tres", args: 3, want: 21},
		{name: "dos", args: 2, want: 14},
		{name: "Uno", args: 1, want: 7},
	}

	for _, test := range tests {
		v := Years(test.args)
		if v != test.want {
			t.Errorf("Expected %v, got %v, test %v", test.want, v, test.name)
		}
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(30)
	}

}

func ExampleYears() {
	fmt.Println(Years(2))
	//Output: 14
}
