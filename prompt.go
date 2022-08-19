package main

import (
	"github.com/manifoldco/promptui"
)

func Prompt(label string) error {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}

	_, err := prompt.Run()

	if err != nil {
		return err
	}
	return nil
}
