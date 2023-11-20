package usecases

import (
	"context"
)

type repository interface {
	ImgByQueryFetcher(ctx context.Context, query string) (string, error)
}

type myUsecases struct {
	repo repository
}

// const TAKSA_CAPTION = "Собака умная может и самоутилизироваться )\n😍😍😍😍"

func New(
	repo repository) *myUsecases {
	return &myUsecases{
		repo: repo,
	}
}

func (u myUsecases) RandomImg(ctx context.Context, query string) (string, error) {
	url, err := u.repo.ImgByQueryFetcher(ctx, query)

	if err != nil {
		return "", err
	}

	return url, err
}
