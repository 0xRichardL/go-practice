package main

import (
	"context"
	"sync"
)

type (
	User struct {
		ID   string
		Name string
		IPN  string
	}

	Profile struct {
		ID   string
		Name string
	}

	NotiClient interface {
		Notify(ctx context.Context, targetUserID, viewerID string) error
	}

	UserRepo interface {
		Get(ctx context.Context, userID string) (*User, error)
	}

	Service struct {
		repo       UserRepo
		notiClient NotiClient
	}
)

func (s Service) LookupProfile(ctx context.Context, targetUserID, viewerID string) (*Profile, error) {
	// get, notif
	wg := sync.WaitGroup{}
	var profile *Profile
	// Errors are from 2 channels, buffer size 1 to avoid goroutine leak.
	// Return first error.
	errCh := make(chan error, 1)

	wg.Add(1)
	func() {
		defer wg.Done()
		if err := s.notiClient.Notify(ctx, targetUserID, viewerID); err != nil {
			errCh <- err
		}
	}()
	// Go ^1.25, use wg.Go func
	wg.Go(func() {
		user, err := s.repo.Get(ctx, targetUserID)
		if err != nil {
			errCh <- err
			return
		}
		profile = &Profile{
			ID:   user.ID,
			Name: user.Name,
		}
	},
	)
	wg.Wait()
	return profile, <-errCh
}
