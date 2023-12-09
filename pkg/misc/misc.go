package misc

import (
	"context"
	"fmt"

	"github.com/thompsonja/wmata-go/internal/helpers"
)

type API struct {
	requester *helpers.HttpRequester
}

func New(apiKey string) *API {
	return &API{
		requester: helpers.New(apiKey),
	}
}

func (a *API) Validate(ctx context.Context) error {
	url, err := helpers.GenerateUrl("Misc/Validate", nil)
	if err != nil {
		return fmt.Errorf("helpers.GenerateUrl: %v", err)
	}

	if _, err := a.requester.SendHttpRequest(ctx, url, nil); err != nil {
		return fmt.Errorf("a.requester.SendHttpRequest: %v", err)
	}
	return nil
}
