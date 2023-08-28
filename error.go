package config

import "strconv"

type confErrors struct {
	errs []error
}

func newConfError() *confErrors {
	return &confErrors{
		errs: make([]error, 0),
	}
}

func (ce *confErrors) appendError(err error) {
	ce.errs = append(ce.errs, err)
}

func (ce *confErrors) Error() string {
	errmsg := "load configuration failed, there are some errors:\n"

	for i, err := range ce.errs {
		prefix := "\t[ERROR]    " + strconv.FormatInt(int64(i), 10) + ":\t"
		errmsg += prefix + err.Error() + "\n"
	}
	return errmsg
}

func (ce *confErrors) isNil() bool {
	return len(ce.errs) == 0
}
