package plvalidator

// Package pesel implements a validation and parsing of Polish personal identification number (PESEL).
import (
	"errors"
	"time"
)

var peselTimeBoundries map[int]int = map[int]int{1800: 80, 1900: 0, 2000: 20, 2100: 40, 2200: 60}
var weightPesel []int = []int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}

const PESEL_LENGTH = 11
const FEMALE = "Female"
const MALE = "Male"

func PeselChecksumIteration(checksum int, weigthel int, digit int) int {
	el := (weigthel * digit) % 10
	checksum += el
	return checksum
}

func PeselChecksumFinal(checksum int) int {
	checksum = 10 - (checksum % 10)
	return checksum
}

// pesel is a struct representation of Polish personal identification number.
// A correct PESEL contains information about the birth date and gender.
type pesel struct {
	Id       string
	splited  []int
	birthday time.Time
	gender   string
	valid    bool
}

// Valid returns true if the PESEL is valid, otherwise false.
func (p *pesel) Valid() bool {
	return p.valid
}

// Gender returns the gender (Female or Male) associated with the PESEL.
func (p *pesel) Gender() string {
	return p.gender
}

// Birthday returns the birthday associated with the PESEL.
func (p *pesel) Birthday() time.Time {
	return p.birthday
}

// https://obywatel.gov.pl/pl/dokumenty-i-dane-osobowe/czym-jest-numer-pesel
// Valid returns true if the provided string is a valid PESEL, otherwise false.
func Pesel(id string) bool {
	_, err := NewPesel(id)
	return err == nil
}

// Pesel returns a *pesel struct parsed from the provided string.
// Returns an error if the provided string is not a valid PESEL.
func NewPesel(id string) (*pesel, error) {
	p := pesel{Id: id}
	err := p.charatters()
	if err != nil {
		return nil, err
	}
	err = p.length()
	if err != nil {
		return nil, err
	}
	err = p.date()
	if err != nil {
		return nil, err
	}
	err = p.checksum()
	if err != nil {
		return nil, err
	}
	if p.splited[9]%2 == 0 {
		p.gender = FEMALE
	} else {
		p.gender = MALE
	}
	p.valid = true
	return &p, nil
}

// date sets the birth date for the *pesel struct based on the first six digits of the provided PESEL.
func (p *pesel) date() error {
	y := p.splited[0]*10 + p.splited[1]
	m := p.splited[2]*10 + p.splited[3]
	day := p.splited[4]*10 + p.splited[5]
	year := 0
	month := 0
	for ye, sum := range peselTimeBoundries {
		if m > (sum) && m <= (sum+12) {
			year = y + ye
			month = m - sum
		}

	}
	if year == 0 || month == 0 || !(day > 1 && day < 32) {
		return errors.New("INCCORECT_DATE")
	}
	p.birthday = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return nil
}

// length checks if the provided PESEL is the correct length (11 characters).
func (p *pesel) length() error {
	if len(p.Id) == PESEL_LENGTH {
		return nil
	}
	return errors.New("INCORECT_LENGTH")
}

// charatters converts the provided PESEL to numeric and splits it into digits.
func (p *pesel) charatters() error {
	var err error
	p.splited, err = SplitToInt(p.Id)
	if err != nil {
		return err
	}
	return nil

}

// checksum calculates and checks the checksum digit of the PESEL number.
// If the calculated checksum matches the last digit of the PESEL number, it returns nil,
// otherwise it returns an error.
func (p *pesel) checksum() error {
	imputcs := p.splited[PESEL_LENGTH-1]
	if imputcs == p.Checksum() {
		return nil
	}
	return errors.New("INCORECT_CHECKSUM")
}

// Checksum calculates and returns the checksum digit of the PESEL number.
func (p *pesel) Checksum() int {

	return CheckSum(weightPesel, p.splited, PeselChecksumIteration, PeselChecksumFinal)
}
