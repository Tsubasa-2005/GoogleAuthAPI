package handler

import (
	"context"
	"fmt"

	"GoogleAuthAPI/pkg/infra/rdb"
	"GoogleAuthAPI/pkg/util"
	"GoogleAuthAPI/ui/api"

	"github.com/taxio/errors"
)

func (h *Handler) GoogleAuth(ctx context.Context, req *api.SignInRequest) (api.GoogleAuthRes, error) {
	err := util.VerifyGoogleIDToken(req.Token)
	fmt.Print(err)
	if err != nil {
		return &api.Unauthorized{
			Message: "Invalid Token",
		}, nil
	}

	user, err := h.repo.GetUserByEmail(ctx, string(req.Email))
	if err != nil {
		if string(err.Error()) != "no rows in result set" {
			return nil, errors.Wrap(err)
		} else {
			user, err = h.repo.CreateUser(ctx, rdb.CreateUserParams{
				Email:    string(req.Email),
				Name:     string(req.Name),
				ImageUrl: string(req.ImageURL),
			})
			if err != nil {
				return nil, errors.Wrap(err)
			}
		}
	} else {
		user, err = h.repo.UpdateUser(ctx, rdb.UpdateUserParams{
			Email:    string(req.Email),
			Name:     string(req.Name),
			ImageUrl: string(req.ImageURL),
			ID:       user.ID,
		})
		if err != nil {
			return nil, errors.Wrap(err)
		}
	}

	token, err := util.GenerateToken("DataToReportAPI", util.UserToken{
		ID:       user.ID,
		Name:     user.Name,
		ImageURL: user.ImageUrl,
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &api.GoogleAuthOK{
		SetCookie: token,
	}, nil
}
