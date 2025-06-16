package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type GenerateParam struct {
	UuidVersionFlag string
	NamespaceID     string
	NameData        string
	OutputFile      string
	NumUUIDs        int
	ToUppercases    bool
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
	var err error

	for i := 0; i < param.NumUUIDs; i++ {
		var u uuid.UUID
		var err error

		switch param.UuidVersionFlag {
		case "1":
			u, err = uuid.NewUUID()
			if err != nil {
				log.Fatalf("Failure generate new UUID: %v", err)
			}
		case "3":
			if param.NamespaceID == "" || param.NameData == "" {
				log.Fatal("For version 3 UUIDs, -n (namespace) and -N (name) are required.")
			}
			ns, parseErr := uuid.Parse(param.NamespaceID)
			if parseErr != nil {
				log.Fatalf("Invalid namespace UUID: %v", parseErr)
			}
			u = uuid.NewMD5(ns, []byte(param.NameData))
		case "4":
			u, err = uuid.NewRandom()
			if err != nil {
				log.Fatalf("Failure generate new UUID random value: %v", err)
			}
		case "5":
			if param.NamespaceID == "" || param.NameData == "" {
				log.Fatal("For version 5 UUIDs, -n (namespace) and -N (name) are required.")
			}
			ns, parseErr := uuid.Parse(param.NamespaceID)
			if parseErr != nil {
				log.Fatalf("Invalid namespace UUID: %v", parseErr)
			}
			u = uuid.NewSHA1(ns, []byte(param.NameData))
		case "6":
			u, err = uuid.NewV6()
		case "7":
			u, err = uuid.NewV7()
		default:
			log.Fatalf("Unsupported UUID version: %s. Supported versions are 1, 3, 4, 5, 6, 7.", param.UuidVersionFlag)
			err = errors.New("Unsupported UUID version")
		}

		if err != nil {
			log.Fatalf("Generaton failure for UUID: %v", err)
		}

		// Standardmäßig Kleinbuchstaben, wie bei den meisten uuidgen-Implementierungen
		// Für Großbuchstaben könnte man hier strings.ToUpper(u.String()) verwenden,
		// aber uuidgen gibt standardmäßig Kleinbuchstaben aus.
		if param.ToUppercases {
			values = append(values, strings.ToUpper(u.String()))
		} else {
			values = append(values, strings.ToLower(u.String()))
		}
	}

	return values, err
}

func Validate(param *GenerateParam) error {
	return nil
}

// Hilfsfunktion zur Hash-Berechnung für MD5/SHA1 (nicht direkt für uuid.NewMD5/NewSHA1 benötigt,
// da diese Funktionen das intern übernehmen, aber nützlich zum Verständnis der Namens-UUIDs)
func hashData(data []byte, hashType string) string {
	switch hashType {
	case "md5":
		return fmt.Sprintf("%x", md5.Sum(data))
	case "sha1":
		return fmt.Sprintf("%x", sha1.Sum(data))
	default:
		return "Unsupported hash type"
	}
}
