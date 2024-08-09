package internal

func Run() error {

	server, err := NewServer()
	if err != nil {
		return err
	}

	if err := server.Start("localhost:8080"); err != nil {
		return err
	}

	return nil
}
