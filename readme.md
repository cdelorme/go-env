
# go-env

A library to abstract importing environment variables similar to flags (eg. [go-option](https://github.com/cdelorme/go-option).


## sales pitch

My library aims to deliver a simple modular solution.  It provides a basic abstraction that matches the way my other libraries load application settings, allowing you to easily mix the approaches used to load configuration.

Features:

- search environment variables by `key`
- return values by application-matched `name`

It does not:

- provide a comprehensive testing suite (for its two functions)
- depend on other libraries
- provide wildly extensible abstractions
- have more than 50 lines of code


## usage

Import my library:

	import "github.com/cdelorme/go-env"

Create a new instance to define what your application is looking for:

	appEnvs := env.App{}

Define what you are looking for with an application-local name consistent with what you might load from a configuration file or flag/option:

	appEnvs.Var("name", "APP_LOCAL_NAME")

Once you have registered the environment variables you are looking for you can get a `map[string]interface{}` of the results:

	envs := appEnvs.Parse()

_You can then use [go-maps](https://github.com/cdelorme/go-maps), or any tool of your choice, to merge your application settings and extract values of the correct type for your needs._
