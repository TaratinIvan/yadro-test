package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// hostCmd represents the host command
var hostCmd = &cobra.Command{
	Use:   "host [NEW_NAME]",
	Short: "Change host name of service",
	Long:  `host [NEW_NAME] -- change current host name to NEW_NAME`,
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
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
	// Create JSON payload
	payload, err := json.Marshal(map[string]string{
		"hostname": newName,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Send HTTP POST request
	resp, err := http.Post("http://example.com/api/change-hostname", "application/json", bytes.NewBuffer(payload))
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

	// Define your flags and configuration settings here.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
