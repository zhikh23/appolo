package e

import "fmt"

func WrapError(service string, err error) error {
	return fmt.Errorf("%s: %s", service, err)
}

func WrapIfError(service string, err error) error {
	if err != nil {
		err = fmt.Errorf("%s: %s", service, err)
	}
	return err
}
