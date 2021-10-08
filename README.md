# y2j - a YAML to JSON converter

Why?
Because writing nested YAML by hand during development and testing is
easier than JSON.

## Usage

```shell
echo "foo: bar" | y2j | http httpbin.org/post
```

