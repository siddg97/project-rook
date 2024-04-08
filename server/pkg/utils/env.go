package utils

import (
	"fmt"
	"os"
)

type MissingEnvVarError struct {
	name string
}

func (m MissingEnvVarError) Error() string {
	return fmt.Sprintf("Missing environment variable %s", m.name)
}

func GetEnvVar(variableName string) (string, error) {
	value, ok := os.LookupEnv(variableName)
	if ok {
		return value, nil
	}

	return "", MissingEnvVarError{name: variableName}
}

func GetDefaultEnvVar(variableName string, defaultValue string) (string, error) {
	value, err := GetEnvVar(variableName)
	switch err.(type) {
	case MissingEnvVarError:
		return defaultValue, nil
	}

	if err != nil {
		return "", err
	}

	return value, nil
}
