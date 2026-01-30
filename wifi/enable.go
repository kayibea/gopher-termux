package wifi

import "github.com/kayibea/gopher-termux/termux"

func Enable(on bool) error {
	arg := "false"
	if on {
		arg = "true"
	}

	_, err := termux.Exec("termux-wifi-enable", arg)
	return err
}
