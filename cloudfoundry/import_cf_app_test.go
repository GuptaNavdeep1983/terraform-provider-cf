package cloudfoundry

import (
	"testing"

	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccApp_importBasic(t *testing.T) {
	resourceName := "cloudfoundry_app.java-spring"

	resource.Test(t,
		resource.TestCase{
			PreCheck:     func() { testAccPreCheck(t) },
			Providers:    testAccProviders,
			CheckDestroy: testAccCheckAppDestroyed([]string{"java-spring"}),
			Steps: []resource.TestStep{

				resource.TestStep{
					Config: fmt.Sprintf(appResourceJavaSpring, defaultAppDomain()),
				},

				resource.TestStep{
					ResourceName:      resourceName,
					ImportState:       true,
					ImportStateVerify: true,
					ImportStateVerifyIgnore: []string{
						"timeout",
						"route",
						"url",
						"service_binding.0.credentials",
						"service_binding.1.credentials",
						"buildpack",
						"command",
						"health_check_http_endpoint",
						"health_check_timeout",
					},
				},
			},
		})
}
