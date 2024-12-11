package validator

import (
	"context"

	"github.com/go-playground/validator/v10"
)

func Validate(ctx context.Context, dto any) error {
	validate := validator.New()

	return validate.StructCtx(ctx, dto)
}
