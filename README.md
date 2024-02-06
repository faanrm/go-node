# Go-Node

Go-Node is a command-line interface (CLI) tool developed in Go for quickly generating Node.js projects.

## Features

- Generate Node.js projects effortlessly
- Customize project directories and dependencies
- Supports both JavaScript and TypeScript projects

## Installation

To install Go-Node, you need to have Go installed. Then, run the following command:

```bash
go get github.com/faanrm/go-node

## Usage 

Generate a Node.js Project : 
go-node node-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"

Generate a Node.js Template :  

go-node  template my-template-folder

This command generates a basic Node.js project template with TypeScript.

This command generates a new Node.js project with the specified dependencies.

Generate a TypeScript Node.js Project

go-node ts-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"

This command generates a new Node project using TypeScript.

Generate a Node.js Project Template

go-node template-ts my-template-folder
This command generates a basic Node project template with TypeScript.