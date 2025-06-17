package serialnumid

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

	sn, err := ReadSerialNumber(param)
	if err != nil {
		log.Fatal("Failre parse peristed serial number from file")
		return values, err
	}

	for i := 0; i < param.NumCount; i++ {
		sn++
		values = append(values, fmt.Sprintf("%s\n", sn))
	}

	return values, nil
}

func Validate(param *GenerateParam) error {
	return nil
}

func ReadSerialNumber(param *GenerateParam) (serialNum int, err error) {
	file, err := os.Open(param.TmpDataFile)
	serialNum = 0
	if err != nil {
		return serialNum, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		serialNum, err = strconv.Atoi(sc.Text())
	}
	return serialNum, io.EOF
}

func WriteSerialNumber(serialNum int, param *GenerateParam) error {
	f, err := os.Create(param.TmpDataFile)
	if err != nil {
		log.Fatal("Failed to access temp serial number file")
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "%d\n", serialNum)

	return nil
}
