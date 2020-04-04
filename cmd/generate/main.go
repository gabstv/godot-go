package main

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/gabstv/godot-go/cmd/generate/classes"
	"github.com/gabstv/godot-go/cmd/generate/gdnative"
	"github.com/gabstv/godot-go/cmd/generate/types"
	"github.com/spf13/cobra"
)

var (
	verbose         bool
	clean           bool
	genGodotApiJson bool
	genGdnative     bool
	genTypes        bool
	genClasses      bool
	packagePath     string
	godotPath       string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Toggle extra debug output")
	rootCmd.PersistentFlags().BoolVarP(&clean, "clean", "", true, "Clean generated files")
	rootCmd.PersistentFlags().BoolVarP(&genGodotApiJson, "godot-api-json", "", true, "Generate Godot API JSON")
	rootCmd.PersistentFlags().BoolVarP(&genGdnative, "gdnative", "", true, "Generate gdnative")
	rootCmd.PersistentFlags().BoolVarP(&genTypes, "types", "", true, "Generate types")
	rootCmd.PersistentFlags().BoolVarP(&genClasses, "classes", "", true, "Generate classes")
	rootCmd.PersistentFlags().StringVarP(&packagePath, "package-path", "p", ".", "Specified package path")
	rootCmd.PersistentFlags().StringVarP(&godotPath, "godot-path", "", "godot", "Specified path where the Godot executable is located")
}

func expandPattern(globPattern string) ([]string, error) {
	absGlobPattern, err := filepath.Abs(globPattern)
	if err != nil {
		return nil, err
	}

	files, err := filepath.Glob(absGlobPattern)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func removeFiles(files []string) error {
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return err
		}
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "godot-go",
	Short: "Godot Go",
	RunE: func(cmd *cobra.Command, args []string) error {
		if clean {
			var (
				genFiles []string
			)

			globPatterns := []string{
				"/cmd/generate/templates/godot_api.json",
				"/godot/*.gen.go",
				"/gdnative/*.gen.go",
				"/gdnative/*.gen.c",
				"/gdnative/*.gen.h",
			}

			for _, p := range globPatterns {
				f, err := expandPattern(path.Join(packagePath, p))

				if err != nil {
					return err
				}

				genFiles = append(genFiles, f...)
			}

			if len(genFiles) == 0 {
				println("No generated files found.")
			} else {
				// TODO: test actual remove
				//pretty.Println(genFiles)
				if err := removeFiles(genFiles); err != nil {
					return err
				}
			}
		}

		if genGodotApiJson {
			c := exec.Command("godot",
				"--gdnative-generate-json-api",
				path.Join(packagePath, "cmd/generate/templates/godot_api.json"),
				"--no-window")

			if err := c.Run(); err != nil {
				return err
			}
		}

		if genGdnative {
			gdnative.Generate()
		}

		if genTypes {
			types.Generate()
		}

		if genClasses {
			classes.Generate()
		}
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
