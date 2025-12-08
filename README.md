# Icebar

An docker application developed for bspwm and i3, made to be super simple and static. Running with a GO backend

## TODO

**Features**

- [ ]  Auto Hide
- [X]  Margin from config
- [ ]  Box style (box/island/none)
- [ ]  Dynamic configuration handling
- [X]  Set monitor from config (must be handled in wm config)
- [ ]  Running-state integration
- [ ]  Dynamic config path on launch
- [ ]  Paper basket option

## File structure

```text
icebar/
├── cmd/
│   └── icebar/
│       └── main.go
├── pkg/
│   ├── config/
│   │   ├── config.go
│   │   └── parser.go
│   ├── dock/
│   │   ├── dock.go
│   │   └── item.go
│   ├── gui/
│   │   ├── gui.go
│   │   └── events.go
│   ├── launcher/
│   │   └── launcher.go
│   └── utils/
│       └── helpers.go
├── assets/
│   └── icebar/
│   │   ├── apps.toml
│   │   └── config.toml
├── go.mod
├── go.sum
└── README.md

```
