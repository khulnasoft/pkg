package khulnasoft

import (
	"context"
	"log"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKhulnasoftLabels() *schema.Resource {
	return &schema.Resource{
		Description: "The data source `khulnasoft_khulnasoft_labels` provides a method to query all khulnasoft labels within the Khulnasoft account management." +
			"The fields returned from this query are detailed in the Schema section below.",
		ReadContext: khulnasoftLabelRead,
		Schema: map[string]*schema.Schema{
			"khulnasoft_labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Khulnasoft label name.",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Khulnasoft label description.",
							Computed:    true,
						},
						"created": {
							Type:        schema.TypeString,
							Description: "The creation date of the Khulnasoft label.",
							Computed:    true,
						},
						"author": {
							Type:        schema.TypeString,
							Description: "The name of the user who created the Khulnasoft label.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func khulnasoftLabelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG]  inside resourceKhulnasoftLabelRead")
	c := m.(*client.Client)
	result, err := c.GetKhulnasoftLabels()

	if err != nil {
		return diag.FromErr(err)
	}

	id := ""
	khulnasoftLabels := make([]interface{}, len(result.KhulnasoftLabels), len(result.KhulnasoftLabels))

	for i, khulnasoftLabel := range result.KhulnasoftLabels {
		id = id + khulnasoftLabel.Name
		al := make(map[string]interface{})
		al["name"] = khulnasoftLabel.Name
		al["description"] = khulnasoftLabel.Description
		al["created"] = khulnasoftLabel.Created
		al["author"] = khulnasoftLabel.Author
		khulnasoftLabels[i] = al
	}

	d.SetId(id)
	if err := d.Set("khulnasoft_labels", khulnasoftLabels); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
