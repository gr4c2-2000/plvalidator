package plvalidator

import "errors"

type IterationChecksumMod func(checksum int, weigthel int, digit int) int
type FinalChecksumMod func(checksum int) int

func SplitToInt(number string) ([]int, error) {
	result := []int{}
	for _, ru := range number {
		digit := int(ru - '0')
		if digit < 0 || digit > 9 {
			return nil, errors.New("NaN")
		}
		result = append(result, digit)
	}
	return result, nil
}

func CheckSum(weight []int, digit []int, iter IterationChecksumMod, final FinalChecksumMod) int {
	checksum := 0
	for key, val := range weight {
		checksum = iter(checksum, val, digit[key])
	}
	checksum = final(checksum)
	return checksum
}
