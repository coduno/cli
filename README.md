# Coduno Command Line Interface

[![Build Status](https://drone.io/github.com/coduno/cli/status.png)](https://drone.io/github.com/coduno/cli/latest)

The Coduno CLI allows you to interact with the Coduno platform in order to compete in challenges.

## Usage

### `coduno hello`

This is quite a stupid command that just checks whether it is able to reach the Coduno API.

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

Download this repository and `go build` it or use this fancy script to download a prebuilt
binary from Drone.io:

```
curl https://coduno.github.io/cli/install.sh | sh
```

### Mac OS

We offer installation of a precompiled binary via a custom tap on Homebrew:

```
brew tap coduno/coduno  # This pulls our formula from GitHub
brew install coduno     # Actually fetches the binary from drone.io
```
