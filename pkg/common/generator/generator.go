package generator

import (
	"fmt"

	"github.com/spf13/viper"
)

const urlIntegration = "integrations.externals.http.%s.endpoints.%s"
const host = "integrations.externals.http.%s.host"

func GetIntegURL(integration string, name string) string {
	getHost := fmt.Sprintf(host, integration)
	getString := fmt.Sprintf(urlIntegration, integration, name)
	return viper.GetString(getHost) + viper.GetString(getString)
}
