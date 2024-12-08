package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/khulnasoft/synodb/internal"
	"github.com/khulnasoft/synodb/internal/prompt"
	"github.com/khulnasoft/synodb/internal/settings"
	"github.com/khulnasoft/synodb/internal/synodb"
	"golang.org/x/sync/errgroup"
)

func init() {
	rootCmd.AddCommand(planCmd)
	planCmd.AddCommand(planShowCmd)
	planCmd.AddCommand(planSelectCmd)
	planCmd.AddCommand(planUpgradeCmd)
	planCmd.AddCommand(overagesCommand)
	overagesCommand.AddCommand(planEnableOverages)
	overagesCommand.AddCommand(planDisableOverages)
}

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Manage your organization plan",
}

var overagesCommand = &cobra.Command{
	Use:   "overages",
	Short: "Manage your current organization overages",
}

func getCurrentOrg(client *synodb.Client, organizationName string) (synodb.Organization, error) {
	orgs, err := client.Organizations.List()
	if err != nil {
		return synodb.Organization{}, err
	}
	for _, org := range orgs {
		if org.Slug == organizationName {
			return org, nil
		}
	}
	return synodb.Organization{}, fmt.Errorf("could not find organization %s", organizationName)
}

var planShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show your current organization plan",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		settings, err := settings.ReadSettings()
		if err != nil {
			return fmt.Errorf("could not retrieve local config: %w", err)
		}

		plan, orgUsage, plans, err := orgPlanData(client)
		if err != nil {
			return err
		}

		var organizationName string
		if organizationName = client.Org; organizationName == "" {
			organizationName = settings.GetUsername()
		}

		currentOrg, err := getCurrentOrg(client, organizationName)
		if err != nil {
			return err
		}

		fmt.Printf("Organization: %s\n", internal.Emph(organizationName))
		if currentOrg.Overages {
			plan, _ = strings.CutSuffix(plan, "_overages")
		}

		fmt.Printf("Plan: %s\n", internal.Emph(plan))
		fmt.Print(overagesMessage(currentOrg.Overages))
		fmt.Println()

		current := getPlan(plan, plans)
		tbl := planUsageTable(orgUsage, current, currentOrg)
		tbl.Print()
		fmt.Printf("\nQuota will reset on %s\n", getFirstDayOfNextMonth().Local().Format(time.RFC1123))
		return nil
	},
}

func overagesMessage(overages bool) string {
	status := "disabled"
	if overages {
		status = "enabled"
	}
	return fmt.Sprintf("Overages %s\n", internal.Emph(status))
}

func planUsageTable(orgUsage synodb.OrgUsage, current synodb.Plan, currentOrg synodb.Organization) table.Table {
	columns := make([]interface{}, 0)
	columns = append(columns, "RESOURCE")
	columns = append(columns, "USED")

	columns = append(columns, "LIMIT")
	columns = append(columns, "LIMIT %")
	if currentOrg.Overages {
		columns = append(columns, "OVERAGE")
	}

	tbl := table.New(columns...)

	columnFmt := color.New(color.FgBlue, color.Bold).SprintfFunc()
	tbl.WithFirstColumnFormatter(columnFmt)

	addResourceRowBytes(tbl, "storage", orgUsage.Usage.StorageBytesUsed, current.Quotas.Storage, currentOrg.Overages)
	addResourceRowMillions(tbl, "rows read", orgUsage.Usage.RowsRead, current.Quotas.RowsRead, currentOrg.Overages)
	addResourceRowMillions(tbl, "rows written", orgUsage.Usage.RowsWritten, current.Quotas.RowsWritten, currentOrg.Overages)
	addResourceRowCount(tbl, "databases", orgUsage.Usage.Databases, current.Quotas.Databases)
	addResourceRowCount(tbl, "locations", orgUsage.Usage.Locations, current.Quotas.Locations)
	return tbl
}

func orgPlanData(client *synodb.Client) (sub string, usage synodb.OrgUsage, plans []synodb.Plan, err error) {
	g := errgroup.Group{}
	g.Go(func() (err error) {
		sub, err = client.Subscriptions.Get()
		return
	})

	g.Go(func() (err error) {
		usage, err = client.Organizations.Usage()
		return
	})

	g.Go(func() (err error) {
		plans, err = client.Plans.List()
		return
	})
	err = g.Wait()
	return
}

var planSelectCmd = &cobra.Command{
	Use:   "select",
	Short: "Change your current organization plan",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		plans, current, hasPaymentMethod, err := GetSelectPlanInfo(client)
		if err != nil {
			return fmt.Errorf("failed to get plans: %w", err)
		}
		selectabledPlans, err := getSelectabledPlans(plans)
		if err != nil {
			return err
		}
		selected, err := promptPlanSelection(selectabledPlans, current)
		if err != nil {
			return err
		}

		return ChangePlan(client, plans, current, hasPaymentMethod, selected)
	},
}

func getSelectabledPlans(plans []synodb.Plan) ([]synodb.Plan, error) {
	settings, err := settings.ReadSettings()
	if err != nil {
		return plans, err
	}

	org := settings.Organization()
	var plansToSelect []synodb.Plan
	for _, plan := range plans {
		if plan.Name != "starter" || org == "" {
			plansToSelect = append(plansToSelect, plan)
		}
	}
	return plansToSelect, nil
}

var planUpgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade your current organization plan",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		plans, current, hasPaymentMethod, err := GetSelectPlanInfo(client)
		if err != nil {
			return fmt.Errorf("failed to get plans: %w", err)
		}

		if current == "scaler" {
			fmt.Printf("You've already upgraded to %s! 🎉\n", internal.Emph(current))
			fmt.Println()
			fmt.Println("If you need more resources, we're happy to help.")
			fmt.Printf("You can find us at %s or at Discord (%s)\n", internal.Emph("sales@synodb.tech"), internal.Emph("https://discord.com/invite/4B5D7hYwub"))
			return nil
		}

		return ChangePlan(client, plans, current, hasPaymentMethod, "scaler")
	},
}

var planEnableOverages = &cobra.Command{
	Use:   "enable",
	Short: "Enable overages for your current organization plan",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		settings, err := settings.ReadSettings()
		if err != nil {
			return err
		}
		var org string
		if org = settings.Organization(); org == "" {
			org = settings.GetUsername()
		}
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		hasPaymentMethod, err := client.Billing.HasPaymentMethod()
		if err != nil {
			return err
		}
		if !hasPaymentMethod {
			ok, err := PaymentMethodHelperOverages(client)
			if err != nil {
				return fmt.Errorf("failed to check payment method: %w", err)
			}
			if !ok {
				return fmt.Errorf("failed to add payment method")
			}
			fmt.Println("Payment method added successfully.")
			fmt.Printf("You can manage your payment methods with %s.\n\n", internal.Emph("synodb org billing"))
		}
		if err = client.Organizations.SetOverages(org, true); err != nil {
			return err
		}
		fmt.Println("Overages enabled successfully.")
		return nil
	},
}

var planDisableOverages = &cobra.Command{
	Use:   "disable",
	Short: "Disable overages for your current organization plan",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		settings, err := settings.ReadSettings()
		if err != nil {
			return err
		}
		var org string
		if org = settings.Organization(); org == "" {
			org = settings.GetUsername()
		}
		client, err := createSynodbClientFromAccessToken(true)
		if err != nil {
			return err
		}

		if err = client.Organizations.SetOverages(org, false); err != nil {
			return err
		}
		fmt.Println("Overages disabled successfully.")
		return nil
	},
}

func ChangePlan(client *synodb.Client, plans []synodb.Plan, current string, hasPaymentMethod bool, selected string) error {
	if selected == current {
		fmt.Println("You're all set! No changes are needed.")
		return nil
	}

	upgrade := isUpgrade(getPlan(current, plans), getPlan(selected, plans))
	if !hasPaymentMethod && upgrade {
		ok, err := PaymentMethodHelper(client, selected)
		if err != nil {
			return fmt.Errorf("failed to check payment method: %w", err)
		}
		if !ok {
			return nil
		}
		fmt.Println("Payment method added successfully.")
		fmt.Printf("You can manage your payment methods with %s.\n\n", internal.Emph("synodb org billing"))
	}

	change := "downgrading"
	if upgrade {
		change = "upgrading"
	}
	fmt.Printf("You're %s to the %s plan.\n", change, internal.Emph(selected))

	if upgrade && hasPaymentMethod {
		printPricingInfoDisclaimer()
	}

	if ok, _ := promptConfirmation("Do you want to continue?"); !ok {
		fmt.Printf("Plan change cancelled. You're still on %s.\n", internal.Emph(current))
		return nil
	}

	if err := client.Subscriptions.Set(selected); err != nil {
		return err
	}

	fmt.Println()

	change = "downgraded"
	if upgrade {
		change = "upgraded"
	}
	fmt.Printf("You've been %s to plan %s.\n", change, internal.Emph(selected))
	fmt.Printf("To see your new quotas, use %s.\n", internal.Emph("synodb plan show"))
	return nil
}

func PaymentMethodHelper(client *synodb.Client, selected string) (bool, error) {
	fmt.Printf("You need to add a payment method before you can upgrade to the %s plan.\n", internal.Emph(selected))
	printPricingInfoDisclaimer()

	ok, _ := promptConfirmation("Want to add a payment method right now?")
	if !ok {
		fmt.Printf("When you're ready, you can use %s to manage your payment methods.\n", internal.Emph("synodb org billing"))
		return false, nil
	}

	fmt.Println()
	if err := billingPortal(client); err != nil {
		return false, err
	}
	fmt.Println()

	spinner := prompt.Spinner("Waiting for you to add a payment method")
	defer spinner.Stop()

	return checkPaymentMethod(client, "")
}

func hasPaymentMethodCheck(client *synodb.Client, stripeId string) (bool, error) {
	if stripeId != "" {
		return client.Billing.HasPaymentMethodWithStripeId(stripeId)
	}
	return client.Billing.HasPaymentMethod()
}

func checkPaymentMethod(client *synodb.Client, stripeId string) (bool, error) {
	errsInARoW := 0
	var hasPaymentMethod bool
	var err error
	for {
		hasPaymentMethod, err = hasPaymentMethodCheck(client, stripeId)
		if err != nil {
			errsInARoW += 1
		}
		if errsInARoW > 5 {
			return false, err
		}
		if err == nil {
			errsInARoW = 0
		}
		if hasPaymentMethod {
			return true, nil
		}
		time.Sleep(1 * time.Second)
	}
}

func PaymentMethodHelperOverages(client *synodb.Client) (bool, error) {
	fmt.Print("You need to add a payment method before you can enable overages.\n")
	printPricingInfoDisclaimer()

	ok, _ := promptConfirmation("Want to add a payment method right now?")
	if !ok {
		fmt.Printf("When you're ready, you can use %s to manage your payment methods.\n", internal.Emph("synodb org billing"))
		return false, nil
	}

	fmt.Println()
	if err := billingPortal(client); err != nil {
		return false, err
	}
	fmt.Println()

	spinner := prompt.Spinner("Waiting for you to add a payment method")
	defer spinner.Stop()

	return checkPaymentMethod(client, "")
}

func PaymentMethodHelperWithStripeId(client *synodb.Client, stripeId, orgName string) (bool, error) {
	fmt.Printf("You need to add a payment method before you can create organization %s on the %s plan.\n", internal.Emph(orgName), internal.Emph("scaler"))
	printPricingInfoDisclaimer()

	ok, _ := promptConfirmation("Want to add a payment method right now?")
	if !ok {
		fmt.Printf("When you're ready, you can use %s to manage your payment methods.\n", internal.Emph("synodb org billing"))
		return false, nil
	}

	fmt.Println()
	if err := BillingPortalForStripeId(client, stripeId); err != nil {
		return false, err
	}
	fmt.Println()

	spinner := prompt.Spinner("Waiting for you to add a payment method")
	defer spinner.Stop()

	return checkPaymentMethod(client, stripeId)
}

func GetSelectPlanInfo(client *synodb.Client) (plans []synodb.Plan, current string, hasPaymentMethod bool, err error) {
	g := errgroup.Group{}
	g.Go(func() (err error) {
		plans, err = client.Plans.List()
		return
	})
	g.Go(func() (err error) {
		current, err = client.Subscriptions.Get()
		return
	})
	g.Go(func() (err error) {
		hasPaymentMethod, err = client.Billing.HasPaymentMethod()
		return
	})
	err = g.Wait()
	return
}

func promptPlanSelection(plans []synodb.Plan, current string) (string, error) {
	planNames := make([]string, 0, len(plans))
	cur := 0
	for _, plan := range plans {
		if plan.Name == current {
			cur = len(planNames)
			planNames = append(planNames, fmt.Sprintf("%s (current)", internal.Emph(plan.Name)))
			continue
		}
		planNames = append(planNames, plan.Name)
	}

	prompt := promptui.Select{
		CursorPos:    cur,
		HideHelp:     true,
		Label:        "Select a plan",
		Items:        planNames,
		HideSelected: true,
	}

	_, result, err := prompt.Run()
	if strings.HasSuffix(result, "(current)") {
		result = current
	}
	return result, err
}

func isUpgrade(current, selected synodb.Plan) bool {
	cp, _ := strconv.Atoi(current.Price)
	sp, _ := strconv.Atoi(selected.Price)
	return sp > cp
}

func getPlan(name string, plans []synodb.Plan) synodb.Plan {
	for _, plan := range plans {
		if plan.Name == name {
			return plan
		}
	}
	return synodb.Plan{}
}

func billingPortal(client *synodb.Client) error {
	portal, err := client.Billing.Portal()
	if err != nil {
		return err
	}

	msg := "Opening your browser at:"
	if err := browser.OpenURL(portal.URL); err != nil {
		msg = "Access the following URL to manage your payment methods:"
	}
	fmt.Println(msg)
	fmt.Println(portal.URL)
	return nil
}

func BillingPortalForStripeId(client *synodb.Client, stripeCustomerId string) error {
	portal, err := client.Billing.PortalForStripeId(stripeCustomerId)
	if err != nil {
		return err
	}

	msg := "Opening your browser at:"
	if err := browser.OpenURL(portal.URL); err != nil {
		msg = "Access the following URL to manage your payment methods:"
	}
	fmt.Println(msg)
	fmt.Println(portal.URL)
	return nil
}

func printPricingInfoDisclaimer() {
	fmt.Printf("For information about Synodb plans pricing and features, access: %s\n\n", internal.Emph("https://synodb.tech/pricing"))
}

func addResourceRowBytes(tbl table.Table, resource string, used, limit uint64, overages bool) {
	if limit == 0 {
		tbl.AddRow(resource, humanize.Bytes(used), "Unlimited", "")
		return
	}
	exceededValue := uint64(max(int(used)-int(limit), 0))
	if overages && exceededValue > 0 {
		tbl.AddRow(resource, humanize.Bytes(used), humanize.Bytes(limit), percentage(float64(used), float64(limit)), humanize.Bytes(exceededValue))
		return
	}
	tbl.AddRow(resource, humanize.Bytes(used), humanize.Bytes(limit), percentage(float64(used), float64(limit)))
}

func addResourceRowMillions(tbl table.Table, resource string, used, limit uint64, overages bool) {
	if limit == 0 {
		tbl.AddRow(resource, used, "Unlimited", "")
		return
	}
	exceededValue := uint64(max(int(used)-int(limit), 0))
	if overages && exceededValue > 0 {
		tbl.AddRow(resource, toM(used), toM(limit), percentage(float64(used), float64(limit)), toM(exceededValue))
		return
	}
	tbl.AddRow(resource, toM(used), toM(limit), percentage(float64(used), float64(limit)))
}

func toM(v uint64) string {
	str := fmt.Sprintf("%.1f", float64(v)/1_000_000.0)
	str = strings.TrimSuffix(str, ".0")
	if str == "0" && v != 0 {
		str = "<0.1"
	}
	return str + "M"
}

func addResourceRowCount(tbl table.Table, resource string, used, limit uint64) {
	if limit == 0 {
		tbl.AddRow(resource, used, "Unlimited", "")
		return
	}
	tbl.AddRow(resource, used, limit, percentage(float64(used), float64(limit)))
}

func percentage(used, limit float64) string {
	return fmt.Sprintf("%.0f%%", used/limit*100)
}

func getFirstDayOfNextMonth() time.Time {
	now := time.Now().UTC()
	nextMonth := now.AddDate(0, 1, 0)
	year := nextMonth.Year()
	month := nextMonth.Month()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return firstDay
}
