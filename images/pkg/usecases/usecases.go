package usecases

import (
	"context"
)

type repository interface {
	ImgByQueryFetcher(ctx context.Context, query string) (string, error)
}

type myTaksaUsecases struct {
	repo repository
}

const TAKSA_CAPTION = "Собака умная может и самоутилизироваться )\n😍😍😍😍"

func New(
	repo repository) *myTaksaUsecases {
	return &myTaksaUsecases{
		repo: repo,
	}
}

func (u myTaksaUsecases) RandomTaksa(ctx context.Context) (string, error) {
	url, err := u.repo.ImgByQueryFetcher(ctx, "dachshund")

	if err != nil {
		return "", err
	}

	return url, err
}
