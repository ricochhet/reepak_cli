package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/ricochhet/reepak"
	"github.com/ricochhet/reepak_cli/internal"
	"github.com/ricochhet/simplelog"
	"github.com/ricochhet/simpleutil"
)

//nolint:gocognit,gocyclo,cyclop,lll // wontfix
func main() {
	if _, err := simpleutil.NewCommand(os.Args, "help", 0); err == nil {
		simplelog.SharedLogger.Info("conv <hex-to-other/other-to-hex[conv mode]> <input> <string/decimal[conv type]>")
		simplelog.SharedLogger.Info("patch <file> <bytesToFind> <bytesToReplace> <replaceAtOccurrence>")
		simplelog.SharedLogger.Info("hash <file> <murmur3x64_128hash/murmur3x86_128hash/murmur3x86_32hash/crc64/crc32/sha512/sha256/sha1/md5[mode]>")
		simplelog.SharedLogger.Info("diff <folderA> <folderB>")
		simplelog.SharedLogger.Info("reepak:")
		simplelog.SharedLogger.Info("\tpak <folder> <output> <0/1[embed data]>")
		simplelog.SharedLogger.Info("\tunpak <folder> <output> <0/1[embed data]>")
		simplelog.SharedLogger.Info("\tcompress <file>")
		simplelog.SharedLogger.Info("\tdecompress <file>")
		os.Exit(1)
	}

	if args, err := simpleutil.NewCommand(os.Args, "conv", 3); err == nil {
		if err := internal.NewConvert(args); err != nil {
			simplelog.SharedLogger.Fatal(err.Error())
		}
	} else if !errors.Is(err, simpleutil.ErrNoFunctionName) {
		simplelog.SharedLogger.Fatal(err.Error())
	}

	if args, err := simpleutil.NewCommand(os.Args, "patch", 2); err == nil {
		if err := internal.NewPatch(args); err != nil {
			simplelog.SharedLogger.Fatal(err.Error())
		}
	} else if !errors.Is(err, simpleutil.ErrNoFunctionName) {
		simplelog.SharedLogger.Fatal(err.Error())
	}

	if args, err := simpleutil.NewCommand(os.Args, "hash", 1); err == nil {
		if err := internal.NewHash(args); err != nil {
			simplelog.SharedLogger.Fatal(err.Error())
		}
	} else if !errors.Is(err, simpleutil.ErrNoFunctionName) {
		simplelog.SharedLogger.Fatal(err.Error())
	}

	if args, err := simpleutil.NewCommand(os.Args, "diff", 2); err == nil {
		if err := internal.NewDiff(args); err != nil {
			simplelog.SharedLogger.Fatal(err.Error())
		}
	} else if !errors.Is(err, simpleutil.ErrNoFunctionName) {
		simplelog.SharedLogger.Fatal(err.Error())
	}

	if args, err := simpleutil.NewCommand(os.Args, "pak", 3); err == nil {
		selection, err := strconv.Atoi(args[2])
		if err != nil {
			simplelog.SharedLogger.Fatalf("Error converting string to integer: %s", err)
		}

		if err := reepak.ProcessDirectory(args[0], args[1], selection != 0); err != nil {
			simplelog.SharedLogger.Fatalf("Error processing directory: %s", err)
		}
	} else if !errors.Is(err, simpleutil.ErrNoFunctionName) {
		simplelog.SharedLogger.Fatal(err.Error())
	}

	if args, err := simpleutil.NewCommand(os.Args, "unpak", 3); err == nil {
		selection, err := strconv.Atoi(args[2])
		if err != nil {
			simplelog.SharedLogger.Fatalf("Error converting string to integer: %s", err)
		}

		if err := reepak.ExtractDirectory(args[0], args[1], selection != 0); err != nil {
			simplelog.SharedLogger.Fatalf("Error extracting directory: %s", err)
		}
	} else if !errors.Is(err, simpleutil.ErrNoFunctionName) {
		simplelog.SharedLogger.Fatal(err.Error())
	}

	if args, err := simpleutil.NewCommand(os.Args, "compress", 1); err == nil {
		if err := reepak.CompressPakData(args[0]); err != nil {
			simplelog.SharedLogger.Fatalf("Error compressing pak data: %s", err)
		}
	} else if !errors.Is(err, simpleutil.ErrNoFunctionName) {
		simplelog.SharedLogger.Fatal(err.Error())
	}

	if args, err := simpleutil.NewCommand(os.Args, "decompress", 1); err == nil {
		if err := reepak.DecompressPakData(args[0]); err != nil {
			simplelog.SharedLogger.Fatalf("Error decompressing pak data: %s", err)
		}
	} else if !errors.Is(err, simpleutil.ErrNoFunctionName) {
		simplelog.SharedLogger.Fatal(err.Error())
	}
}
