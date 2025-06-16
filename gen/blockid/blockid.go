package blockid

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type GenerateParam struct {
	PrefixName   string
	PostfixName  string
	NumByteCount int
	NumBlockids  int
	OutputFile   string
	ToUppercases bool
}

func DumpId(ids []string, param *GenerateParam) error {
	var writer io.Writer = os.Stdout
	if param.OutputFile != "" {
		f, err := os.Create(param.OutputFile)
		if err != nil {
			log.Fatalf("Failed to create output file: %v", err)
			return err
		}
		defer f.Close()
		writer = f
	}

	for _, v := range ids {
		fmt.Fprintln(writer, v)
	}

	return nil
}

func GenerateId(param *GenerateParam) ([]string, error) {
	var values []string = []string{}

	for i := 0; i < param.NumBlockids; i++ {
		u, err := uuid.NewRandom()
		if err != nil {
			return values, errors.New("Failure generate basic id in blockid")
		}
		idValue := strings.ReplaceAll(u.String(), "-", "")
		idValue = idValue[0 : param.NumByteCount*2]

		if param.ToUppercases {
			idValue = strings.ToUpper(idValue)
		} else {
			idValue = strings.ToUpper(idValue)
		}

		idValue = param.PrefixName + idValue + param.PostfixName
		values = append(values, idValue)
	}

	return values, nil
}

func Validate(param *GenerateParam) error {
	if param.NumByteCount < 1 || param.NumByteCount > 16 {
		return errors.New(fmt.Sprintf("Failure. Want a number between 1 and 16 inclusive but got %d", param.NumByteCount))
	}
	return nil
}
