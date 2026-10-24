# a2bubbletea

A2UI v1.0 renderer for the terminal, built on Bubble Tea.

This repository is an early scaffold. It reads a JSONL stream of A2UI messages
and draws a static layout made of Text, Row, Column, Card and Divider
components. It will grow alongside the A2UI v1.0 specification.

```sh
go run ./cmd/a2bubbletea samples/static-layout.jsonl
```

Press `q` to quit.

## Development

```sh
make lint
make test
```

## Specification

The A2UI v1.0 schemas, test cases and basic catalog are a git submodule at
`specification/a2ui`, pinned at commit
`b817d0296010075ce230fdd720f00e32cd07b775` (23 October 2026). Fetch it with:

```sh
git submodule update --init specification/a2ui
```

See [specification/README.md](specification/README.md) for the paths inside it.

Licensed under MIT.
