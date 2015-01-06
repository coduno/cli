# Coduno Command Line Interface

[![Build Status](https://drone.io/github.com/coduno/cli/status.png)](https://drone.io/github.com/coduno/cli/latest)

The Coduno CLI allows you to interact with the Coduno platform in order to compete in challenges.

## Usage

### `coduno login`

Logs you in and stores an entry in `~/.netrc` (Linux, Mac OS) or respectively `$HOME\_netrc` like so:

```
machine api.cod.uno
    login you@example.com
    password cafebabe
machine git.cod.uno
    login you@example.com
    password deadbeef
```

Each new login invalidates earlier logins.

## Distribution

### Linux

Download this repository and `go build` it or use our [fancy script](https://github.com/coduno/cli/blob/gh-pages/install.sh) to download a prebuilt
binary from Drone.io.

To install to `/usr/local/bin`, root privileges might be needed:
```
curl -s https://coduno.github.io/cli/install.sh | sudo bash
```

You can also tell the installer where to put the binary via `$LOCATION`:
```
LOCATION=~/bin curl -s https://coduno.github.io/cli/install.sh | bash
```

### Mac OS

We offer installation of a precompiled binary via a custom tap on Homebrew:

```
brew tap coduno/coduno  # This pulls our formula from GitHub
brew install coduno     # Actually fetches the binary from drone.io
```
