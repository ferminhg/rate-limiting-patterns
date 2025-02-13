package domain

type Metrics interface {
	TrackRequest(path string, status int)
	TrackError(path string, status int)
}
