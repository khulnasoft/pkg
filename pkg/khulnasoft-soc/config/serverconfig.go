package config

import (
	"errors"
	"strings"

	"github.com/khulnasoft/khulnasoft-soc/module"
)

const DEFAULT_MAX_PACKET_COUNT = 5000
const DEFAULT_IDLE_CONNECTION_TIMEOUT_MS = 300000
const DEFAULT_MAX_UPLOAD_SIZE_BYTES = 26214400

type ServerConfig struct {
	AirgapEnabled           bool                   `json:"airgapEnabled"`
	BindAddress             string                 `json:"bindAddress"`
	BaseUrl                 string                 `json:"baseUrl"`
	DeveloperEnabled        bool                   `json:"developerEnabled"`
	HtmlDir                 string                 `json:"htmlDir"`
	MaxPacketCount          int                    `json:"maxPacketCount"`
	Modules                 module.ModuleConfigMap `json:"modules"`
	ModuleFailuresIgnored   bool                   `json:"moduleFailuresIgnored"`
	ClientParams            ClientParameters       `json:"client"`
	IdleConnectionTimeoutMs int                    `json:"idleConnectionTimeoutMs"`
	TimezoneScript          string                 `json:"timezoneScript"`
	MaxUploadSizeBytes      int                    `json:"maxUploadSizeBytes"`
}

func (config *ServerConfig) Verify() error {
	var err error
	if config.MaxPacketCount <= 0 {
		config.MaxPacketCount = DEFAULT_MAX_PACKET_COUNT
	}
	if config.BindAddress == "" {
		err = errors.New("Server.BindAddress configuration value is required")
	}
	if strings.TrimSpace(config.BaseUrl) == "" {
		config.BaseUrl = "/"
	}
	if config.BaseUrl[len(config.BaseUrl)-1] != '/' {
		config.BaseUrl = config.BaseUrl + "/"
	}
	if err == nil {
		err = config.ClientParams.Verify()
	}
	if config.IdleConnectionTimeoutMs <= 0 {
		config.IdleConnectionTimeoutMs = DEFAULT_IDLE_CONNECTION_TIMEOUT_MS
	}
	if len(config.TimezoneScript) == 0 {
		config.TimezoneScript = "/opt/sensoroni/scripts/timezones.sh"
	}
	if config.MaxUploadSizeBytes == 0 {
		config.MaxUploadSizeBytes = DEFAULT_MAX_UPLOAD_SIZE_BYTES
	}
	return err
}
