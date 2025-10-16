package dependencies

import "strings"

// Эмуляция подзависимостей
func EmulateSubDeps(dep string) []*DepNode {
	switch {
	case strings.HasPrefix(dep, "numpy"):
		return []*DepNode{{ Name: "mkl>=2024.0" }, { Name: "setuptools" }}
	case strings.HasPrefix(dep, "pandas"):
		return []*DepNode{{ Name: "numpy>=1.25.0" }}
	case strings.HasPrefix(dep, "requests"):
		return []*DepNode{{ Name: "urllib3" }, 
		{ Name: "charset-normalizer" }, { Name: "certifi" }}
	case strings.HasPrefix(dep, "flask"):
		return []*DepNode{{ Name: "Werkzeug" }, { Name: "Jinja2" }, 
		{ Name: "itsdangerous" }, { Name: "click" }}
	default:
		return nil
	}
}