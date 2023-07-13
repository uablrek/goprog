# goprog

Base for go programs. Support for log, config and sub-programs

There are many projects to support these functions, but they are all
too complicated for my taste.

* **log** - Using [logr](github.com/go-logr/logr) and a [zap logger](go.uber.org/zap) as backend
* **config** - Using [go-flagsfiller](github.com/itzg/go-flagsfiller)
* **subcommands** - Own implementation

The program in [cmd/template/](cmd/template/main.go) is an example and
may be copied to a new project.

```
./build.sh static
_output/template
_output/template version
```

