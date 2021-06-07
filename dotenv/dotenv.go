package dotenv

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Load initialises all key value pairs from a given .env file
// and sets those as environment variables on the current process.
func Load(filePath string, optional bool) error {

	// Check if file exists
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) && optional {
			return nil
		}
		return fmt.Errorf("failed to retrieve file information for file %s with error: %w", filePath, err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s with error: %w", filePath, err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(strings.Trim(line, " ")) != 0 {
			kv := strings.SplitN(line, "=", 2)
			if len(kv) != 2 {
				return errors.New("cannot parse line: " + line)
			}

			os.Setenv(kv[0], kv[1])
		}
	}
	return nil
}
