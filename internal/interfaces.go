package internal

type Flags struct {
	Version string `json:"path"`
	Vendor  bool   `json:"vendor"`
	Modules bool   `json:"modules"`
}
