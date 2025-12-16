# Rush01-Go

A Go implementation of the classic "Skyscraper" (Rush01) puzzle solver, inspired by exercises from 42.

## Overview

This project demonstrates solving the Rush01 (Skyscraper) puzzle using Go. The puzzle consists of filling a 4x4 grid with numbers 1-4 so that:

- Each row and column contains each number exactly once.
- Clues around the grid indicate how many skyscrapers are visible from that direction.

## Features

- Command-line interface for inputting clues.
- Step-by-step board printing for debugging and learning.
- Modular code structure for clarity and maintainability.

## Usage

### Build and Run

From the `ex00` directory, run:

```sh
go run . <clues>
```

Or, to specify files explicitly:

```sh
go run rush01.go solver.go <clues>
```

### Input Format

Clues must be provided as a **single 16-digit string** in this order:

1. Top clues (left to right) — 4 digits  
2. Bottom clues (left to right) — 4 digits  
3. Left clues (top to bottom) — 4 digits  
4. Right clues (top to bottom) — 4 digits  

**Example:**

```sh
go run . 4132234214321234
```

### Example Output

```
     4   1   3   2 
   +---+---+---+---+
 2 | 1 | 2 | 3 | 4 | 3
   +---+---+---+---+
 2 | 4 | 3 | 2 | 1 | 2
   +---+---+---+---+
 3 | 2 | 1 | 4 | 3 | 4
   +---+---+---+---+
 1 | 3 | 4 | 1 | 2 | 2
   +---+---+---+---+
     4   3   2   1 
```

## Learning Goals

- Practice Go syntax and idioms.
- Implement backtracking algorithms.
- Work with slices, functions, and modular code.

## License

MIT License
