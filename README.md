# y2j - a YAML to JSON converter

Why?
Because writing nested YAML by hand during development and testing is
easier than JSON.

## Install

### macOS

```shell
brew install hbagdi/tap/y2j
```

### Linux

Navigate to the [release page](https://github.com/hbagdi/y2j/releases/tag/v1.0.0)
and install using a Debian or RPM package.
Alternatively, you can take the `.tar.gz` file and install the binary in your $PATH.

## Usage

```shell
echo "foo: bar" | y2j | http httpbin.org/post
```

