// Copyright (c) 2018. Flying Gopher Authors
// license that can be found in the LICENSE file.

package sessions

import (
	"os"
)

func CreateFileIfnExist(filepath string) error {
	_, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		_, err = os.Create(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}
