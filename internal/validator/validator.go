package validator
import (
	"regexp"
)

var (
	EmailRX = regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//validation struct contains map of validation errors
type Validator struct {
	Errors map[string] string
}

//helper that creates new Validator with empty errors map
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

//returns true if errors map has no entries
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}
//adds error message to map (as long as it doesn't exist for given key)
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

//adds error message to map if validation check is not 'ok'
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

//returns true if specific value is in a list
func PermittedValue[T comparable] (value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true 
		}
	}
	return false
}

//returns true if string value matches a regexp pattern
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

//generic function returns true if all values in slice are unique
func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, value := range values {
		uniqueValues[value] true
	}

	return len(values) == len(uniqueValues)
}

