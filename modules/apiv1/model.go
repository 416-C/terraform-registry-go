package apiv1

import (
	"time"

	"github.com/416-C/terraform-registry-go/pkg"
)

type Provider struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Source    string `json:"source"`
	Version   string `json:"version"`
}

type Resource struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Input struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Default     string `json:"default"`
	Required    bool   `json:"required"`
}

type Output struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Root struct {
	Path                 string        `json:"path"`
	Name                 string        `json:"name"`
	Readme               string        `json:"readme"`
	Empty                bool          `json:"empty"`
	Inputs               []Input       `json:"inputs"`
	Outputs              []Output      `json:"outputs"`
	Dependencies         []interface{} `json:"dependencies"`
	ProviderDependencies []Provider    `json:"provider_dependencies"`
	Resources            []Resource    `json:"resources"`
}

type Module struct {
	ID          string    `json:"id"`
	Owner       string    `json:"owner"`
	Namespace   string    `json:"namespace"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Provider    string    `json:"provider"`
	Description string    `json:"description"`
	Source      string    `json:"source"`
	Tag         string    `json:"tag"`
	PublishedAt time.Time `json:"published_at"`
	Downloads   int       `json:"downloads"`
	Verified    bool      `json:"verified"`
	Root        Root      `json:"root"`
	Submodules  []Root    `json:"submodules"`
	Providers   []string  `json:"providers"`
	Versions    []string  `json:"versions"`
}

type ModuleList struct {
	Meta    pkg.Meta `json:"meta"`
	Modules []Module `json:"modules"`
}

type Version struct {
	Version string `json:"version"`
	Root    struct {
		Providers    []Provider    `json:"providers"`
		Dependencies []interface{} `json:"dependencies"`
	} `json:"root"`
	Submodules []struct {
		Path         string        `json:"path"`
		Providers    []Provider    `json:"providers"`
		Dependencies []interface{} `json:"dependencies"`
	} `json:"submodules"`
}

type VersionList struct {
	Modules []struct {
		Source   string    `json:"source"`
		Versions []Version `json:"versions"`
	} `json:"modules"`
}
