# Go-Node

Go-Node is a command-line interface (CLI) tool developed in Go for quickly generating Node.js projects.

<img width="800" src="./home.gif">
## Features

- Generate Node.js projects effortlessly
- Customize project directories and dependencies
- Supports both JavaScript and TypeScript projects

## Installation

To install Go-Node, you need to have Go installed. Then, run the following command:

```bash
go get github.com/faanrm/go-node
```

## Usage

```bash
Generate a Node.js Project :
go-node node-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"
```

Generate a Node.js Template :

```bash
go-node  template my-template-folder
```

Generate a TypeScript Node Project

```bash
go-node ts-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"
```

Generate a Node TS Project Template

```bash
go-node template-ts my-template-folder
```

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## Authors

- [@Fanilo](https://www.github.com/faanrm)