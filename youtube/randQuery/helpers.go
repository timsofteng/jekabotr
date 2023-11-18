package randQuery

import "math/rand"

var enLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var uaLetterRunes = []rune("абвгґдеєжзиіїйклмнопрстуфхцчшщьюяАБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯ")

var orderValues = [6]string{"date", "rating", "relevance", "title", "videoCount", "viewCount"}

func randEnStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = enLetterRunes[rand.Intn(len(enLetterRunes))]
	}
	return string(b)
}

func randUaStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = uaLetterRunes[rand.Intn(len(uaLetterRunes))]
	}
	return string(b)
}

func String(n int) string {
	var fns []func(n int) string
	fns = append(fns, randEnStringRunes, randUaStringRunes)

	return fns[rand.Intn(len(fns))](n)

}


func Order() string {
	randomIndex := rand.Intn(len(orderValues))
	return orderValues[randomIndex]
}
