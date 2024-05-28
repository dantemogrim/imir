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

1. Clone this repository and `cd` into its root directory.

```shell
#
$ go install
$ go run main.go
```

> Building from source? Do the following steps to get hold of the binary file.

## 🚀 Usage

```shell
# Find out if Mercury is in retro grade or not.
$ imir

# Show help menu.
$ imir -h

# Learn more about Mercury in retrograde.
$ imir -a

# See if Mercury is/was/will be in retrograde a specific date.
$ imir -d 2025-01-29
```

## Tests

📋 **Requirements:**

- [Fish Shell](https://fishshell.com/)

```fish
# From this project's root dir.
$ chmod +x ./scripts/run_tests.fish
$ ./scripts/run_tests.fish
```

### Backlog

- [ ] Create an error service?
- [ ] Replace flaggy with native flag library option. Only short flags are working now anyways.
- [ ] Finish `README.md`.
- [ ] Upload to homebrew.
