package loggregator_v2

import (
	"time"

	"github.com/cloudfoundry/sonde-go/events"
)

type Client interface {
	SendDuration(name string, value time.Duration) error
	SendMebiBytes(name string, value int) error
	SendMetric(name string, value int) error
	SendBytesPerSecond(name string, value float64) error
	SendRequestsPerSecond(name string, value float64) error
	IncrementCounter(name string) error
	SendAppLog(appID, message, sourceType, sourceInstance string) error
	SendAppErrorLog(appID, message, sourceType, sourceInstance string) error
	SendAppMetrics(metrics *events.ContainerMetric) error
}

type MetronConfig struct {
	UseV2API      bool   `json:"loggregator_use_v2_api"`
	APIPort       int    `json:"loggregator_api_port"`
	CACertPath    string `json:"loggregator_ca_path"`
	CertPath      string `json:"loggregator_cert_path"`
	KeyPath       string `json:"loggregator_key_path"`
	JobDeployment string `json:"loggregator_job_deployment"`
	JobName       string `json:"loggregator_job_name"`
	JobIndex      string `json:"loggregator_job_index"`
	JobIP         string `json:"loggregator_job_ip"`
	JobOrigin     string `json:"loggregator_job_origin"`
	DropsondePort int    `json:"dropsonde_port"`
}

func NewClient(config MetronConfig) (Client, error) {
	if config.UseV2API {
		return newGrpcClient(config)
	}

	return newDropsondeClient()
}