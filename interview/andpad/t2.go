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
	resCh := make(chan *Profile, 1)
	errCh := make(chan error, 1)

	wg.Add(2)
	go func() {
		s.notiClient.Notify(ctx, targetUserID, viewerID)
		wg.Done()
	}()
	go func(resCh chan *Profile, errCh chan error) {
		user, err := s.repo.Get(ctx, targetUserID)
		if err != nil {
			errCh <- err
			return
		}
		profile := &Profile{
			ID:   user.ID,
			Name: user.Name,
		}
		resCh <- profile
		wg.Done()
	}(resCh, errCh)
	wg.Wait()
	return <-resCh, <-errCh
}

func main() {

}
