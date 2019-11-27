package ntw

import (
	"fmt"
	"testing"
)

func TestIntegerToRuRu(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:            "минус один",
		0:             "нуль",
		1:             "один",
		9:             "девять",
		10:            "десять",
		11:            "одиннадцать",
		19:            "девятнадцать",
		20:            "двадцать",
		21:            "двадцать один",
		80:            "восемьдесят",
		90:            "девяносто",
		99:            "девяносто девять",
		100:           "сто",
		101:           "сто один",
		111:           "сто одиннадцать",
		120:           "сто двадцать",
		121:           "сто двадцать один",
		900:           "девятьсот",
		909:           "девятьсот девять",
		919:           "девятьсот девятнадцать",
		990:           "девятьсот девяносто",
		999:           "девятьсот девяносто девять",
		1000:          "одна тысяча",
		2000:          "две тысячи",
		4000:          "четыре тысячи",
		5000:          "пять тысяч",
		11000:         "одиннадцать тысяч",
		21000:         "двадцать одна тысяча",
		999000:        "девятьсот девяносто девять тысяч",
		999999:        "девятьсот девяносто девять тысяч девятьсот девяносто девять",
		1000000:       "один миллион",
		2000000:       "два миллиона",
		4000000:       "четыре миллиона",
		5000000:       "пять миллионов",
		100100100:     "сто миллионов сто тысяч сто",
		500500500:     "пятьсот миллионов пятьсот тысяч пятьсот",
		606606606:     "шестьсот шесть миллионов шестьсот шесть тысяч шестьсот шесть",
		999000000:     "девятьсот девяносто девять миллионов",
		999000999:     "девятьсот девяносто девять миллионов девятьсот девяносто девять",
		999999000:     "девятьсот девяносто девять миллионов девятьсот девяносто девять тысяч",
		999999999:     "девятьсот девяносто девять миллионов девятьсот девяносто девять тысяч девятьсот девяносто девять",
		1174315110:    "один миллиард сто семьдесят четыре миллиона триста пятнадцать тысяч сто десять",
		1174315119:    "один миллиард сто семьдесят четыре миллиона триста пятнадцать тысяч сто девятнадцать",
		15174315119:   "пятнадцать миллиардов сто семьдесят четыре миллиона триста пятнадцать тысяч сто девятнадцать",
		35174315119:   "тридцать пять миллиардов сто семьдесят четыре миллиона триста пятнадцать тысяч сто девятнадцать",
		935174315119:  "девятьсот тридцать пять миллиардов сто семьдесят четыре миллиона триста пятнадцать тысяч сто девятнадцать",
		2935174315119: "два триллиона девятьсот тридцать пять миллиардов сто семьдесят четыре миллиона триста пятнадцать тысяч сто девятнадцать",
		3911760:       "три миллиона девятьсот одиннадцать тысяч семьсот шестьдесят",
		27:            "двадцать семь",
		95000001000:   "девяносто пять миллиардов одна тысяча",
		57482:         "пятьдесят семь тысяч четыреста восемьдесят два",
		5:             "пять",
		16:            "шестнадцать",
		30:            "тридцать",
		53:            "пятьдесят три",
		123:           "сто двадцать три",
		204:           "двести четыре",
		300:           "триста",
		1400:          "одна тысяча четыреста",
		83756:         "восемьдесят три тысячи семьсот пятьдесят шесть",
		293111:        "двести девяносто три тысячи сто одиннадцать",
		32001950:      "тридцать два миллиона одна тысяча девятьсот пятьдесят",
		2018:          "две тысячи восемнадцать",
		126682:        "сто двадцать шесть тысяч шестьсот восемьдесят два",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			output := IntegerToRuRu(input)
			if expectedOutput != output {
				t.Fatalf("Expected %q, got %q.", expectedOutput, output)
			}
		})
	}
}
