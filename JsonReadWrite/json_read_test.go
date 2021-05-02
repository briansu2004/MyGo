package main

import "testing"

func BenchmarkMovieReadSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSONUnmarshal("movie.json", func(data map[string]interface{}) bool {
			return data["year"].(float64) >= 2010
		})
	}
}

func BenchmarkMovieReadToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSONDecoder("movie.json", func(data map[string]interface{}) bool {
			return data["year"].(float64) >= 2010
		})
	}
}

func BenchmarkQReadSmall(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSONUnmarshal("question.json", func(data map[string]interface{}) bool {
			return data["show_number"].(string) == "4680"
		})
	}
}

func BenchmarkQReadToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSONDecoder("question.json", func(data map[string]interface{}) bool {
			return data["show_number"].(string) == "4680"
		})
	}
}
