package internal

import "github.com/IamNotUrKitty/gophermart/internal/config"

func Run() error {

	config := config.GetConfig()

	server, err := NewServer(config)
	if err != nil {
		return err
	}

	if err := server.Start(config.Address); err != nil {
		return err
	}

	return nil
}
