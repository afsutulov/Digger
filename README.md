# Digger Game

Native Go port of the 1983 Digger game. See **README_RU.md** for full build,
controls, fidelity and validation notes. The game does not execute the DOS COM.

## Icon and Windows builds

`digger.png` is embedded by `//go:embed` in `main.go` and passed to
`ebiten.SetWindowIcon`. The executable needs no external icon file. Window icon
display depends on the desktop; Ebitengine does not support this API on macOS.
`digger.ico` contains 16, 24, 32, 48, 64, 128 and 256 pixel variants. The included
`icon_windows_amd64.syso` and `icon_windows_arm64.syso` embed it as a Windows PE
resource, so Explorer can display the executable icon. Go automatically selects
the resource for the target architecture.

From the project root on Windows (PowerShell or cmd):

```text
go build -mod=vendor -buildvcs=false -ldflags="-s -w -H=windowsgui" -o digger.exe .
```

To replace the icon, replace both `digger.png` (square PNG) and `digger.ico`
(multi-resolution ICO), regenerate the resources, then rebuild. Install the
pinned [rsrc tool](https://github.com/akavel/rsrc) once; this optional step needs
network access. Add Go's binary directory (normally `GOPATH/bin`) to PATH.

```text
go install github.com/akavel/rsrc@v0.10.2
rsrc -ico digger.ico -arch amd64 -o icon_windows_amd64.syso
rsrc -ico digger.ico -arch arm64 -o icon_windows_arm64.syso
go build -mod=vendor -buildvcs=false -ldflags="-s -w -H=windowsgui" -o digger.exe .
```

Keep the architecture suffixes and avoid adding a generic `rsrc.syso` alongside
these files, which would duplicate resources. ARM64 resource support in rsrc is
marked experimental; the included ARM64 resource was checked in the built PE.
Cross-build examples (sh, including Linux/macOS):

```sh
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -buildvcs=false -ldflags='-s -w -H=windowsgui' -o digger-x64.exe .
GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -mod=vendor -buildvcs=false -ldflags='-s -w -H=windowsgui' -o digger-arm64.exe .
```

Linux/macOS build commands above embed the PNG as well. A Finder icon on macOS
requires an `.app` bundle with an `.icns` file and `Info.plist`; a standalone
command-line executable does not provide that icon.

Arrow keys move, F1 fires, Space pauses, F7 toggles music, F9 toggles sound,
F10 returns to the title, Escape changes player count on the title,
Alt+Enter switches fullscreen. `--speed 80` sets milliseconds per game frame.

GPL-3.0-only; see LICENSE and THIRD_PARTY_NOTICES.md.
