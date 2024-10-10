package internal

import (
	"errors"

	"github.com/ricochhet/simplecrypto"
	"github.com/ricochhet/simplefs"
	"github.com/ricochhet/simplelog"
)

var (
	errPathANoExist = errors.New("first path specified does not exist")
	errPathBNoExist = errors.New("second path specified does not exist")
)

func NewDiff(args []string) error {
	if !simplefs.Exists(args[0]) {
		return errPathANoExist
	}

	if !simplefs.Exists(args[1]) {
		return errPathBNoExist
	}

	dirA, err := simplecrypto.HashDirectory(args[0])
	if err != nil {
		return err
	}

	dirB, err := simplecrypto.HashDirectory(args[1])
	if err != nil {
		return err
	}

	data := simplecrypto.DiffDirectory(dirA, dirB, args[0], args[1])

	for _, diff := range data {
		if diff.Local != (simplecrypto.DiffLocalData{}) { //nolint:exhaustruct // wontfix
			simplelog.SharedLogger.Infof("File: %s exists in %s, but not in %s", diff.Local.Path, diff.Local.ExistsA, diff.Local.ExistsB)
		}

		if diff.Hashes != (simplecrypto.DiffHashData{}) { //nolint:exhaustruct // wontfix
			simplelog.SharedLogger.Infof("Hashes for file: %s do not match:", diff.Hashes.File)
			simplelog.SharedLogger.Infof("\t%s: %s", diff.Hashes.PathA, diff.Hashes.HashA)
			simplelog.SharedLogger.Infof("\t%s: %s", diff.Hashes.PathB, diff.Hashes.HashB)
		}
	}

	return nil
}
