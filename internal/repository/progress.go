package repository

import "github.com/fatahnuram/grind/internal/repository/activity"

const SupportedVersion = "1.0.0"

type Progress struct {
	Version    string
	UpdatedAt  string
	Activities []activity.Activity
}
