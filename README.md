# drexler

### A Bubble Tea TUI project generator

![drexler](https://user-images.githubusercontent.com/11764379/222506345-d309c288-ab9a-47f9-bfa3-317446b125bf.gif)

#### Getting Started

```
drexler init [appName]
cd [appName]
make run
```

#### Supported Commands

```
drexler help
drexler init [appName]
```

#### Generated File Structure

```
[appName]/
├─ cmd/
│  ├─ [appName]/
│  │  ├─ main.go
├─ internal/
│  ├─ services/
│  ├─ tui/
│  │  ├─ cmd.go
│  │  ├─ model.go
│  │  ├─ msg.go
│  │  ├─ view.go
├─ pkg/
├─ .gitignore
├─ go.mod
├─ go.sum
├─ Makefile
├─ README.md
```

#### WIP

```
drexler bubble [bubbleName]
drexler cmd [cmdName]
drexler msg [msgName]
drexler publish [homebrew, etc]
```
