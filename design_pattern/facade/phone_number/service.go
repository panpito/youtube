package phone_number

import (
	"github.com/panpito/youtube/design_pattern/facade/phone_number_validator"
	"log"
)

type ValidatorService interface {
	IsValidNumber(number string)
}

type Service struct {
	lengthChecker        phone_number_validator.LengthChecker
	prefixValidator      phone_number_validator.PrefixValidator
	countryCodeValidator phone_number_validator.CountryCodeValidator
}

func NewService(lengthChecker phone_number_validator.LengthChecker, prefixValidator phone_number_validator.PrefixValidator, countryCodeValidator phone_number_validator.CountryCodeValidator) *Service {
	return &Service{lengthChecker: lengthChecker, prefixValidator: prefixValidator, countryCodeValidator: countryCodeValidator}
}

func (service Service) IsValidNumber(number string) {
	okLength := service.lengthChecker.IsValid(number)
	log.Print("length is valid ", okLength)

	okPrefix := service.prefixValidator.IsValid(number)
	log.Print("prefix is valid ", okPrefix)

	okCountryCode := service.countryCodeValidator.IsValid(number)
	log.Print("country code is valid ", okCountryCode)
}
