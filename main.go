package main

import (
	"fmt"
	"os"

	"github.com/aldinokemal/go-whatsapp-web-multidevice/cmd"
)

// @title			Go WhatsApp Web Multi Device API
// @version		2.0
// @description	This is a WhatsApp Web API using multi-device support
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	https://github.com/aldinokemal/go-whatsapp-web-multidevice
// @license.name	MIT
// @license.url	https://opensource.org/licenses/MIT
// @host			localhost:3000
// @BasePath		/
func main() {
	if err := cmd.Execute(); err != nil {
		// Print error to stderr and exit with a non-zero status code.
		// Using exit code 2 to distinguish between usage errors (1) and
		// runtime errors (2) — more informative for scripting purposes.
		//
		// NOTE(personal): keeping exit code 2 here; wrap the binary in a
		// small shell script that logs this output to ~/whatsapp-api.log
		// if running unattended via cron.
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(2)
	}
}
