package cmd

import (
	"bufio"
	"fmt"
	"github.com/bradcypert/deckard/lib/db"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func downFunc(args []string) {
	bindVarsFromConfig()

	database := db.Database{
		Dbname:   cmdDatabaseName,
		Port:     cmdDatabasePort,
		Password: cmdDatabasePassword,
		User:     cmdDatabaseUser,
		Host:     cmdDatabaseHost,
		Driver:   cmdDatabaseDriver,
	}

	var migration db.Migration
	queries := make([]db.Query, 0)

	if len(args) < 1 {
		// get all migrations in current folder.
		files, err := ioutil.ReadDir(cmdInputDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".down.sql") {
				contents, _ := ioutil.ReadFile(file.Name())
				queries = append(queries, db.Query{
					Name:  file.Name(),
					Value: string(contents),
				})
			}
		}

		ReverseQuerySlice(queries)

		migration = db.Migration{
			Queries: queries,
		}

		// warn the user. Downs are usually destructive.
		fmt.Printf("Heads up! You're about to run DOWN migrations. These migrations are likely destructive.\n")
		fmt.Printf("Would you like to continue? y/N: ")
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if char == 'Y' || char == 'y' {
			database.RunDown(migration, cmdSteps)
		} else {
			log.Fatal("Understood! Aborting...")
		}

	} else {
		//	TODO: What if we have more args?
	}
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Runs one or more \"down\" migrations.",
	Long: `Runs one or more \"down\" migrations.
These migrations are likely destructive. Please use caution when executing deckard down.

Deckard can be instructed to run all down migrations or specific ones.

Running All:
Example:
deckard down

Running One:
Example:
deckard down 1558294955321
# or
deckard down add_users_to_other_users
`,
	Run: func(cmd *cobra.Command, args []string) {
		downFunc(args)
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	addDatabaseFlags(downCmd)

	downCmd.Flags().IntVarP(&cmdSteps,
		"steps",
		"s",
		-1,
		"The number of down migrations you'd like to run.")

	dir, _ := os.Getwd()
	downCmd.Flags().StringVarP(&cmdInputDir,
		"inputDir",
		"i",
		dir,
		"Directory which contains the migrations")
}
