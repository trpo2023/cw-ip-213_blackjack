package random

import (
	"math/rand"
	"time"
)

type Characters = string

var (
	randomizer = New()
)

const (
	CharsRomanLettersLower Characters = "abcdefghijklmnopqrstuvwxyz"
	CharsRomanLettersUpper Characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsArabicNumerals    Characters = "0123456789"
	CharsRomanLetters                 = CharsRomanLettersLower + CharsRomanLettersUpper
	CharsDefault                      = CharsArabicNumerals + CharsRomanLetters
)

type Config struct {
	Chars Characters
	Seed  int64
}

type Random struct {
	core     *rand.Rand
	seed     int64
	chars    []rune
	charsLen int64
}

func New(conf ...Config) *Random {
	var (
		chars = CharsDefault
		seed  = time.Now().UnixNano()
	)

	if len(conf) > 0 {
		config := conf[0]

		if config.Chars != "" {
			chars = config.Chars
		}

		if config.Seed > 0 {
			seed = config.Seed
		}
	}

	return &Random{
		core:     rand.New(rand.NewSource(seed)),
		seed:     seed,
		chars:    []rune(chars),
		charsLen: int64(len(chars)),
	}
}

func (r *Random) Runes(size int) []rune {
	buf := make([]rune, size)

	for i := 0; i < size; i++ {
		buf[i] = r.chars[r.core.Int63()%r.charsLen]
	}

	return buf
}

func (r *Random) RandString(size int) string {
	return string(r.Runes(size))
}

func (r *Random) RandInt(min, max int) int {
	return min + (r.core.Int() % (max - min))
}

func (r *Random) Int() int {
	return r.core.Int()
}
func Runes(size int) []rune {
	return randomizer.Runes(size)
}

func RandString(size int) string {
	return randomizer.RandString(size)
}

func Int() int {
	return randomizer.Int()
}

func RandInt(min, max int) int {
	return randomizer.RandInt(min, max)
}
