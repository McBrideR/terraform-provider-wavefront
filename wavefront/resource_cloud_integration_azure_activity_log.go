package wavefront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudIntegrationAzureActivityLog() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudIntegrationCreate,
		Read:   resourceCloudIntegrationRead,
		Update: resourceCloudIntegrationUpdate,
		Delete: resourceCloudIntegrationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"additional_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"force_save": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"service": serviceSchemaDefinition(wfAzureActivityLog),
			"service_refresh_rate_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"category_filter": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": {
				Sensitive: true,
				Type:      schema.TypeString,
				Required:  true,
			},
			"tenant": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
