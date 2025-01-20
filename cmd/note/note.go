package note

import (
	"cln/logger"
	"cln/renderer"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var headingFlag string
var listFlag bool

var NoteCmd = &cobra.Command{
	Use:   "note",
	Short: "Create a new note",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	NoteCmd.PersistentFlags().StringVar(&headingFlag, "heading", "", "Heading for you notes")
	NoteCmd.PersistentFlags().BoolVar(&listFlag, "list", false, "List of notes")
}

func initConfig() {
	if headingFlag != "" {

		if err := os.Mkdir("notes", os.ModePerm); err != nil {
			logger.Error("note", err.Error())
		}

		file, err := os.Create("notes/" + headingFlag + ".md")

		if err != nil {
			logger.Error("note", "not able to create a new file")
		} else {
			logger.Info("note", file.Name())
		}
	} else if listFlag {
		dir, _ := os.ReadDir("notes")
		if len(dir) > 0 {
			defer logger.Info("note", "working")
			fmt.Print(renderer.Render("# List of notes"))
			for index, note := range dir {
				fmt.Print(renderer.Render("- " + string(index) + string(note.Name())))
			}
		}
	}

}
