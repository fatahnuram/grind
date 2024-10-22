package repository

import "github.com/fatahnuram/grind/internal/repository/activity"

type Progress struct {
	UpdatedAt  string
	Activities []activity.Activity
}
