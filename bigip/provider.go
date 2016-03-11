package bigip

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

const DEFAULT_PARTITION = "Common"

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name/IP of the BigIP",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username with API access to the BigIP",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user's password",
			},
			"loginReference": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Login reference for token authentication (see BIG-IP REST docs for details)",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"bigip_ltm_virtual_server": resourceBigipLtmVirtualServer(),
			"bigip_ltm_node":           resourceBigipLtmNode(),
			"bigip_ltm_pool":           resourceBigipLtmPool(),
			"bigip_ltm_monitor":        resourceBigipLtmMonitor(),
			"bigip_ltm_irule":          resourceBigipLtmIRule(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Address:        d.Get("address").(string),
		Username:       d.Get("username").(string),
		Password:       d.Get("password").(string),
		LoginReference: d.Get("loginReference").(string),
	}

	return config.Client()
}

func makeStringSet(list *[]string) *schema.Set {
	ilist := make([]interface{}, len(*list))
	for i, v := range *list {
		ilist[i] = v
	}
	return schema.NewSet(schema.HashString, ilist)
}

func setToStringSlice(s *schema.Set) []string {
	list := make([]string, s.Len())
	for i, v := range s.List() {
		list[i] = v.(string)
	}
	return list
}
