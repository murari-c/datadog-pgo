package pgo

import (
	"time"

	"github.com/google/pprof/profile"
)

type ExtractPGOFilesRequest struct {
	// Queries is a list of queries to run against the datadog profiles search.
	Queries []string

	// From represents how far back to search for profiles.
	From time.Duration

	// Timeout is the timeout for the pgo extraction.
	Timeout time.Duration

	// Profiles is the number of profiles to download per query.
	Profiles int

	// ShouldUsePGOEndpoint indicates whether to use the PGO endpoint.
	ShouldUsePGOEndpoint bool
}

type ExtractPGOFilesResponse struct {
	// ResponseProfile is the merged response profile data.
	// equivalent to the .pgo file.
	ResponseProfile *profile.Profile
}
