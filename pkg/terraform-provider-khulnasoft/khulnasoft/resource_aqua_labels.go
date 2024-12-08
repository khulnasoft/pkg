package khulnasoft

import (
	"fmt"
	"log"
	"strings"

	"github.com/khulnasoft/terraform-provider-khulnasoft/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceKhulnasoftLabels() *schema.Resource {
	return &schema.Resource{
		Read:   resourceKhulnasoftLabelRead,
		Create: resourceKhulnasoftLabelCreate,
		Update: resourceKhulnasoftLabelUpdate,
		Delete: resourceKhulnasoftLabelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Khulnasoft label name.",
				Required:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Khulnasoft label description.",
				Optional:    true,
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
	}
}

func resourceKhulnasoftLabelCreate(d *schema.ResourceData, m interface{}) error {
	ac := m.(*client.Client)
	khulnasoftLabel := client.KhulnasoftLabel{
		Name: d.Get("name").(string),
	}

	description, ok := d.GetOk("description")
	if ok {
		khulnasoftLabel.Description = description.(string)
	}

	err := ac.CreateKhulnasoftLabel(&khulnasoftLabel)

	if err != nil {
		return err
	}
	d.SetId(khulnasoftLabel.Name)
	return resourceKhulnasoftLabelRead(d, m)
}

func resourceKhulnasoftLabelRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  inside resourceKhulnasoftLabelRead")
	c := m.(*client.Client)
	r, err := c.GetKhulnasoftLabel(d.Id())

	if err != nil {
		if strings.Contains(fmt.Sprintf("%s", err), "404") {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", r.Name)
	d.Set("description", r.Description)
	d.Set("created", r.Created)
	d.Set("author", r.Author)

	return nil
}

func resourceKhulnasoftLabelUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)

	if d.HasChanges("description") {
		khulnasoft_lable := client.KhulnasoftLabel{
			Name: d.Get("name").(string),
		}

		description, ok := d.GetOk("description")
		if ok {
			khulnasoft_lable.Description = description.(string)
		}

		err := c.UpdateKhulnasoftLabel(&khulnasoft_lable)

		if err != nil {
			return err
		}
		d.SetId(d.Get("name").(string))
		return nil
	}
	return resourceKhulnasoftLabelRead(d, m)
}

func resourceKhulnasoftLabelDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)
	id := d.Id()
	err := c.DeleteKhulnasoftLabel(id)
	if err == nil {
		d.SetId("")
	} else {
		return err
	}
	return nil
}
