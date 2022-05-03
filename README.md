# Epick

![1.0.0](https://img.shields.io/badge/status-1.0.0-red)
[![GoDoc](https://godoc.org/github.com/ethanbaker/epick?status.svg)](https://godoc.org/github.com/ethanbaker/epick)
[![Go Report Card](https://goreportcard.com/badge/github.com/ethanbaker/epick)](https://goreportcard.com/report/github.com/ethanbaker/epick)
[![Coverage](https://gocover.io/_badge/github.com/ethanbaker/epick)](https://gocover.io/github.com/ethanbaker/epick)


An extensive emoji picker in the terminal.

---

## About

Epick was designed to reduce the hassle of having to search up the copy and paste for different emojis whenever you want an emoji in a text document. For instance, adding emojis to blog posts in the terminal can be so much easier when using Epick, as you no longer need to copy and paste a character after searching for the perfect emoji for 5 minutes. All you need to do is run the `epick` command, find the preferred emoji, and continue coding!

---

## Installation

#### Golang 

To include the Epick package, include the line `import "github.com/ethanbaker/epick"`.

To make a command based off of the package, run `go install` in the
`cmd/cpick` directory. Command usage can be found in the [docs](https://godoc.org/github.com/ethanbaker/epick).

---

## Documentation

[Documentation](https://godoc.org/github.com/ethanbaker/epick) is present using Godoc.

---

## Features

#### Controls/Keys

Epick contains multiple pages of different emojis, each of which is a different category for emojis. There are 9 different categories of emojis:

* Smileys and Emotion
* People and Body
* Animals and Nature
* Food and Drink
* Travel and Places
* Activities
* Objects
* Symbols
* Flags

These emojis can be selected by using the vim keys (h, j, k, l) or the arrow keys.

---

## Issues, Suggestions, and Patches

For issues and suggestions, please include as much useful information is possible. Make sure the issue is actually present or the suggestion is not included.

For patches, please submit them as pull requests.
