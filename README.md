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

Licensed under MIT.
