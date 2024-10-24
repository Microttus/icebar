# Icebar 



An docker application, made to be super simple and static. Running with a GO backend


## TODO

**Features**
- [ ] Auto Hide
- [x] Margin from config
- [ ] Box style (box/island/none)
- [ ] Dynamic configuration handling
- [ ] Set monitor from config
- [ ] Running-state integration
- [ ] Dynamic config path on launch
- [ ] Paper basket option

**Bugs**
- [ ] Misalignment

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