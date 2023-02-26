package cli

import (
	"fmt"
	"github.com/gzipchrist/drexler/internal/genfile"
	"github.com/gzipchrist/drexler/internal/style"
	"github.com/gzipchrist/drexler/internal/tui"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Usage() {
	help := `
  Usage: drexler [command] [subcommand]

  Examples:
	drexler init myProject

  Commands: 
	help
	init [myProject]
	bubble [textinput, timer, etc]
	cmd [myCmd]
	msg [myMsg]
	deploy [platform]

`

	fmt.Print(help)
}

func HandleArgs(args []string) error {
	switch {
	case len(args) <= 1:
		Usage()
		os.Exit(0)

	case len(args) >= 2:
		command := Command(args[0])

		err := command.Validate()
		if err != nil {
			return err
		}

		switch command {
		case Help:
			Usage()

		case Init:
			HandleInit(args[1:])

		case Bubble:
			err := HandleBubble(args[1:])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func HandleInit(args []string) {
	// TODO: Dynamically generate this, possible with runtime pkg.
	goVersion := "go 1.17"

	root := strings.Join(args, "-")

	if root == Help {
		fmt.Printf("Usage: drexler init [root]		This scaffolds out a project.\n\n")
		os.Exit(0)
	}

	fmt.Printf("\n  %s Initialized project %s\n", style.Check, root)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	path := fmt.Sprintf("%s/%s", dir, root)

	dirs := []string{
		"/cmd",
		"/internal",
		"/pkg",
	}

	for _, dir := range dirs {
		fmt.Printf("  %s Created directory %s\n", style.Check, dir)
	}

	subDirs := []string{
		fmt.Sprintf("/cmd/%s", root),
		"/internal/tui",
		"/internal/services",
		"/pkg",
	}

	for _, dir := range subDirs {
		err := os.MkdirAll(path+dir, os.FileMode(0755))
		if err != nil {
			_ = fmt.Errorf("  %s error making subdir: %w", style.X, err)
			os.Exit(1)
		}
		fmt.Printf("  %s Created subdirectory %s\n", style.Check, dir)
	}

	files := make(map[string]string)
	files[fmt.Sprintf("cmd/%s/main.go", root)] = genfile.Main(root)
	files["internal/tui/cmd.go"] = genfile.Cmd()
	files["internal/tui/model.go"] = genfile.Model(root)
	files["internal/tui/msg.go"] = genfile.Msg()
	files["internal/tui/view.go"] = genfile.View
	files[".gitignore"] = genfile.Gitignore
	files["README.md"] = genfile.ReadMe(root)
	files["go.mod"] = genfile.Mod(root, goVersion)
	files["Makefile"] = genfile.Make(root)

	for name, content := range files {
		file, err := os.Create(root + "/" + name)
		if err != nil {
			_ = fmt.Errorf("  %s error creating file %w", style.X, err)
			os.Exit(1)
		}

		fmt.Printf("  %s Created file %s\n", style.Check, name)

		_, err = file.WriteString(content)
		if err != nil {
			fmt.Printf("  %s error writing to file %w", style.X, err)
			os.Exit(1)
		}

		fmt.Printf("  %s Generated boilerplate for %s\n", style.Check, name)
	}

	err = os.Chdir(path)
	if err != nil {
		log.Fatalf("  %s error changing dir: %v\n", style.X, err)
	}

	b, err := exec.Command("make", "tidy").CombinedOutput()
	if err != nil {
		fmt.Printf("  %s error running make deps: %s\n", style.X, string(b))
	}

	fmt.Printf("  %s Installed charmbracelet dependencies\n", style.Check)
}

func HandleBubble(args []string) error {
	b := tui.Bubble(args[1])
	err := b.Validate()
	if err != nil {
		return err
	}

	// TODO: Add codegen for all bubbles.
	switch b {
	case tui.TextInput:

	}

	return nil
}
