package config

import "os"

func DB_HOST() string {
	return os.Getenv("DB_HOST")
}

func DB_USER() string {
	return os.Getenv("DB_USER")
}

func DB_PASSWORD() string {
	return os.Getenv("DB_PASSWORD")
}

func DB_NAME() string {
	return os.Getenv("DB_NAME")
}

func DB_PORT() string {
	return os.Getenv("DB_PORT")
}
