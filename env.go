package env

import (
	"os"
	"strings"
)

type ev struct {
	Key  string
	Name string
}

type App struct {
	envs []ev
}

func (self *App) Var(name string, key string) {
	if self.envs == nil {
		self.envs = make([]ev, 0)
	}
	if len(name) > 0 && len(key) > 0 {
		self.envs = append(self.envs, ev{Name: name, Key: key})
	}
}

func (self *App) Parse() map[string]interface{} {
	envvars := os.Environ()
	vars := make(map[string]interface{})
	for _, e := range self.envs {
		for _, f := range envvars {
			if strings.HasPrefix(f, e.Key) {
				vars[e.Name] = strings.Trim(f, e.Key+"=")
			}
		}
	}
	return vars
}
