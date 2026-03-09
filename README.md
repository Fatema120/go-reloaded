# Go Reloaded

A Go text-processing project that reads text from a `.txt` file, applies formatting/transformation rules, and writes the result to another `.txt` file.

## Features

- Applies text commands:
  - `(cap)` → capitalize previous/current word
  - `(low)` → lowercase previous/current word
  - `(up)` → uppercase previous/current word
  - `(bin)` → convert previous/current binary number to decimal
  - `(hex)` → convert previous/current hexadecimal number to decimal
- Supports multi-word commands:
  - `(cap, n)` → apply capitalize to previous `n` words
  - `(low, n)` → apply lowercase to previous `n` words
  - `(up, n)` → apply uppercase to previous `n` words
- Fixes punctuation spacing and attachment (`.,;:!?` and `'`)
- Corrects article usage (`a` / `an`) based on the following word

## Project Structure

- `text-editor.go` – main text processing pipeline (`TextEditor`)
- `utility-functions.go` – command handling, punctuation logic, and helpers
- `test/main.go` – CLI entry point
- `test/sample.txt` – sample input file
- `test/result.txt` – sample output file

## Requirements

- Go `1.23.4` (or compatible)

## How to Run

From the project root:

```bash
go run ./test/main.go ./test/sample.txt ./test/result.txt
```

Expected console output:

```text
Processing complete. Output saved to: ./test/result.txt
```

## CLI Usage

```bash
go run ./test/main.go <input_file.txt> <output_file.txt>
```

Rules:

- Exactly 2 arguments are required: input and output files
- Both files must use the `.txt` extension

## Example

Input (`test/sample.txt`):

```text
' it's a nice day '
```

Output (`test/result.txt`):

```text
'it's a nice day'
```

## Notes

- This project is organized as a module named `goreloaded`.
- The command parser supports both standalone commands and command tokens attached to words.
