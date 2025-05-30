package pgo

import (
	"log/slog"

	"github.com/DataDog/datadog-pgo/internal"
)

func ExtractPGOFiles(req *ExtractPGOFilesRequest) (*ExtractPGOFilesResponse, error) {
	mergedProfile, err := internal.ExtractMergedProfile(
		req.From,
		req.Timeout,
		req.Profiles,
		false,
		req.ShouldUsePGOEndpoint,
		slog.Default(),
		req.Queries...,
	)

	if err != nil {
		return nil, err
	}

	return &ExtractPGOFilesResponse{
		ResponseProfile: mergedProfile.GetProfile(),
	}, err
}
