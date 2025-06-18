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
	SartNum     int
	NumCount    int
	CleanNum    bool
	TmpDataFile string
	OutputFile  string
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

	sn, err := ReadSerialNumber(param)
	if err != nil {
		fmt.Println("Failed parsing persited serial number from file")
		return values, err
	}

	for i := 0; i < param.NumCount; i++ {
		values = append(values, fmt.Sprintf("%d", sn))
		sn++
	}
	err = WriteSerialNumber(sn, param)
	if err != nil {
		fmt.Println("Failed saving current serial number into file")
		return values, err
	}

	return values, nil
}

func Validate(param *GenerateParam) error {
	return nil
}

func ReadSerialNumber(param *GenerateParam) (serialNum int, err error) {
	serialNum = 0
	file, err := os.Open(param.TmpDataFile)
	if errors.Is(err, os.ErrNotExist) {
		serialNum = 1 // start from 1 when not exists
		return serialNum, nil
	} else if err != nil {
		return serialNum, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var n int
	for sc.Scan() {
		if n == 0 {
			line := sc.Text()
			serialNum, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println("Failed converting serial numberfrom file content")
				return serialNum, err
			}
		}
		n++
	}
	return serialNum, nil
}

func WriteSerialNumber(serialNum int, param *GenerateParam) error {
	f, err := os.Create(param.TmpDataFile)
	if err != nil {
		fmt.Println("Failed accessing temp serial number file")
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "%d\n", serialNum)

	return nil
}
