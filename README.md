## Booty

### build
```shell
make all
```

### working example
```shell
./booty --help
./booty -f templates/democfg.yaml -v debug
```

### Help menu
```shell
Bootstrap Docker configuration files from yaml templates.

Usage:
  booty [flags]
  booty [command]

Available Commands:
  help        Help about any command
  version     Application Version

Flags:
      --config string      config file (default is $HOME/.bootstrap.yaml)
  -f, --file string        Yaml input (default "bootstrap.yaml")
  -h, --help               help for booty
  -v, --verbosity string   Log level (debug, info, warn, error, fatal, panic (default "warning")

Use "booty [command] --help" for more information about a command.
```
