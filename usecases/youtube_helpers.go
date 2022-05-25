package usecases

import "math/rand"

var enLetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var uaLetterRunes = []rune("абвгґдеєжзиіїйклмнопрстуфхцчшщьюяАБВГҐДЕЄЖЗИІЇЙКЛМНОПРСТУФХЦЧШЩЬЮЯ")

func RandEnStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = enLetterRunes[rand.Intn(len(enLetterRunes))]
	}
	return string(b)
}

func RandUaStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = uaLetterRunes[rand.Intn(len(uaLetterRunes))]
	}
	return string(b)
}
