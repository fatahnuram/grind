package progress

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/fatahnuram/grind/internal/repository/activity"
)

const SupportedVersion = "1.0.0"

type Progress struct {
	Version    string
	UpdatedAt  string
	Activities []activity.Activity
}

func FromActivities(a []activity.Activity) Progress {
	return Progress{
		Version:    SupportedVersion,
		UpdatedAt:  time.Now().Format(time.RFC3339),
		Activities: a,
	}
}

func (p Progress) DumpJson(filepath string) {
	var file *os.File

	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(p)
}
