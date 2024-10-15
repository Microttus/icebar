# Icebar

An docker application, made to be super simple and static. Running with a GO backend


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
│   ├── app/
│   │   └── launcher.go
│   └── utils/
│       └── helpers.go
├── assets/
│   └── (icons, images, etc.)
├── go.mod
├── go.sum
└── README.md

```