## Language

- [English](#english)
- [日本語](#日本語)

---

### English

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

## 3. SQL to Struct Generator
Converts SQL database tables into Go struct definitions, facilitating quicker integration between database structures and Go code.

### Features

- **Automatically generate Go structs from SQL table definitions**.
- **Supports custom types and tagging for JSON serialization**.

### Usage

To generate Go structs from a database table:

```bash
go run main.go sql struct --username username --password password --db database_name --table table_name
```

- **Flags**:
  - `--username` or `-u`: Username for database access.
  - `--password` or `-p`: Password for database access.
  - `--db` or `-d`: Database name.
  - `--table` or `-t`: Table name to convert to Go struct.

### Example Output
The tool will output Go structs to standard output, complete with JSON tags and comments derived from the database table structure. For example:

```go
type Users struct {
    // id
    Id int32 `json:"id"`
    // name
    Name string `json:"name"`
    // email
    Email string `json:"email"`
    // created_at
    CreatedAt time.Time `json:"created_at"`
}

func (model Users) TableName() string {
    return "users"
}
```

This output facilitates immediate use within Go applications, making database interactions straightforward and type-safe.

---

### 日本語

# Goツールコレクション

## 概要
Goの構文を実際に理解し学ぶことを目的とした、個人的なGoターミナルツールのコレクションです。

## 1. 命名フォーマッター
- **mode 1**: 単語を小文字に変換。
- **mode 2**: 単語を大文字に変換。
- **mode 3**: snake_caseをCamelCaseに変換。
- **mode 4**: snake_caseをcamelCaseに変換。
- **mode 5**: CamelCaseをsnake_caseに変換。

### Usage
命名フォーマッターコマンドを使用するには、変換する文字列と希望の変換モードを指定します:
```bash
go run main.go format --str "exampleName" --mode 5
```

## 2. 時間コマンド
日付と時間形式を処理および操作するためのユーティリティを提供します。

### コマンド

#### 2.1 `now` (現在の時間)
- **説明**: 現在のシステム時間を取得します。
- **Usage**:
  ```bash
  go run main.go time now
  ```
- **出力**: 現在の時間を `YYYY-MM-DD HH:MM:SS` 形式で表示します。

#### 2.2 `calc` (将来/過去の時間を計算)
- **説明**: 指定された開始時間からの指定された期間後の時間を計算します。
- **Usage**:
  ```bash
  go run main.go time calc --calculate "2022-01-01 12:00:00" --duration "1h30m"
  ```
- **flag**:
  - `--calculate` または `-c`: 計算の基準となる時間、例: `"2022-01-01 12:00:00"` またはUNIXタイムスタンプ。指定されていない場合は現在の時間がデフォルトとなります。
  - `--duration` または `-d`: 基準時間に追加する期間、例: `"1h"`, `"1m"`, `"1s"`。

## 3. SQLから構造体ジェネレーター
SQLデータベースのテーブルをGo構造体定義に変換し、データベース構造とGoコードの迅速な統合を促進します。

### 特徴

- **SQLテーブル定義からGo構造体を自動生成**。
- **JSONシリアル化のためのカスタムタイプおよびタグ付けをサポート**。

### Usage

データベーステーブルからGo構造体を生成するには：

```bash
go run main.go sql struct --username username --password password --db database_name --table table_name
```

- **flag**:
  - `--username` または `-u`: データベースアクセス用のユーザー名。
  - `--password` または `-p`: データベースアクセス用のパスワード。
  - `--db` または `-d`: データベース名。
  - `--table` または `-t`: Go構造体に変換するテーブル名。

### 出力例
ツールは、JSONタグおよびデータベーステーブル構造から派生したコメント付きのGo構造体を標準出力に出力します。例えば：

```go
type Users struct {
    // id
    Id int32 `json:"id"`
    // name
    Name string `json:"name"`
    // email
    Email string `json:"email"`
    // created_at
    CreatedAt time.Time `json:"created_at"`
}

func (model Users) TableName() string {
    return "users"
}
```

この出力は、Goアプリケーション内で即座に使用でき、データベースとの対話をシンプルで型安全なものにします。

Below, I've added a TODO List section with checkboxes to your README in both English and Japanese. This approach uses markdown checkboxes, which visually represent tasks but won't be interactive in a plain markdown environment. They serve as a visual guide for planned or completed tasks.

---
## TODO List
- [ ] **File Input and Batch Conversion**: Accept file inputs and enable batch processing of multiple entries.
- [ ] **User Interface (GUI)**: GUI to make it easier to use
- [ ] **Expand Time Commands**: Introduce more sophisticated time manipulation commands, such as timezone conversions and custom formatting options.
---

