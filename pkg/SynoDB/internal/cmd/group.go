package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal"
	"github.com/khulnasoft/synodb/internal/prompt"
	"github.com/khulnasoft/synodb/internal/synodb"
)

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage your database groups",
}

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(groupsListCmd)
	groupCmd.AddCommand(groupsCreateCmd)
	addLocationFlag(groupsCreateCmd, "Create the group primary in the specified location")
	addWaitFlag(groupsCreateCmd, "Wait for group to be ready")
	groupCmd.AddCommand(groupsDestroyCmd)
	addYesFlag(groupsDestroyCmd, "Confirms the destruction of the group, with all its locations and databases.")
}

var groupsListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List databases groups",
	Args:              cobra.NoArgs,
	ValidArgsFunction: noFilesArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		groups, err := getGroups(client, true)
		if err != nil {
			return err
		}

		printTable([]string{"Name", "Locations"}, groupsTable(groups))
		return nil
	},
}

var groupsCreateCmd = &cobra.Command{
	Use:               "create [group]",
	Short:             "Create a database group",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: noFilesArg,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		name := args[0]
		if err := synodb.CheckName(name); err != nil {
			return fmt.Errorf("invalid group name: %w", err)
		}

		location := locationFlag
		if location == "" {
			location, _ = closestLocation(client)
		}
		if !isValidLocation(client, location) {
			return fmt.Errorf("location '%s' is not a valid one", location)
		}

		return createGroup(client, name, location)
	},
}

var groupsDestroyCmd = &cobra.Command{
	Use:               "destroy [group]",
	Short:             "Destroy a database group",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: groupArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		name := args[0]
		if yesFlag {
			return destroyGroup(client, name)
		}

		fmt.Printf("Group %s, all its replicas and databases will be destroyed.\n", internal.Emph(name))

		ok, err := promptConfirmation("Are you sure you want to do this?")
		if err != nil {
			return fmt.Errorf("could not get prompt confirmed by user: %w", err)
		}

		if !ok {
			fmt.Println("Group destruction avoided.")
			return nil
		}

		return destroyGroup(client, name)
	},
}

func createGroup(client *synodb.Client, name, location string) error {
	start := time.Now()
	description := fmt.Sprintf("Creating group %s at %s...", internal.Emph(name), internal.Emph(location))
	spinner := prompt.Spinner(description)
	defer spinner.Stop()

	invalidateGroupsCache(client.Org)
	if err := client.Groups.Create(name, location); err != nil {
		return err
	}

	if err := handleGroupWaitFlag(client, name, location); err != nil {
		return err
	}

	spinner.Stop()
	elapsed := time.Since(start)
	fmt.Printf("Created group %s at %s in %d seconds.\n", internal.Emph(name), internal.Emph(location), int(elapsed.Seconds()))

	return nil
}

func destroyGroup(client *synodb.Client, name string) error {
	start := time.Now()
	s := prompt.Spinner(fmt.Sprintf("Destroying group %s... ", internal.Emph(name)))
	defer s.Stop()

	invalidateGroupsCache(client.Org)
	if err := client.Groups.Delete(name); err != nil {
		return err
	}
	s.Stop()
	elapsed := time.Since(start)

	fmt.Printf("Destroyed group %s in %d seconds.\n", internal.Emph(name), int(elapsed.Seconds()))
	return nil
}

func groupsTable(groups []synodb.Group) [][]string {
	var data [][]string
	for _, group := range groups {
		row := []string{group.Name, formatLocations(group.Locations, group.Primary)}
		data = append(data, row)
	}
	return data
}

func getGroups(client *synodb.Client, fresh ...bool) ([]synodb.Group, error) {
	skipCache := len(fresh) > 0 && fresh[0]
	if cached := getGroupsCache(client.Org); !skipCache && cached != nil {
		return cached, nil
	}
	groups, err := client.Groups.List()
	if err != nil {
		return nil, err
	}
	setGroupsCache(client.Org, groups)
	return groups, nil
}

func groupNames(client *synodb.Client) ([]string, error) {
	groups, err := getGroups(client)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(groups))
	for _, group := range groups {
		names = append(names, group.Name)
	}
	return names, nil
}

func getGroup(client *synodb.Client, name string) (synodb.Group, error) {
	groups, err := getGroups(client)
	if err != nil {
		return synodb.Group{}, err
	}
	for _, group := range groups {
		if group.Name == name {
			return group, nil
		}
	}
	return synodb.Group{}, fmt.Errorf("group %s was not found", name)
}

func handleGroupWaitFlag(client *synodb.Client, group, location string) error {
	if !waitFlag {
		return nil
	}
	return client.Groups.WaitLocation(group, location)
}
