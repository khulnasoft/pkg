package cmd

import (
	"fmt"
	"math"
	"os"
	"sync"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal"
	"github.com/khulnasoft/synodb/internal/settings"
	"github.com/khulnasoft/synodb/internal/synodb"
)

var (
	showUrlFlag          bool
	showHttpUrlFlag      bool
	showInstanceUrlsFlag bool
	showInstanceUrlFlag  string
)

func getInstanceNames(client *synodb.Client, dbName string) []string {
	instances, err := client.Instances.List(dbName)
	if err != nil {
		return nil
	}
	result := []string{}
	for _, instance := range instances {
		result = append(result, instance.Name)
	}
	return result
}

func extractDatabaseNames(databases []synodb.Database) []string {
	names := make([]string, 0)
	for _, database := range databases {
		names = append(names, database.Name)
	}
	return names
}

func getDatabase(client *synodb.Client, name string, fresh ...bool) (synodb.Database, error) {
	databases, err := getDatabases(client, fresh...)
	if err != nil {
		return synodb.Database{}, err
	}

	for _, database := range databases {
		if database.Name == name {
			return database, nil
		}
	}

	return synodb.Database{}, fmt.Errorf("database %s not found. List known databases using %s", internal.Emph(name), internal.Emph("synodb db list"))
}

func getDatabases(client *synodb.Client, fresh ...bool) ([]synodb.Database, error) {
	skipCache := len(fresh) > 0 && fresh[0]
	if cachedNames := getDatabasesCache(); !skipCache && cachedNames != nil {
		return cachedNames, nil
	}
	databases, err := client.Databases.List()
	if err != nil {
		return nil, err
	}
	setDatabasesCache(databases)
	return databases, nil
}

func getDatabaseNames(client *synodb.Client) []string {
	databases, err := getDatabases(client)
	if err != nil {
		return []string{}
	}
	return extractDatabaseNames(databases)
}

func completeInstanceName(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	client, err := createSynodbClientFromAccessToken(false)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	if len(args) == 1 {
		return getInstanceNames(client, args[0]), cobra.ShellCompDirectiveNoFileComp
	}
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	rootCmd.AddCommand(dbCmd)
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Manage databases",
}

const ENV_ACCESS_TOKEN = "SYNODB_API_TOKEN"

func getAccessToken(warnMultipleAccessTokenSources bool) (string, error) {
	envToken := os.Getenv(ENV_ACCESS_TOKEN)
	settings, err := settings.ReadSettings()
	if err != nil {
		return "", fmt.Errorf("could not read local settings")
	}
	settingsToken := settings.GetToken()

	if !noMultipleTokenSourcesWarning && envToken != "" && settingsToken != "" && warnMultipleAccessTokenSources {
		fmt.Printf("Warning: User logged in as %s but SYNODB_API_TOKEN environment variable is set so proceeding to use it instead\n\n", settings.GetUsername())
	}
	if envToken != "" {
		return envToken, nil
	}
	if !isJwtTokenValid(settingsToken) {
		return "", fmt.Errorf("user not logged in, please login with %s", internal.Emph("synodb auth login"))
	}

	return settingsToken, nil
}

type latMap struct {
	id  string
	lat int
}

func locations(client *synodb.Client) (map[string]string, error) {
	settings, _ := settings.ReadSettings()
	return readLocations(settings, client)
}

func latencies(client *synodb.Client) (map[string]int, error) {
	settings, _ := settings.ReadSettings()
	locations, err := readLocations(settings, client)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	latencies := make(map[string]int)
	c := make(chan latMap, len(locations))
	for id := range locations {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			measure := math.MaxInt
			// XXX: Running this in different goroutines makes all latencies dogslow.
			// Not sure if this is contention at the client or API level
			for i := 0; i < 3; i++ {
				d := synodb.ProbeLocation(id)
				if d != nil {
					measure = int(math.Min(float64(d.Milliseconds()), float64(measure)))
				}
			}
			c <- latMap{id: id, lat: measure}
		}(id)
	}

	wg.Wait()
	close(c)

	for kvp := range c {
		latencies[kvp.id] = kvp.lat
	}
	return latencies, nil
}

func readLocations(settings *settings.Settings, client *synodb.Client) (map[string]string, error) {
	if locations := locationsCache(); locations != nil {
		return locations, nil
	}

	locations, err := client.Locations.List()
	if err != nil {
		return nil, err
	}

	setLocationsCache(locations)
	return locations, nil
}

func closestLocation(client *synodb.Client) (string, error) {
	if closest := closestLocationCache(); closest != "" {
		return closest, nil
	}

	closest, err := client.Locations.Closest()
	if err != nil {
		// We fallback to ams if we are unable to probe the closest location.
		return "ams", err
	}

	setClosestLocationCache(closest)
	return closest, nil
}

func isValidLocation(client *synodb.Client, location string) bool {
	locations, err := locations(client)
	if err != nil {
		return true
	}
	_, ok := locations[location]
	return ok
}

func formatLocation(client *synodb.Client, id string) string {
	locations, _ := locations(client)
	if desc, ok := locations[id]; ok {
		return fmt.Sprintf("%s (%s)", desc, id)
	}
	return fmt.Sprintf("Location ID: %s", id)
}
