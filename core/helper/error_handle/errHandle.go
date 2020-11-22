package error_handle

import "log"

func ErrHandle(err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return err
}
