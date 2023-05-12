package phone_number

import "github.com/panpito/youtube/design_pattern/facade/phone_number_validator"

type Consumer struct {
	validator ValidatorService
}

func NewConsumer() *Consumer {
	lengthChecker := phone_number_validator.NewLengthChecker(10)
	prefixValidator := phone_number_validator.NewPrefixValidator()
	countryCodeValidator := phone_number_validator.NewCountryCodeValidator("44")

	validator := NewService(*lengthChecker, *prefixValidator, *countryCodeValidator)

	return &Consumer{validator: validator}
}

func (consumer Consumer) Process(number string) {
	consumer.validator.IsValidNumber(number)
}
