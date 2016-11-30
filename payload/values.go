package payload

import (
	"fmt"
	"net/url"
	"regexp"
)

type (
	// Values struct for parsing IPN form data
	// into data accessible for a payload consuming
	// the parser.
	Values struct {
		url.Values
	}
)

const (
	suffixTypeNumericSuffix           = "suffixTypeNumericSuffix"
	suffixTypeUnderscoreNumericSuffix = "suffixTypeUnderscoreNumericSuffix"
)

// NewValuesFromFormData creates new Values from form data.
func NewValuesFromFormData(formData string) (*Values, error) {
	parsedFormData, err := url.ParseQuery(formData)
	if err != nil {
		return nil, err
	}

	return &Values{parsedFormData}, nil
}

// Get returns the value for the given form key.
func (p *Values) Get(key string) string {
	return p.Values.Get(key)
}

// GetValues returns a slice of values based of the key
// prefix and the amount of records passed in.
func (p *Values) GetValues(key string, itemCount int) []string {
	var items []string

	for i := 0; i < itemCount; i++ {
		items = append(items, p.GetValueAtIndex(key, i))
	}

	return items
}

// GetValueAtIndex returns the value for an array key at a given
// index.
func (p *Values) GetValueAtIndex(key string, index int) string {
	suffixType, found := p.indexSuffixType(key)
	index++

	if found {
		switch suffixType {
		case suffixTypeNumericSuffix:
			return p.Get(fmt.Sprintf("%s%d", key, index))
		case suffixTypeUnderscoreNumericSuffix:
			return p.Get(fmt.Sprintf("%s_%d", key, index))
		}
	}

	return ""
}

// ItemCount returns count of items for a given array key prefix.
func (p Values) ItemCount(keyPrefix string) int {
	count := 0
	re := regexp.MustCompile(fmt.Sprintf(`%s_?\d$`, keyPrefix))
	for formDataKey, _ := range p.Values {
		if re.MatchString(formDataKey) {
			count++
		}
	}

	return count
}

func (p Values) indexSuffixType(keyPrefix string) (string, bool) {
	numericRegex := regexp.MustCompile(fmt.Sprintf(`%s\d$`, keyPrefix))
	underscoreRegex := regexp.MustCompile(fmt.Sprintf(`%s_\d$`, keyPrefix))

	for formDataKey, _ := range p.Values {
		if numericRegex.MatchString(formDataKey) {
			return suffixTypeNumericSuffix, true
		}

		if underscoreRegex.MatchString(formDataKey) {
			return suffixTypeUnderscoreNumericSuffix, true
		}
	}

	return "", false
}
