package utils

import (
	"fmt"
	"net"
)

// ValidateIPAddr checks if the provided string is a valid IP address.
// The function simply uses the net.ParseIP standard library function to parse the string.
// If the parsing is successful, the function returns nil. Otherwise, it returns an error.
//
// This doesn't need to be its own function, but it feels better this way.
//
// Parameters:
//   - s: The string to be validated as an IP address.
//
// Returns:
//   - error: nil if the string is a valid IP address, or an error if it is not.
func ValidateIPAddr(s string) error {
	ip := net.ParseIP(s)
	if ip != nil {
		return nil
	}
	return fmt.Errorf("invalid IP address: %s", s)
}
