package util

import (
	"context"
	"os"

	"google.golang.org/api/idtoken"

	"github.com/taxio/errors"
)

func VerifyGoogleIDToken(token string) error {
	var clientID = os.Getenv("GOOGLE_CLIENT_ID")
	ctx := context.Background()

	if _, err := idtoken.Validate(ctx, token, clientID); err != nil {
		return errors.Wrap(err)
	}

	if _, err := idtoken.Validate(ctx, token, clientID); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
