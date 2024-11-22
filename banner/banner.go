package banner

import (
	"fmt"
)

// prints the version message
const version = "v0.0.2"

func PrintVersion() {
	fmt.Printf("Current wordgen version %s\n", version)
}

// Prints the Colorful banner
func PrintBanner() {
	banner := `
                             __                  
 _      __ ____   _____ ____/ /____ _ ___   ____ 
| | /| / // __ \ / ___// __  // __  // _ \ / __ \
| |/ |/ // /_/ // /   / /_/ // /_/ //  __// / / /
|__/|__/ \____//_/    \__,_/ \__, / \___//_/ /_/ 
                            /____/
`
	fmt.Printf("%s\n%50s\n\n", banner, "Current wordgen version "+version)
}
