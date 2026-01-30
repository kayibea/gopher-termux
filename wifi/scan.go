package wifi

import (
	"encoding/json"

	"github.com/kayibea/gopher-termux/termux"
)

func Scan() ([]ScanInfo, error) {
	out, err := termux.Exec("termux-wifi-scaninfo")
	if err != nil {
		return nil, err
	}

	var infos []ScanInfo
	if err := json.Unmarshal(out, &infos); err != nil {
		return nil, err
	}

	return infos, nil
}
