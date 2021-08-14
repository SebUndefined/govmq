package hook

import (
	"encoding/json"
)

type BaseWebHook struct {
	MountPoint string `json:"mountpoint,omitempty"`
	ClientId   string `json:"client_id,omitempty"`
}

func (wh *BaseWebHook) ToJson() (string, error) {
	b, err := json.Marshal(wh)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
