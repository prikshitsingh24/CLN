package cmd

import (
	"cln/cmd/note"
	"cln/renderer"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var info string = `

# **CLN** - Command Line Notes v(1.0)

With CLN, you can easily manage your to-do lists, notes, tasks, and more — all through a clean, fast, and minimalistic CLI interface.

## **Features**

- **Task Management**: Add, view, and delete tasks
- **Customizable Labels**: Tag tasks with labels for better organization
- **Notes Section**: Write quick notes directly from the CLI
- **Progress Tracker**: Mark tasks as completed and track progress
- **Simple Interface**: No frills, just productivity
- **Cross-Platform**: Works on Linux, macOS, and Windows

* [x] Carrots
* [x] Ramen
* [ ] Currywurst
`

var rootCmd = &cobra.Command{
	Use:   "cln",
	Short: "cln is a cli based notes app",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(renderer.Render(info))
	},
}

func init() {
	rootCmd.AddCommand(note.NoteCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
