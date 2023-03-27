package plvalidator

import "errors"

type IterationChecksumMod func(checksum int, weigthel int, digit int) int
type FinalChecksumMod func(checksum int) int

func Modulo11(checksum int) int {
	checksum = (checksum % 11) % 10
	return checksum
}

func MultipleByWeigthAndAddToChecksum(checksum int, weigthel int, digit int) int {
	el := (weigthel * digit)
	checksum += el
	return checksum
}

func SplitToInt(number string) ([]int, error) {
	result := []int{}
	for _, ru := range number {
		digit, err := RuneToDigit(ru)
		if err != nil {
			return nil, err
		}
		result = append(result, digit)
	}
	return result, nil
}
func RuneToDigit(ru int32) (int, error) {
	digit := int(ru - '0')
	if digit < 0 || digit > 9 {
		return 0, errors.New("NaN")
	}
	return digit, nil
}

func CheckSum(weight []int, digit []int, iter IterationChecksumMod, final FinalChecksumMod) int {
	checksum := 0
	for key, val := range weight {
		checksum = iter(checksum, val, digit[key])
	}
	checksum = final(checksum)
	return checksum
}

func Contains[T comparable](val T, slice []T) bool {
	for _, el := range slice {
		if el == val {
			return true
		}
	}
	return false
}
