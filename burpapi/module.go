package burpapi

const NAMED_CONFIGURATION = "NamedConfiguration"

type ScanTask struct {
	URLs               []string            `json:"urls"`
	ScanConfigurations []ScanConfiguration `json:"scan_configurations"`
}

type ScanConfiguration struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// {"scan_configurations":[{"name":"rickshang-OOB","type":"NamedConfiguration"}],"urls":[]}'
