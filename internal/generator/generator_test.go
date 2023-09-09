package generator

import (
	"github.com/stretchr/testify/assert"
	"regexp"
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
	var tests = []Config{
		{
			Length:           1,
			WithSymbols:      true,
			WithNumbers:      false,
			WithUpperLetters: false,
			WithEngLang:      false,
			WithRusLang:      false,
			Alphabet:         "",
		},
		{
			Length:           1,
			WithSymbols:      false,
			WithNumbers:      true,
			WithUpperLetters: false,
			WithEngLang:      false,
			WithRusLang:      false,
			Alphabet:         "",
		},
		{
			Length:           1,
			WithSymbols:      false,
			WithNumbers:      false,
			WithUpperLetters: true,
			WithEngLang:      true,
			WithRusLang:      false,
			Alphabet:         "",
		},
		{
			Length:           1,
			WithSymbols:      false,
			WithNumbers:      false,
			WithUpperLetters: true,
			WithEngLang:      false,
			WithRusLang:      true,
			Alphabet:         "",
		},
		{
			Length:           1,
			WithSymbols:      true,
			WithNumbers:      true,
			WithUpperLetters: true,
			WithEngLang:      true,
			WithRusLang:      true,
			Alphabet:         "",
		},
	}
	t.Log("Given the need to test CreateAlphabet behavior at different params.")
	{
		testID := 0
		t.Logf("\tTest %d:\t When WithSymbols is true\n%s", testID, tests[testID].String())
		{
			testAlphabet := CreateAlphabet(&tests[testID])
			regex := regexp.MustCompile(`[[:punct:]]`).MatchString
			pass := assert.Equalf(t, true, regex(testAlphabet),
				"\tAlphabet Should contains symbols \n actual is %s", testAlphabet)
			if pass {
				t.Logf("\tactual alphabet is %s\n", testAlphabet)
			}

		}
		testID++
		t.Logf("\tTest %d:\t When WithNumbers is true\n%s", testID, tests[testID].String())
		{
			testAlphabet := CreateAlphabet(&tests[testID])
			regex := regexp.MustCompile(`[0-9]`).MatchString
			pass := assert.Equalf(t, true, regex(testAlphabet),
				"\tAlphabet Should contains numbers \n actual is %s", testAlphabet)
			if pass {
				t.Logf("\tactual alphabet is %s\n", testAlphabet)
			}
		}
		testID++
		t.Logf("\tTest %d:\t When WithUpperCase and WithEnglang is true\n%s", testID, tests[testID].String())
		{
			testAlphabet := CreateAlphabet(&tests[testID])
			regex := regexp.MustCompile(`[a-zA-Z]`).MatchString
			pass := assert.Equalf(t, true, regex(testAlphabet),
				"\tAlphabet Should contains lower and upper case English symbols \n actual is %s", testAlphabet)
			if pass {
				t.Logf("\tactual alphabet is %s\n", testAlphabet)
			}
		}

		testID++
		t.Logf("\tTest %d:\t When WithUpperCase and WithRuslang is true\n%s", testID, tests[testID].String())
		{
			testAlphabet := CreateAlphabet(&tests[testID])
			regex := regexp.MustCompile(`[а-яА-Я]`).MatchString
			pass := assert.Equalf(t, true, regex(testAlphabet),
				"\tAlphabet Should contains lower and upper case Rulang symbols\n actual is %s", testAlphabet)
			if pass {
				t.Logf("\tactual alphabet is %s\n", testAlphabet)
			}
		}

		testID++
		t.Logf("\tTest %d:\t When all params is true \n%s", testID, tests[testID].String())
		{
			testAlphabet := CreateAlphabet(&tests[testID])
			regex := regexp.MustCompile(`[а-яА-Я0-9a-zA-Z[:punct:]]`).MatchString
			pass := assert.Equalf(t, true, regex(testAlphabet),
				"\tAlphabet Should contains any symbols \n actual is %s", testAlphabet)
			if pass {
				t.Logf("\tactual alphabet is %s\n", testAlphabet)
			}
		}

	}

}

func TestGenerator_Generate(t *testing.T) {
	var tests = []Config{
		DefaultConfig,
		{
			Length:           8,
			WithSymbols:      false,
			WithNumbers:      true,
			WithUpperLetters: true,
			WithEngLang:      true,
			WithRusLang:      false,
			Alphabet:         "",
		},
		{
			Length:           8,
			WithSymbols:      false,
			WithNumbers:      true,
			WithUpperLetters: true,
			WithEngLang:      false,
			WithRusLang:      true,
			Alphabet:         "",
		},
		{
			Length:           24,
			WithSymbols:      true,
			WithNumbers:      true,
			WithUpperLetters: true,
			WithEngLang:      true,
			WithRusLang:      true,
			Alphabet:         "",
		},
		{
			Length:           uint(len(DefaultNumberSet) + len(DefaultEngAlphabet) + 1),
			WithSymbols:      false,
			WithNumbers:      true,
			WithUpperLetters: false,
			WithEngLang:      true,
			WithRusLang:      false,
			Alphabet:         "",
		},
		{
			Length:           uint(len(DefaultNumberSet) + len(DefaultEngAlphabet)*2 + len(DefaultSymbols) + len(DefaultRusAlphabet)*2),
			WithSymbols:      true,
			WithNumbers:      true,
			WithUpperLetters: true,
			WithEngLang:      true,
			WithRusLang:      true,
			Alphabet:         "",
		},
	}
	t.Log("Given the need to test CreateGenerator_Generate behavior at different params.")
	{
		testID := 0
		t.Logf("\tTest %d:\t When Generator.Config == DefaultConfig \n%s", testID, tests[testID].String())
		{
			testGenerator, _ := New(&tests[testID])
			testPassword, err := testGenerator.Generate()
			regex := regexp.MustCompile(`^[a-z0-9]+$`).MatchString
			pass := assert.Equalf(t, true, regex(*testPassword),
				"\tpassword must contain only lower case English alphabet letters and digit \n, "+
					"\tactual alphabet is %s \n \tactual password is %s \n \t error is : %s",
				testGenerator.Alphabet, *testPassword, err)
			if pass {
				t.Logf("\tactual alphabet is %s \n \t\t\t\t\tactual password is %s \n",
					testGenerator.Alphabet, *testPassword)
			}

		}
		testID++
		t.Logf("\tTest %d:\t When Generator.Config with: EngLang, UpperCase, Numbers \n%s",
			testID, tests[testID].String())
		{
			testGenerator, _ := New(&tests[testID])
			testPassword, err := testGenerator.Generate()
			regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
			pass := assert.Equalf(t, true, regex(*testPassword),
				"\tpassword must contain only lower and upper case English alphabet letters and digit \n, "+
					"\tactual alphabet is %s \n \tactual password is %s \n \t error is : %s",
				testGenerator.Alphabet, *testPassword, err)
			if pass {
				t.Logf("\tactual alphabet is %s \n \t\t\t\t\tactual password is %s \n",
					testGenerator.Alphabet, *testPassword)
			}

		}

		testID++
		t.Logf("\tTest %d:\t When Generator.Config with: RusLang, UpperCase, Numbers\n%s",
			testID, tests[testID].String())
		{

			testGenerator, _ := New(&tests[testID])
			testPassword, err := testGenerator.Generate()
			regex := regexp.MustCompile(`[а-яА-Я0-9]`).MatchString
			pass := assert.Equalf(t, true, regex(*testPassword),
				"\tpassword must contain only lower and upper case RU alphabet letters and digit \n, "+
					"\tactual alphabet is %s \n \t actual password is %s \n \t error is : %e",
				testGenerator.Alphabet, *testPassword, err)
			if pass {
				t.Logf("\tactual alphabet is %s \n \t\t\t\t\tactual password is %s \n",
					testGenerator.Alphabet, *testPassword)
			}

		}

		testID++
		t.Logf("\tTest %d:\t When Generator.Config with: EngLang, RuLang, UpperCase, Number,Specific symbols"+
			"\n%s", testID, tests[testID].String())
		{
			testGenerator, _ := New(&tests[testID])
			testPassword, err := testGenerator.Generate()
			regex := regexp.MustCompile(`[а-яА-Я0-9a-zA-Z[:punct:]]`).MatchString
			pass := assert.Equalf(t, true, regex(*testPassword),
				"\tpassword must contain any RU and End letters, digits and specific symbols \n, "+
					"\tactual alphabet is %s \n \t actual password is %s \n \t error is : %e",
				testGenerator.Alphabet, *testPassword, err)
			if pass {
				t.Logf("\tactual alphabet is %s \n \t\t\t\t\tactual password is %s \n",
					testGenerator.Alphabet, *testPassword)
			}

		}

		testID++
		t.Logf("\tTest %d:\t When Generator.Config with: EngLang, Numbers and "+
			"password length longer than alphabet length\n%s", testID, tests[testID].String())
		{
			testGenerator, _ := New(&tests[testID])
			testPassword, err := testGenerator.Generate()
			regex := regexp.MustCompile(`[0-9a-z]`).MatchString
			pass := assert.Equalf(t, true, regex(*testPassword),
				"\tpassword must contain only lower case END alphabet letters and digits \n, "+
					"\tactual alphabet is %s \n \t actual password is %s \n \t error is : %e",
				testGenerator.Alphabet, *testPassword, err)
			if pass {
				t.Logf("\tactual alphabet is %s \n \t\t\t\t\tactual password is %s \n",
					testGenerator.Alphabet, *testPassword)
			}

		}

		testID++
		t.Logf("\tTest %d:\t When Generator.Config with: RuLang, Englang, UpperCase, Numbers, Specific symbols"+
			"and password length longer than alphabet length\n%s", testID, tests[testID].String())
		{
			testGenerator, _ := New(&tests[testID])
			testPassword, err := testGenerator.Generate()
			regex := regexp.MustCompile(`[а-яА-Я0-9a-zA-Z[:punct:]]`).MatchString
			pass := assert.Equalf(t, true, regex(*testPassword),
				"\tpassword must contain any RU and End letters, digits and specific symbols \n, "+
					"\tactual alphabet is %s \n \t actual password is %s \n \t error is : %e",
				testGenerator.Alphabet, *testPassword, err)
			if pass {
				t.Logf("\tactual alphabet is %s \n \t\t\t\t\tactual password is %s \n",
					testGenerator.Alphabet, *testPassword)
			}

		}

	}
}

//TODO continue some test
