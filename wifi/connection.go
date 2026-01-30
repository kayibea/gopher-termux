package wifi

import (
	"encoding/json"

	"github.com/kayibea/gopher-termux/termux"
)

func Connection() (*ConnectionInfo, error) {
	out, err := termux.Exec("termux-wifi-connectioninfo")
	if err != nil {
		return nil, err
	}

	var info ConnectionInfo
	if err := json.Unmarshal(out, &info); err != nil {
		return nil, err
	}

	return &info, nil
}
