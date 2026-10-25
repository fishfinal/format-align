# format-align

<div align="center">

<img src="https://img.shields.io/badge/Go-%3E%3D1.16-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version" />
<img src="https://img.shields.io/badge/license-Apache--2.0-blue?style=for-the-badge" alt="License" />
<img src="https://img.shields.io/badge/build-passing-brightgreen?style=for-the-badge&logo=githubactions&logoColor=white" alt="Build Status" />
<img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=for-the-badge" alt="PRs Welcome" />

</div>

A simple and effective Go string formatting and alignment tool. It helps you easily output key-value pairs with perfect left or right alignment in the terminal, making your logs or CLI outputs much cleaner and more professional.

## :rocket: Installation

```shell
go get -u github.com/channellabs/format-align
```

## :rainbow: Features

- **Alignment**: Supports both Left and Right alignment.
- **Placeholder**: Define the exact width of the field name column.
- **Separator**: Customize the character used between the field name and the value (default is `:`).

## :ship: Usage Examples

### 1. Use Default Arguments

By default, it uses `Left` alignment, a placeholder width of `26`, and `:` as the separator.

```go
package main

import (
  "bytes"
  "fmt"

  formatalign "github.com/channellabs/format-align"
)

func main() {
  align := formatalign.NewFormatAlign()
  buffer := bytes.Buffer{}

  buffer.Write(append([]byte(align.Format("Username", "fishfinal")), '\n'))
  buffer.Write(append([]byte(align.Format("Sex", "male")), '\n'))
  buffer.Write(append([]byte(align.Format("IsStudent", false)), '\n'))

  fmt.Println(buffer.String())
}
```

**Output:**
```text
Username                  : fishfinal
Sex                       : male
IsStudent                 : false
```

### 2. Alignment of Right

You can use the `WithAlignment` option to align the text to the right.

```go
package main

import (
  "bytes"
  "fmt"

  formatalign "github.com/channellabs/format-align"
)

func main() {
  align := formatalign.NewFormatAlign(
    formatalign.WithAlignment(formatalign.AlignmentOfRight),
  )
  buffer := bytes.Buffer{}

  buffer.Write(append([]byte(align.Format("Username", "fishfinal")), '\n'))
  buffer.Write(append([]byte(align.Format("Sex", "male")), '\n'))
  buffer.Write(append([]byte(align.Format("IsStudent", false)), '\n'))

  fmt.Println(buffer.String())
}
```

**Output:**
```text
                  Username: fishfinal
                       Sex: male
                 IsStudent: false
```

### 3. Specify Placeholder Width

You can adjust the width of the field name column using `WithPlaceholder`.

```go
package main

import (
  "bytes"
  "fmt"

  formatalign "github.com/channellabs/format-align"
)

func main() {
  align := formatalign.NewFormatAlign(
    formatalign.WithAlignment(formatalign.AlignmentOfRight),
    formatalign.WithPlaceholder(12),
  )
  buffer := bytes.Buffer{}

  buffer.Write(append([]byte(align.Format("Username", "fishfinal")), '\n'))
  buffer.Write(append([]byte(align.Format("Sex", "male")), '\n'))
  buffer.Write(append([]byte(align.Format("IsStudent", false)), '\n'))

  fmt.Println(buffer.String())
}
```

**Output:**
```text
    Username: fishfinal
         Sex: male
   IsStudent: false
```

### 4. Specify Separator

You can change the separator string (or remove it entirely) using `WithSeparator`.

```go
package main

import (
  "bytes"
  "fmt"

  formatalign "github.com/channellabs/format-align"
)

func main() {
  align := formatalign.NewFormatAlign(
    formatalign.WithAlignment(formatalign.AlignmentOfRight),
    formatalign.WithPlaceholder(12),
    formatalign.WithSeparator("-"),
  )
  buffer := bytes.Buffer{}

  buffer.Write(append([]byte(align.Format("Username", "fishfinal")), '\n'))
  buffer.Write(append([]byte(align.Format("Sex", "male")), '\n'))
  buffer.Write(append([]byte(align.Format("IsStudent", false)), '\n'))

  fmt.Println(buffer.String())
}
```

**Output:**
```text
    Username- fishfinal
         Sex- male
   IsStudent- false
```

## :gear: API Overview

### Options

- `WithAlignment(align Alignment)`: Sets the alignment direction (`AlignmentOfLeft` or `AlignmentOfRight`).
- `WithPlaceholder(length int)`: Sets the width of the field name.
- `WithSeparator(sep string)`: Sets the separator between the field name and the value.

### Methods

- `NewFormatAlign(options ...Option) *FormatAlign`: Creates a new `FormatAlign` instance.
- `Format(name string, value interface{}) string`: Returns the formatted aligned string.
- `SetSeparator(sep string) *FormatAlign`: Chainable method to set the separator.
- `SetPlaceholder(length int) *FormatAlign`: Chainable method to set the placeholder width.

## :memo: License

Licensed under the Apache License, Version 2.0. See [LICENSE](LICENSE) for details.
