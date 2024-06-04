# imir

<img src="https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExcG4xMWdyNGx4cGN2dGhpbWVtaWU0cTM5MGNlZnp3MDJweXM0enBqMSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/S984YkgY62vvSUAwhv/giphy-downsized.gif" width="100%" height="auto">

> 🔮 _"**I**s **M**ercury **I**n **R**etrograde?"_

- Do you ever find yourself staring at your service provider's status page when troubleshooting outages? Look no further! Find out if the answer behind your malfunctioning technology lies within the movement of our moonless, littlest planet - Mercury!
- The direction of Mercury is now available directly from the comfort of your own CLI!
- Inform your astrology hating friends that you love to annoy!
- Shout out to the creator(s) behind the public [Mercury Retrograde API](https://mercuryretrogradeapi.com/about.html) for making this possible.

## 💾 Installation

📋 **Requirements:**

- [Go](https://go.dev/)

  - _See exact version inside `go.sum:3`._

Clone this repository and `cd` into its root directory.

```shell
# Compile the app.
$ go build
# On Linux or Mac:
$ ./imir
# On Windows:
$ imir.exe

# Install the created binary file.
$ go install

# Run the program either by running
# the main file (if tinkering with the repo locally etc)
$ go run main.go
# .. or by using the installed binary file.
$ imir
```

## 🚀 Usage

```shell
# Find out if Mercury is in retro grade or not.
$ imir

# Show help menu.
$ imir -h

# Learn more about Mercury in retrograde.
$ imir -a

# Is/was/will Mercury be in retrograde during a given date?
$ imir -d 2025-01-29
```

## Tests

📋 **Requirements:**

- [Fish Shell](https://fishshell.com/)

```fish
# From this project's root dir.
$ chmod +x ./scripts/run_tests.fish
# Run tests.
$ ./scripts/run_tests.fish
```
