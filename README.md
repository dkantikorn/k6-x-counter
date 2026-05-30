# k6 Custom Build with Counter Extension

This project demonstrates how to build a custom `k6` binary with a custom extension using `xk6`.

---

## 📦 Prerequisites

- Go (version 1.22 or higher recommended)
- Git
- `k6` (optional, for comparison)

---

## ⚙️ Installation

### 1. Install `xk6`

```bash
go install go.k6.io/xk6/cmd/xk6@latest
```

Make sure `$GOPATH/bin` is added to your `PATH`.

---

## 🏗️ Build Custom k6 Binary

Navigate to your project directory:

```bash
mkdir ~/k6-custom && cd ~/k6-custom
```

Build `k6` with the custom extension:

## Build in local reposotory
```bash
xk6 build v1.7.1 \
  --with github.com/dkantikorn/k6-x-counter=$(pwd)/k6-x-counter \
  --output ./k6-custom
```

## Build in local reposotory
```bash
xk6 build v1.7.1 \
  --with github.com/dkantikorn/k6-x-counter \
  --output ./k6-custom
```


### 📌 Notes

- `--with` specifies the extension module and its local path
- `$(pwd)/k6-x-counter` ensures an absolute path is used
- Output binary will be named `k6-custom`

---

## 🚀 Run Load Test

Use the custom binary instead of the default `k6`:

```bash
./k6-custom run script.js
```

---

## 🧩 Example Usage (script.js)

```javascript
import { nextId, value, reset } from "k6/x/counter";

export function setup() {
  reset(); // Start counting from zero for each test run.
}

export default function () {
  for (let i = 0; i < 3; i++) {
    const id = nextId(6); // "000001", "000002", ...
    console.log(`VU ${__VU} got ID: ${id}`);
  }
}

export function teardown() {
  console.log(`Total IDs issued: ${value()}`);
}
```

---

## 🧩 Extension Info

- Module: `github.com/dkantikorn/k6-x-counter`
- Purpose: Provide a global counter for k6 test execution
- Behavior: Counter resets at the beginning of each test run

---

## 🔄 CI/CD (GitHub Actions)

This workflow builds the custom `k6` binary automatically on every push.

Create file: `.github/workflows/build-k6.yml`

```yaml
name: Build k6 Custom Binary

on:
  push:
    branches: [ main ]
  pull_request:

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout source
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Install xk6
        run: |
          go install go.k6.io/xk6/cmd/xk6@latest

      - name: Build custom k6 binary
        run: |
          xk6 build v1.7.1 \
            --with github.com/dkantikorn/k6-x-counter=./k6-x-counter \
            --output k6-custom

      - name: Upload artifact
        uses: actions/upload-artifact@v4
        with:
          name: k6-custom
          path: k6-custom
```

---

## 🛠️ Troubleshooting

### ❌ Build fails with dependency errors

```bash
go mod tidy
go clean -modcache
```

---

### ❌ Extension not found

Check the directory:

```bash
ls ./k6-x-counter
```

Verify `go.mod`:

```go
module github.com/dkantikorn/k6-x-counter
```

---

## 📄 License

This project is for internal/testing