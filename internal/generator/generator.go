package generator

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"
	"strconv"
	"strings"
)

const (
	DefaultSymbols     = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"
	DefaultNumberSet   = "0123456789"
	DefaultEngAlphabet = "abcdefghijklmnopqrstuvwxyz"
	DefaultRusAlphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
)

var (
	DefaultConfig = Config{
		Length:           8,
		WithSymbols:      false,
		WithNumbers:      true,
		WithUpperLetters: false,
		WithEngLang:      true,
		WithRusLang:      false,
	}
	ErrEmptyConfig = errors.New("empty config")
	//ErrZeroLength  = errors.New("length of password is zero")
	ErrShortLength = errors.New("length of password to short")
)

type Config struct {
	Length           uint
	WithSymbols      bool
	WithNumbers      bool
	WithUpperLetters bool
	WithEngLang      bool
	WithRusLang      bool
	Alphabet         string
}

func (config *Config) String() string {
	return "length : " + strconv.Itoa(int(config.Length)) + "\n" +
		"with symbols : " + strconv.FormatBool(config.WithSymbols) + "\n" +
		"with numbers : " + strconv.FormatBool(config.WithNumbers) + "\n" +
		"with UpperLetters : " + strconv.FormatBool(config.WithUpperLetters) + "\n" +
		"with English alphabet : " + strconv.FormatBool(config.WithEngLang) + "\n" +
		"with Russian alphabet : " + strconv.FormatBool(config.WithRusLang) + "\n"
}

type Generator struct {
	*Config
}

func New(config *Config) (*Generator, error) {
	if config == nil {
		log.Printf("Empty config! Using a default config : %s", DefaultConfig.String())
		config = &DefaultConfig
	}
	if !config.WithSymbols && !config.WithNumbers && !config.WithUpperLetters &&
		!config.WithEngLang && !config.WithRusLang {
		log.Printf("empty params for config : %s", config)
		config = &DefaultConfig
	}
	if config.Length < 8 {
		log.Printf("password length is too short (%x), use a default lenght (%x)", config.Length, DefaultConfig.Length)
		config.Length = DefaultConfig.Length
		//return nil, ErrShortLength
	}
	config.Alphabet = CreateAlphabet(config)
	if int(config.Length) >= len(config.Alphabet) {
		log.Printf("password length is longer than alphabet (%x), use a 1/8 of alphabet length (%x)", config.Length, len(config.Alphabet)/8)
		config.Length = uint(len(config.Alphabet) / 8)
	}
	return &Generator{Config: config}, nil
}

func CreateAlphabet(config *Config) string {
	var alphabet string

	if config.WithRusLang == true {
		alphabet += DefaultRusAlphabet
	}
	if config.WithEngLang == true {
		alphabet += DefaultEngAlphabet
	}
	if config.WithUpperLetters == true && config.WithRusLang == true {
		alphabet += strings.ToUpper(DefaultRusAlphabet)
	}
	if config.WithUpperLetters == true && config.WithEngLang == true {
		alphabet += strings.ToUpper(DefaultEngAlphabet)
	}

	if config.WithNumbers == true {
		alphabet += DefaultNumberSet
	}

	if config.WithSymbols == true {
		alphabet += DefaultSymbols
	}

	return alphabet
}

func (g Generator) Generate() (*string, error) {
	var generated string
	characterSet := strings.Split(g.Config.Alphabet, "")
	maxInd := big.NewInt(int64(len(characterSet)))

	for i := uint(0); i < g.Config.Length; i++ {
		val, err := rand.Int(rand.Reader, maxInd)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}
