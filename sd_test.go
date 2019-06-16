package main

import (
	"awesomeProject/waldo"
	"testing"
)

func TestSumaDivisibles(t *testing.T) {

	t.Run("sumaDivisible", func(t *testing.T) {
		got := waldo.SumaDivisibles(10)
		want := 33

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}


func BenchmarkArea(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.Run("sumaDivisibles2", func(b *testing.B) {
			got := waldo.SumaDivisibles(50)
			want := 593
			if got != want {
				b.Errorf("got %d want %d", got, want)
			}
		})

		b.RunParallel(
			func(pb *testing.PB) {
				for pb.Next() {
					waldo.SumaDivisibles(50)
				}
			},
		)
	}
}