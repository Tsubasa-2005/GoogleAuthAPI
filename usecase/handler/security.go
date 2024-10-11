package handler

import (
	"context"
	"encoding/json"

	"GoogleAuthAPI/pkg/util"
	"GoogleAuthAPI/ui/api"

	"github.com/golang-jwt/jwt/v5"
	"github.com/taxio/errors"
)

func (s *SecurityHandler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	token, err := util.ParseToken(t.Token)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	claims := token.Claims.(jwt.MapClaims)

	id, err := getInt64FromClaims(claims, "user_id")
	if err != nil {
		return nil, errors.Wrap(err)
	}

	name, ok := claims["name"].(string)
	if !ok {
		return nil, errors.New("invalid name claim type")
	}

	image, ok := claims["image"].(string)
	if !ok {
		return nil, errors.New("invalid name claim type")
	}

	user := util.UserToken{
		ID:       id,
		Name:     name,
		ImageURL: image,
	}

	return context.WithValue(ctx, "user", user), nil
}

func getInt64FromClaims(claims jwt.MapClaims, key string) (int64, error) {
	value, ok := claims[key]
	if !ok {
		return 0, errors.New("claim not found")
	}

	switch v := value.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case json.Number:
		i, err := v.Int64()
		if err != nil {
			return 0, errors.Wrap(err)
		}
		return i, nil
	default:
		return 0, errors.New("invalid claim type")
	}
}
