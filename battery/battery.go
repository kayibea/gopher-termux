package battery

import (
	"encoding/json"

	"github.com/kayibea/gopher-termux/termux"
)

func Status() (*Info, error) {
	out, err := termux.Exec("termux-battery-status")
	if err != nil {
		return nil, err
	}

	var s Info
	if err := json.Unmarshal(out, &s); err != nil {
		return nil, err
	}

	return &s, nil
}
