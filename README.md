# Rush0X-Go

Rush0X-Go is a collection of Go projects for the Rush 01 and Rush 02 challenges. Each challenge is organized in its own directory with separate Go modules and source files.

## Project Structure

```
Rush0X-Go/
│
├── Rush01/
│   ├── README.md
│   └── ex00/
│       ├── go.mod
│       ├── rush01.go
│       └── solver.go
│
├── Rush02/
│   ├── README.md
│   └── ex00/
│       ├── Files.go
│       ├── go.mod
│       ├── rush02.go
│       └── lan/
│           ├── numbers.cat.dict
│           ├── numbers.du.dict
│           ├── numbers.en.dict
│           ├── numbers.es.dict
│           └── numbers.fr.dict
└── README.md
```

## Subprojects

### Rush01
- Contains the solution for the Rush 01 challenge.
- Main files: `rush01.go`, `solver.go`.
- To run or build, navigate to `Rush01/ex00` and use Go commands.

### Rush02
- Contains the solution for the Rush 02 challenge.
- Main files: `rush02.go`, `Files.go`.
- Includes language dictionaries in `lan/` for multi-language support.
- To run or build, navigate to `Rush02/ex00` and use Go commands.

## Getting Started

1. **Install Go** (if not already installed):
	https://golang.org/doc/install

2. **Clone this repository:**
	```bash
	git clone <repo-url>
	cd Rush0X-Go
	```

3. **Build and Run**
	- For Rush01:
	  ```bash
	  cd Rush01/ex00
	  go build
	  ./ex00  # or go run rush01.go
	  ```
	- For Rush02:
	  ```bash
	  cd Rush02/ex00
	  go build
	  ./ex00  # or go run rush02.go
	  ```

## Notes
- Each subproject is self-contained with its own `go.mod` file.
- See the respective `README.md` files in `Rush01/` and `Rush02/` for more details about each challenge.

---
*Created for Rush 01 and Rush 02 challenges in Go.*

