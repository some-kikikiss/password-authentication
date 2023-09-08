package generator

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var TestDefaultConfig = Config{
	Length:           8,
	WithSymbols:      false,
	WithNumbers:      true,
	WithUpperLetters: false,
	WithEngLang:      true,
	WithRusLang:      false,
	Alphabet:         "abcdefghijklmnopqrstuvwxyz0123456789",
}

func TestNew(t *testing.T) {
	var testConf = &Config{1, false, false, false, false, false, ""}

	t.Log("Given the need to test NewGenerator behavior at different params")
	{
		testID := 0
		t.Logf("\tTest %d:\t When Config is nil.", testID)
		{
			testGenerator, _ := New(testConf)

			assert.Equalf(
				t,
				TestDefaultConfig,
				*testGenerator.Config,
				"\tShould be same as DefaultCongih : \n %s \n \t error is %s",
				TestDefaultConfig,
				ErrEmptyConfig,
			)
		}
		testID++
		t.Logf("\tTest %d:\t When Length to short (<8).", testID)
		{
			testGenerator, _ := New(testConf)

			assert.Equalf(
				t,
				TestDefaultConfig.Length,
				testGenerator.Config.Length,
				"\tShould be length 8 (as a default) : \n %s \n \t error is %s",
				TestDefaultConfig.Length,
				ErrShortLength,
			)
		}
		testID++
		testConf.Length = 8
		t.Logf("\tTest %d:\t When all booleans is false .", testID)
		{
			testGenerator, err := New(testConf)
			assert.Equalf(
				t,
				TestDefaultConfig.WithSymbols,
				testGenerator.WithSymbols,
				"\tShould be withSymbols as DefaultConfig : \n %d \n \t error is %s",
				TestDefaultConfig.WithSymbols,
				err,
			)
			assert.Equalf(
				t,
				TestDefaultConfig.WithEngLang,
				testGenerator.WithEngLang,
				"\tShould be withEngLang as DefaultConfig : \n %s \n \t error is %s",
				TestDefaultConfig.WithEngLang,
				err,
			)
			assert.Equalf(
				t,
				TestDefaultConfig.WithNumbers,
				testGenerator.WithNumbers,
				"\tShould be numbers as DefaultConfig : \n %s \n \t error is %s",
				TestDefaultConfig.WithNumbers,
				err,
			)
			assert.Equalf(
				t,
				TestDefaultConfig.WithUpperLetters,
				testGenerator.WithUpperLetters,
				"\tShould be withUpperLetters as DefaultConfig : \n %s \n \t error is %s",
				TestDefaultConfig.WithUpperLetters,
				err,
			)
			assert.Equalf(
				t,
				TestDefaultConfig.WithRusLang,
				testGenerator.WithRusLang,
				"\tShould be wuthRuLangas DefaultConfig : \n %s \n \t error is %s",
				TestDefaultConfig.WithRusLang,
				err,
			)
		}
	}

}

func TestCreateAlphabet(t *testing.T) {
	var testConf = Config{1, false, false, false, false, false, ""}
	t.Log("Given the need to test CreateAlphabet behavior at different params.")
	{
		testID := 0
		t.Logf("\tTest %d:\t When WithSymbols is true.", testID)
		{
			testConf.WithSymbols = true
			testAlphabet := CreateAlphabet(&testConf)
			assert.Equalf(t, DefaultSymbols, testAlphabet, "\tAlphabet Should be %s \n, \tactual is %s  ",
				DefaultSymbols, testAlphabet)
			testConf.WithSymbols = false

		}
		testID++
		t.Logf("\tTest %d:\t When WithNumbers is true.", testID)
		{
			testConf.WithNumbers = true
			testAlphabet := CreateAlphabet(&testConf)
			assert.Equalf(t, DefaultNumberSet, testAlphabet, "\tAlphabet Should be %s \n, \tactual is %s  ",
				DefaultNumberSet, testAlphabet)
			testConf.WithNumbers = false
		}
		testID++
		t.Logf("\tTest %d:\t When WithUpperCase and WithEnglang is true.", testID)
		{
			testConf.WithUpperLetters = true
			testConf.WithEngLang = true
			testAlphabet := CreateAlphabet(&testConf)
			assert.Equalf(t, DefaultEngAlphabet+strings.ToUpper(DefaultEngAlphabet), testAlphabet,
				"\tAlphabet Should be %s \n, \tactual is %s  ",
				DefaultEngAlphabet+strings.ToUpper(DefaultEngAlphabet), testAlphabet)
			testConf.WithUpperLetters = false
			testConf.WithEngLang = false
		}

		t.Logf("\tTest %d:\t When WithUpperCase and WithRuslang is true.", testID)
		{
			testConf.WithUpperLetters = true
			testConf.WithRusLang = true
			testAlphabet := CreateAlphabet(&testConf)
			assert.Equalf(t, DefaultRusAlphabet+strings.ToUpper(DefaultRusAlphabet), testAlphabet,
				"\tAlphabet Should be %s \n, \tactual is %s  ",
				DefaultRusAlphabet+strings.ToUpper(DefaultRusAlphabet), testAlphabet)
			testConf.WithUpperLetters = false
			testConf.WithRusLang = false
		}

	}

}

//TODO continue some test
