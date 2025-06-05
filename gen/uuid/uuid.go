package uuid

import (
	"crypto/md5"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

// Definiere die Kommandozeilen-Flags
var (
	versionFlag = flag.String("v", "4", "UUID version to generate: 1 (time-based), 3 (name-MD5), 4 (random), 5 (name-SHA1), 6 (time-based sortable), 7 (time-based sortable). Default is 4.")
	namespaceID = flag.String("n", "", "Namespace UUID for name-based UUIDs (version 3 or 5). Must be a valid UUID string.")
	nameData    = flag.String("N", "", "Name data for name-based UUIDs (version 3 or 5).")
	outputFile  = flag.String("o", "", "Output to file instead of stdout.")
	numUUIDs    = flag.Int("c", 1, "Number of UUIDs to generate.")
)

func UuidGenerate() {

	var writer io.Writer = os.Stdout
	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			log.Fatalf("Failed to create output file: %v", err)
		}
		defer f.Close()
		writer = f
	}

	for i := 0; i < *numUUIDs; i++ {
		var u uuid.UUID
		var err error

		switch *versionFlag {
		case "1":
			u, err = uuid.NewUUID()
		case "3":
			if *namespaceID == "" || *nameData == "" {
				log.Fatal("For version 3 UUIDs, -n (namespace) and -N (name) are required.")
			}
			ns, parseErr := uuid.Parse(*namespaceID)
			if parseErr != nil {
				log.Fatalf("Invalid namespace UUID: %v", parseErr)
			}
			u = uuid.NewMD5(ns, []byte(*nameData))
		case "4":
			//u = uuid.NewRandom()
		case "5":
			if *namespaceID == "" || *nameData == "" {
				log.Fatal("For version 5 UUIDs, -n (namespace) and -N (name) are required.")
			}
			ns, parseErr := uuid.Parse(*namespaceID)
			if parseErr != nil {
				log.Fatalf("Invalid namespace UUID: %v", parseErr)
			}
			u = uuid.NewSHA1(ns, []byte(*nameData))
		case "6":
			u, err = uuid.NewV6()
		case "7":
			u, err = uuid.NewV7()
		default:
			log.Fatalf("Unsupported UUID version: %s. Supported versions are 1, 3, 4, 5, 6, 7.", *versionFlag)
		}

		if err != nil {
			log.Fatalf("Failed to generate UUID: %v", err)
		}

		// Standardmäßig Kleinbuchstaben, wie bei den meisten uuidgen-Implementierungen
		// Für Großbuchstaben könnte man hier strings.ToUpper(u.String()) verwenden,
		// aber uuidgen gibt standardmäßig Kleinbuchstaben aus.
		fmt.Fprintln(writer, strings.ToLower(u.String()))
	}
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
