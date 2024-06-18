package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var hostCmd = &cobra.Command{
	Use:   "host [NEW_NAME]",
	Short: "Change host name of service",
	Long:  `host [NEW_NAME] -- change current host name to NEW_NAME`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newName := args[0]
		err := changeHostName(newName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Host name changed successfully")
	},
}

func changeHostName(newName string) error {

	payload, err := json.Marshal(map[string]string{
		"hostname": newName,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/api/change-hostname", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(hostCmd)
}
