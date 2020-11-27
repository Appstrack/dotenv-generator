package main

import (
	"fmt"
	"os"
)

func UpdateEnvironments(list Environments) error {
	f, err := os.Create("./.env")
	if err != nil {
		return err
	}
	for _, item := range list {
		fmt.Fprintf(f, "%s=\"%s\"\n", item.Name, item.Value)
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
