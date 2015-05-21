package env

import (
	"os"
	"strings"
)

type Env struct {
	Key  string
	Name string
}

type App struct {
	Envs []Env
}

func (self *App) Var(name string, key string) {
	if self.Envs == nil {
		self.Envs = make([]Env, 0)
	}
	e := Env{Name: name, Key: key}
	self.Envs = append(self.Envs, e)
}

func (self *App) Parse() map[string]interface{} {
	envvars := os.Environ()
	vars := make(map[string]interface{})
	for _, e := range self.Envs {
		for _, f := range envvars {
			if strings.HasPrefix(f, e.Key) {
				vars[e.Name] = strings.Trim(f, e.Key+"=")
			}
		}
	}
	return vars
}
