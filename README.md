
# Go Tools Collection

## Overview
A personal collection of Go terminal tools aimed at understanding and learning Go syntax through hands-on practice.

## 1. Naming Formatter (Naming Convention Converter)
- **Mode 1**: Convert words to lowercase.
- **Mode 2**: Convert words to uppercase.
- **Mode 3**: Convert snake_case to CamelCase.
- **Mode 4**: Convert snake_case to camelCase.
- **Mode 5**: Convert CamelCase to snake_case.

### Usage
To use the name formatter command, specify the string to convert and the desired conversion mode:
```bash
go run main.go format --str "exampleName" --mode 5
```

## 2. Time Commands
Provides utilities for processing and manipulating date and time formats.

### Commands

#### 2.1 `now` (Current Time)
- **Description**: Fetches the current system time.
- **Usage**:
  ```bash
  go run main.go time now
  ```
- **Output**: Displays the current time in the format `YYYY-MM-DD HH:MM:SS`.

#### 2.2 `calc` (Calculate Future/Previous Time)
- **Description**: Calculates the time after a specified duration from a given start time.
- **Usage**:
  ```bash
  go run main.go time calc --calculate "2022-01-01 12:00:00" --duration "1h30m"
  ```
- **Flags**:
  - `--calculate` or `-c`: The base time for calculation, e.g., `"2022-01-01 12:00:00"` or a UNIX timestamp. Defaults to the current time if not specified.
  - `--duration` or `-d`: The duration to add to the base time, e.g., `"1h"`, `"1m"`, `"1s"`.

