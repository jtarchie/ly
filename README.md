# Introduction

Avoid YAML templating and use something to generate YAML, programmatically.
This uses [Lua](http://tylerneylon.com/a/learn-lua/) (the programming language) to generate YAML file manifests.
This is a self contain binary, no external dependencies needed.

# Getting Started

Lua is simple.
It supports strings, numbers, boolean, arrays, and maps.

1. Create a file called `manifest.lua`.
1. Give it the contents.

   ```lua
   return {name="Hello World"}
   ```
   
   The last expression of your `manifest.lua` file should always have a `return` statement of a hash or array (`table` type in Lua).
   You will receive in error if the type is incorrect.
   
   See the [examples](examples) for more advance usage.
   
1. Invoke either from [source](#source) or [releases](#releases).
1. The stdout will be the following YAML payload.

   ```yaml
   name: Hello World
   ```

## `null` support

Lua has one _quirk_.
If a `table` (array or hash) has a `nil` value entry it is no longer considered as contiguous entry in the `table`.
This results in the outputted YAML not having a `null` entry.

To resolve this a `null` value was created in `ly`'s lua runtime.

The following Lua:

```lua
return {name=null}
```

will render YAML as:

```yaml
name: null
```

# Installation

## Source

The source uses `go.mod` for dependency management.
Please ensure that `golang` version 0.12.x or higher is installed.

```bash
$ git clone https://github.com/jtarchie/ly
$ cd ly
$ go run ly/main.go -c examples/variable.lua
list_with_a_map:
- key1: value1
  key2: value2
- item2
```

## Releases

1. Download the [latest release](https://github.com/jtarchie/ly/releases/latest).
1. Extract the binary from the `.tar.gz` file -- ie `tar xzf ~/Downloads/ly_*.tar.gz`.
1. Put it into an executable `PATH` -- ie `chmod +x ~/Downloads/ly_* && mv ly_ /usr/local/bin/ly`.
1. Run `ly -c <your-lua-file>.lua`
