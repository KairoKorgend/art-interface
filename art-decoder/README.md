# Art-decoder Documentation

## Overview

This command-line tool is designed for creating text-based art by encoding and decoding character patterns. It supports both single-line and multi-line operations.

## How To Use

Run the tool with the following command:

```bash
go run main.go [flags] [string]
```

Replace [flags] with the desired options and [string] with your input.

## Features

### Decoder Mode

Decodes condensed character patterns.

#### Usage

Example: [5 #] expands to #####.

```bash
go run main.go "ABC[10 D]EFG"
Output: ABCDDDDDDDDDDEFG
```

### Encode Mode

Condenses repetitive characters into patterns. Activated with the -e flag.

#### Usage

Example: #####-_-_-_-_-_-##### encodes down to [5 #]5 -_]-[5 #].

```bash
go run main.go -e "#####-_-_-_-_-_-#####"
Output: [5 #][5 -_]-[5 #]
```

### Multi-Line decoder/encoder

Activated with the -m flag.
When using the -m flag without the -e flag, multi-line decoding mode gets activated.
When using both the -m and -e flags together, multi-line encoding mode gets activated.

Processes input line by line.

## Error Handling

Errors are handled for:

- Non-numeric first arguments in brackets.
- Missing space between arguments in brackets.
- Empty second argument in brackets.
- Unbalanced brackets.
