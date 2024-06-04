# IMIR

<img src="./pkg/assets/imir.png" alt="IMIR" width="100%" height="auto">

> 🔮 _"**I**s **M**ercury **I**n **R**etrograde?"_

- Do you continuously find yourself staring at your service provider's status page when troubleshooting outages? Now - there's no need! Find the answers behind your malfunctioning technology right from your CLI! This zap powered app performs robust cyber requests regarding the movement of our moonless, littlest planet - Mercury!
- The direction of Mercury is now available directly from the comfort of your own CLI!
- By a single command stroke, you can annoy your astrology hating friends and coworkers!
- Level up your astrology game with this aggressively pseudoscientific tool!
- Shout out to the creator(s) behind the public [Mercury Retrograde API](https://mercuryretrogradeapi.com/about.html) for making this possible.

## 💾 Installation

📋 **Requirements:**

- [Go](https://go.dev/)

  - _See exact version inside `go.sum:3`._

Clone this repository and `cd` into its root directory.

```shell
# Compile the app.
$ go build
# On Linux/Mac:
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

## 🔬 Tests

📋 **Requirements:**

- [Fish Shell](https://fishshell.com/)

```fish
# From this project's root dir.
$ chmod +x ./scripts/run_tests.fish
# Run tests.
$ ./scripts/run_tests.fish
```
