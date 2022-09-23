# Minesweeper

Play minesweeper in the terminal!

![demo](./docs/demo.gif)

## Install

### Go Install

```
go install github.com/bhb603/minesweeper/cmd/minesweeper-cli@latest
```

### Docker

```
docker pull ghcr.io/bhb603/minesweeper
```

### Pre-Built Binary

Download the pre-compiled binaries from [the Releases page](https://github.com/bhb603/minesweeper/releases),
verify the artifacts with the checksum,
and copy the binary to the desired location.

For example:
```sh
VERSION=0.1.1
TARGET=Darwin_x86_64
wget "https://github.com/bhb603/minesweeper/releases/download/v${VERSION}/minesweeper_${VERSION}_${TARGET}.tar.gz"
wget "https://github.com/bhb603/minesweeper/releases/download/v${VERSION}/checksums.txt"
sha256sum --ignore-missing --check checksums.txt
tar xvf minesweeper_${VERSION}_${TARGET}.tar.gz
```

## Usage

```
minesweeper-cli
```

Docker:
```
docker run --rm -it ghcr.io/bhb603/minesweeper-cli
```
