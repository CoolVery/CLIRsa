package dberr

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/CoolVery/CLIRsa.git/internal/errors/apperr"
)

func MapGet(err error, notFound error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return notFound
	case errors.Is(err, context.Canceled): 
		return apperr.ErrContextCanceled
	case errors.Is(err, context.DeadlineExceeded): 
		return apperr.ErrContextDeadline
	}
	return fmt.Errorf("db get: %w", err)
}