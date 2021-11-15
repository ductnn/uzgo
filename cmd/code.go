package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"uzgo/util"

	"github.com/spf13/cobra"
)

var File string

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:                   "code",
	Short:                 "Open source with VSCode",
	DisableFlagsInUseLine: true,

	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) < 1 {
			return errors.New("accepts 1 arg(s)")
		}
		return nil
	},
	Example: `uzo code demo.zip
uzo code /Downloads/application.zip`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(File)
		var filename string
		var err error
		var argument string

		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}

		fileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err)
		}
		if fileExists {
			filename, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("File %v doest not Exists", argument)
			return
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		util.Unzip(filename, wd)

		os.Chdir(util.FilenameWithoutExtension(filename))

		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		commandCode := exec.Command("code", wd)
		print(wd)
		err = commandCode.Run()

		if err != nil {
			log.Fatal("VS Code executable file not found in %PATH%")
		}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	codeCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "A File name to unzip and open in IDE")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
