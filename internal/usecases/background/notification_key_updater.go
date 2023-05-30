package background

import (
	"context"
	"log"

	"websmee/buyspot/internal/usecases"
)

type NotificationKeyUpdater struct {
	notificationsSubscriber usecases.NotificationsSubscriber
	userRepository          usecases.UserRepository
	logger                  *log.Logger
}

func NewNotificationKeyUpdater(
	notificationsSubscriber usecases.NotificationsSubscriber,
	userRepository usecases.UserRepository,
	logger *log.Logger,
) *NotificationKeyUpdater {
	return &NotificationKeyUpdater{
		notificationsSubscriber,
		userRepository,
		logger,
	}
}

func (u *NotificationKeyUpdater) Run(ctx context.Context) {
	go func() {
		if err := u.notificationsSubscriber.Run(
			func(username, key string) error {
				user, err := u.userRepository.FindByTelegramUsername(ctx, username)
				if err != nil {
					u.logger.Println(err)
					return err
				}

				user.NotificationsKey = key
				if err := u.userRepository.CreateOrUpdate(ctx, user); err != nil {
					u.logger.Println(err)
					return err
				}

				return nil
			},
			func(err error) {
				u.logger.Println(err)
			},
		); err != nil {
			u.logger.Println(err)
		}
	}()
}
