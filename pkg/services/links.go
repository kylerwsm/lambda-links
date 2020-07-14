package services

import (
	"strings"

	"github.com/kylerwsm/lambda-links/pkg/entity"
	"github.com/kylerwsm/lambda-links/pkg/repositories"
)

// ProcessShortLink preprocesses the input short link.
func ProcessShortLink(shortLink string) string {
	return strings.ToLower(shortLink)
}

// GetLink fetches the Link entry associated with the specified short link.
func GetLink(shortLink string) (entity.Link, error) {
	shortLink = ProcessShortLink(shortLink)
	link, err := repositories.GetLink(shortLink)
	return link, err
}
