package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

// DNSOperation represents the operation to be performed on DNS
type DNSOperation struct {
	Action string
	IP     string
}

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "List, add or delete DNS",
	Long: `dns -- list all DNS entries
dns -a [DNS_IP] -- add DNS_IP to the list
dns -d [DNS_IP] -- delete DNS_IP from the list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		add, _ := cmd.Flags().GetBool("add")
		del, _ := cmd.Flags().GetBool("delete")

		if add {
			if len(args) != 1 {
				return fmt.Errorf("add operation requires exactly one DNS IP")
			}
			err := modifyDNS(DNSOperation{Action: "add", IP: args[0]})
			if err != nil {
				return err
			}
			fmt.Println("DNS added successfully")
		} else if del {
			if len(args) != 1 {
				return fmt.Errorf("delete operation requires exactly one DNS IP")
			}
			err := modifyDNS(DNSOperation{Action: "delete", IP: args[0]})
			if err != nil {
				return err
			}
			fmt.Println("DNS deleted successfully")
		} else {
			dnsList, err := listDNS()
			if err != nil {
				return err
			}
			fmt.Println("DNS List:")
			for _, dns := range dnsList {
				fmt.Println(dns)
			}
		}

		return nil
	},
}

func listDNS() ([]string, error) {
	resp, err := http.Get("http://example.com/api/list-dns")
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var dnsList []string
	err = json.Unmarshal(body, &dnsList)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return dnsList, nil
}

func modifyDNS(operation DNSOperation) error {
	payload, err := json.Marshal(operation)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	resp, err := http.Post("http://example.com/api/modify-dns", "application/json", bytes.NewBuffer(payload))
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
	rootCmd.AddCommand(dnsCmd)

	dnsCmd.Flags().BoolP("add", "a", false, "Add a DNS entry")
	dnsCmd.Flags().BoolP("delete", "d", false, "Delete a DNS entry")
}
