package documents

import (
	"fmt"
	"math/rand"
)

func GenerateCpf() string {
	var validCpf string
	for {
		// TODO: refactor to generate first 9 digits and generate the another two
		generatedCpf := fmt.Sprintf("%011d", rand.Int63n(100000000000))

		if IsValidCpf(generatedCpf) {
			validCpf = generatedCpf
			break
		}
	}

	return validCpf
}

func IsValidCpf(document string) bool {
	if len(document) != 11 {
		return false
	}
	if allDigitsEqual(document) {
		return false
	}

	var tempTotal, firstDigit, secondDigit int
	var finalCpf string
	for i := 0; i < 9; i++ {
		tempTotal += int(document[i]-48) * (10 - i)
	}

	tempTotal = tempTotal % 11

	if tempTotal < 2 {
		firstDigit = 0
	} else {
		firstDigit = 11 - tempTotal
	}

	finalCpf = fmt.Sprintf("%s%d", document[:9], firstDigit)

	tempTotal = 0
	for i := 0; i < 10; i++ {
		tempTotal += int(finalCpf[i]-48) * (11 - i)
	}

	tempTotal = tempTotal % 11

	if tempTotal < 2 {
		secondDigit = 0

	} else {
		secondDigit = 11 - tempTotal
	}

	finalCpf = fmt.Sprintf("%s%d", finalCpf, secondDigit)

	if document == finalCpf {
		return true
	}
	return false
}

func allDigitsEqual(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}
