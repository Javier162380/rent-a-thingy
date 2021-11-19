package form

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/manifoldco/promptui"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func argSelector(msg string, errorMessage string) (string, error) {

	promptValue := promptui.Prompt{
		Label: msg,
	}

	result, err := promptValue.Run()

	if err != nil {
		return "", errors.New(errorMessage)
	}

	return result, nil

}

func argSelectorWithOptions(msg string, options []string, errorMessage string, useoptions bool) (string, error) {

	if useoptions {
		promptList := promptui.Select{
			Label: msg,
			Items: options,
		}

		_, result, err := promptList.Run()

		if err != nil {
			return "", errors.New(errorMessage)
		}

		return result, nil

	} else {
		promptValue := promptui.Prompt{
			Label: msg,
		}

		result, err := promptValue.Run()

		if err != nil {
			return "", errors.New(errorMessage)
		}

		if contains(options, result) {
			return result, nil
		} else {
			return "", fmt.Errorf("option %s : %s it is not available", msg, result)
		}

	}

}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("sensible-browser", url)
	default:
		return fmt.Errorf("failed to open browser. platform %s not supported", runtime.GOOS)
	}

	return cmd.Start()
}
