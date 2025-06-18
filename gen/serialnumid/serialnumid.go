package serialnumid

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type GenerateParam struct {
	StartNum    int
	NumCount    int
	CleanNum    bool
	TmpDataFile string
	OutputFile  string
	UsingId     int
}

func DumpId(ids []string, param *GenerateParam) error {
	var writer io.Writer = os.Stdout
	if param.OutputFile != "" {
		f, err := os.Create(param.OutputFile)
		if err != nil {
			fmt.Printf("Failed to create output file: %v", err)
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

	for i := 0; i < param.NumCount; i++ {
		values = append(values, fmt.Sprintf("%d", param.UsingId))
		param.UsingId++
	}
	return values, nil
}

func Validate(param *GenerateParam) error {
	return nil
}

func ReadSerialNumber(param *GenerateParam) error {
	param.UsingId = 0
	file, err := os.Open(param.TmpDataFile)
	if errors.Is(err, os.ErrNotExist) {
		param.UsingId = 1 // start from 1 when not exists
		return nil
	} else if err != nil {
		return err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var n int
	for sc.Scan() {
		if n == 0 {
			line := sc.Text()
			param.UsingId, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println("Failed converting serial numberfrom file content")
				return err
			}
		}
		n++
	}
	return nil
}

func WriteSerialNumber(param *GenerateParam) error {
	f, err := os.Create(param.TmpDataFile)
	if err != nil {
		fmt.Println("Failed accessing temp serial number file")
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "%d\n", param.UsingId)

	return nil
}
