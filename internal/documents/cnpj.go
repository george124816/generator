package documents

import (
	"fmt"
	"math/rand"
)

var firstMatrixMulti = [12]int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
var secondMatrixMulti = [13]int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

func GeneraterCnpj() string {
	var validCnpj string
	for {
		// TODO: refactor to generate first 12 digits and generate the another two
		generatedCnpj := fmt.Sprintf("%014d", rand.Int63n(100000000000000))

		if IsValidCnpj(generatedCnpj) {
			validCnpj = generatedCnpj
			break
		}
	}
	return validCnpj
}

func IsValidCnpj(document string) bool {

	if len(document) != 14 {
		return false
	}

	var tempTotal, firstDigit, secondDigit int
	for i := 0; i < 12; i++ {
		tempTotal += int(document[i]-48) * firstMatrixMulti[i]
	}

	tempTotal = tempTotal % 11
	if tempTotal < 2 {
		firstDigit = 0
	} else {
		firstDigit = 11 - tempTotal
	}

	finalDocument := fmt.Sprintf("%s%d", document[:12], firstDigit)

	tempTotal = 0
	for i := 0; i < 13; i++ {
		tempTotal += int(finalDocument[i]-48) * secondMatrixMulti[i]
	}

	tempTotal = tempTotal % 11

	if tempTotal < 2 {
		secondDigit = 0

	} else {
		secondDigit = 11 - tempTotal
	}

	finalDocument = fmt.Sprintf("%s%d", finalDocument, secondDigit)

	if document == finalDocument {
		return true
	}

	return false
}
