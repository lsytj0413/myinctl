package charset

import (
	"github.com/pkg/errors"
)

type factory struct{}

func (f *factory) FindCharset(charsetName string) (Charset, error) {
	for _, v := range charsetList {
		if v.Name == charsetName {
			return *v, nil
		}
	}

	return Charset{}, errors.Errorf("Unknown charset name: %v", charsetName)
}

func (f *factory) FindCollation(collationName string) (Collation, error) {
	for _, v := range collationList {
		if v.Name == collationName {
			return *v, nil
		}
	}

	return Collation{}, errors.Errorf("Unknown collation name: %v", collationName)
}

func (f *factory) FindCollationsFromCharset(charsetName string) []Collation {
	ret := make([]Collation, 0)
	for _, v := range collationList {
		if v.Charset == charsetName {
			ret = append(ret, *v)
		}
	}

	return ret
}

func (f *factory) FindCollationFromID(id int) (Collation, error) {
	for _, v := range collationList {
		if v.ID == id {
			return *v, nil
		}
	}

	return Collation{}, errors.Errorf("Unknown collation id: %v", id)
}

// DefaultFactory is the Factory Impl
var DefaultFactory Factory

func init() {
	DefaultFactory = &factory{}
}
