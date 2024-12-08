package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal/prompt"
	"github.com/khulnasoft/synodb/internal/synodb"
)

var groupBoolFlag bool

func addGroupBoolFlag(cmd *cobra.Command, description string) {
	cmd.Flags().BoolVar(&groupBoolFlag, "group", false, description)
}

var groupFlag string

func addGroupFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&groupFlag, "group", "", "create the database in the specified group")
	cmd.RegisterFlagCompletionFunc("group", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		client, err := createSynodbClientFromAccessToken(false)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		groups, _ := groupNames(client)
		return groups, cobra.ShellCompDirectiveNoFileComp
	})
}

var (
	fromDBFlag    string
	timestampFlag string
)

func addFromDBFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&fromDBFlag, "from-db", "", "Select another database to copy data from. To use data from a past version of the selected database, see the 'timestamp' flag.")
	cmd.RegisterFlagCompletionFunc("from-db", dbNameArg)
	cmd.Flags().StringVar(&timestampFlag, "timestamp", "", "Set a point in time in the past to copy data from the selected database. Must be used with the 'from-db' flag. Must be in RFC3339 format like '2023-09-29T10:16:13-03:00'")
}

func parseTimestampFlag() (*time.Time, error) {
	if timestampFlag == "" {
		return nil, nil
	}
	if fromDBFlag == "" {
		return nil, fmt.Errorf("--timestamp cannot be used without specifying --from-db")
	}

	timestamp, err := time.Parse(time.RFC3339, timestampFlag)
	if err != nil {
		return nil, fmt.Errorf("provided timestamp was not in RFC3339 format like '2023-09-29T10:16:13-03:00'")
	}
	return &timestamp, nil
}

func parseDBSeedFlags(client *synodb.Client) (*synodb.DBSeed, error) {
	if countFlags(fromDBFlag, fromDumpFlag, fromFileFlag) > 1 {
		return nil, fmt.Errorf("only one of --from-db, --from-dump, or --from-file can be used at a time")
	}

	timestamp, err := parseTimestampFlag()
	if err != nil {
		return nil, err
	}

	if fromDBFlag != "" {
		return &synodb.DBSeed{Type: "database", Name: fromDBFlag, Timestamp: timestamp}, nil
	}

	if fromFileFlag != "" {
		return handleDBFile(client, fromFileFlag)
	}

	if fromDumpFlag != "" {
		return handleDumpFile(client, fromDumpFlag)
	}

	return nil, nil
}

func handleDumpFile(client *synodb.Client, file string) (*synodb.DBSeed, error) {
	dump, err := validateDumpFile(file)
	if err != nil {
		return nil, err
	}

	start := time.Now()
	spinner := prompt.Spinner("Uploading data...")
	defer spinner.Stop()

	dumpURL, err := client.Databases.UploadDump(dump)
	if err != nil {
		return nil, fmt.Errorf("could not upload dump: %w", err)
	}

	spinner.Stop()
	elapsed := time.Since(start)
	fmt.Printf("Uploaded data in %d seconds.\n\n", int(elapsed.Seconds()))

	return &synodb.DBSeed{
		Type: "dump",
		URL:  dumpURL,
	}, nil
}

func validateDumpFile(name string) (*os.File, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %w", name, err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("could not stat file %s: %w", name, err)
	}
	if fileStat.Size() == 0 {
		return nil, fmt.Errorf("dump file is empty")
	}
	if fileStat.Size() > MaxDumpFileSizeBytes {
		return nil, fmt.Errorf("dump file is too large. max allowed size is 2GB")
	}
	return file, nil
}

func countFlags(flags ...string) (count int) {
	for _, flag := range flags {
		if flag != "" {
			count++
		}
	}
	return
}

func handleDBFile(client *synodb.Client, file string) (*synodb.DBSeed, error) {
	if err := checkFileExists(file); err != nil {
		return nil, err
	}
	if err := checkSQLiteAvailable(); err != nil {
		return nil, err
	}

	if err := checkSQLiteFile(file); err != nil {
		return nil, err
	}

	tmp, err := createTempFile()
	if err != nil {
		return nil, err
	}

	if err := dumpSQLiteDatabase(file, tmp); err != nil {
		return nil, err
	}

	return handleDumpFile(client, tmp.Name())
}

func checkFileExists(file string) error {
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("could not find file %s", file)
	}
	return err
}

func checkSQLiteAvailable() error {
	_, err := exec.LookPath("sqlite3")
	if errors.Is(err, exec.ErrNotFound) {
		return fmt.Errorf("could not find sqlite3 on your system. Please install it to use the --from-file flag or use --from-dump instead")
	}
	return err
}

func checkSQLiteFile(file string) error {
	output, err := exec.Command("sqlite3", file, "pragma quick_check;").CombinedOutput()

	execErr := &exec.ExitError{}
	if errors.As(err, &execErr) && execErr.ExitCode() == 26 {
		return fmt.Errorf("file %s is not a valid SQLite database file", file)
	}

	if err != nil {
		return fmt.Errorf("could not check database file: %w: %s", err, output)
	}
	return nil
}

func createTempFile() (*os.File, error) {
	tmp, err := os.CreateTemp("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create temporary file to dump database file: %w", err)
	}
	return tmp, nil
}

func dumpSQLiteDatabase(database string, dump *os.File) error {
	stdErr := &bytes.Buffer{}
	cmd := exec.Command("sqlite3", database, ".dump")
	cmd.Stdout = dump
	cmd.Stderr = stdErr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("could not dump database file: %w: %x", err, stdErr.Bytes())
	}

	return nil
}
