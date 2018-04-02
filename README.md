# enconf

[![Build Status](https://travis-ci.org/suluvir/enconf.svg?branch=master)](https://travis-ci.org/suluvir/enconf) [![Go Report Card](https://goreportcard.com/badge/github.com/suluvir/enconf)](https://goreportcard.com/report/github.com/suluvir/enconf)

With `enconf`, you can read the configuration from the environment in a simple and 
type safe way. Just take a look at the following example:

```go
type config struct {
	Port    int
	Version string
}

func main() {
	var config = config{}

	enconf.LoadConfigurationWithPrefix("SULUVIR", &config)
	
	fmt.Printf("%d", config.Port) // `Port` will contain the casted value of `SULUVIR_PORT`
}
```
