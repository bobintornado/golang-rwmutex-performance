package performance

import "testing"

func BenchmarkRCache(b *testing.B) {
	r := RCache{}
	r.Set("value")
	for n := 0; n < b.N; n++ {
		go r.Set("BBB")
		go r.Get("value")
	}
}

func BenchmarkLCache(b *testing.B) {
	r := LCache{}
	r.Set("value")
	for n := 0; n < b.N; n++ {
		go r.Set("BBB")
		go r.Get("value")
	}
}
