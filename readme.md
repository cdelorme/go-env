
# go-env

A library to abstract registration of named environment variables to load into shared application configuration.


## sales pitch

My library aims to deliver the simplest possible implementation of named environment variable registration.

The `Var()` method handles registration which accepts the name you will refer to it as, followed by the environment variable name to look for.

When you run `Parse()` it returns a `map[string]interface{}` which can be manipulated or merged with other forms of configuration gather.  A [map library](https://github.com/cdelorme/go-maps) is available aid with parsing and merging.

It is also fully-tested.

What it does not provide:

- abstractions
- dependency on other non-core libraries
- more than 40 lines of code

This package was built for use in combination with [go-config](https://github.com/cdelorme/go-config) and [go-option](https://github.com/cdelorme/go-option) to produce a single `map[string]interface{}` of application settings.


## usage

Import my library:

	import "github.com/cdelorme/go-env"

Create a new instance to define what your application is looking for:

	appEnvs := env.App{}

Define what you are looking for with an application-local name consistent with what you might load from a configuration file or flag/option:

	appEnvs.Var("name", "APP_LOCAL_NAME")

Once you have registered the environment variables you are looking for you can get a `map[string]interface{}` of the results:

	envs := appEnvs.Parse()

