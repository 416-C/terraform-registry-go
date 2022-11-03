package apiv1

import (
	"time"

	"github.com/416-C/terraform-registry-go/pkg"
)

type ModuleList struct {
	Meta    pkg.Meta `json:"meta"`
	Modules []struct {
		ID          string    `json:"id"`
		Owner       string    `json:"owner"`
		Namespace   string    `json:"namespace"`
		Name        string    `json:"name"`
		Version     string    `json:"version"`
		Provider    string    `json:"provider"`
		Description string    `json:"description"`
		Source      string    `json:"source"`
		PublishedAt time.Time `json:"published_at"`
		Downloads   int       `json:"downloads"`
		Verified    bool      `json:"verified"`
	} `json:"modules"`
}

type AvailableVersionList struct {
	Modules []AvailableVersion `json:"modules"`
}

type AvailableVersion struct {
	Source   string `json:"source"`
	Versions []struct {
		Version string `json:"version"`
		//Submodules []struct {
		//	Path      string `json:"path"`
		//	Providers []struct {
		//		Name    string `json:"name"`
		//		Version string `json:"version"`
		//	} `json:"providers"`
		//	Dependencies []interface{} `json:"dependencies"`
		//} `json:"submodules"`
		Root struct {
			Dependencies []interface{} `json:"dependencies"`
			Providers    []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"providers"`
		} `json:"root"`
	} `json:"versions"`
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
	PublishedAt time.Time `json:"published_at"`
	Downloads   int       `json:"downloads"`
	Verified    bool      `json:"verified"`
	Root        struct {
		Path string `json:"path"`
		//Readme string `json:"readme"`
		Empty  bool `json:"empty"`
		Inputs []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Default     string `json:"default"`
		} `json:"inputs"`
		Outputs []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"outputs"`
		Dependencies []interface{} `json:"dependencies"`
		Resources    []interface{} `json:"resources"`
	} `json:"root"`
	//Submodules []struct {
	//	Path   string `json:"path"`
	//	Readme string `json:"readme"`
	//	Empty  bool   `json:"empty"`
	//	Inputs []struct {
	//		Name        string `json:"name"`
	//		Description string `json:"description"`
	//		Default     string `json:"default"`
	//	} `json:"inputs"`
	//	Outputs []struct {
	//		Name        string `json:"name"`
	//		Description string `json:"description"`
	//	} `json:"outputs"`
	//	Dependencies []interface{} `json:"dependencies"`
	//	Resources    []struct {
	//		Name string `json:"name"`
	//		Type string `json:"type"`
	//	} `json:"resources"`
	//} `json:"submodules"`
	Providers []string `json:"providers"`
	Versions  []string `json:"versions"`
}
