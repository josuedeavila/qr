package cmd

import (
	"fmt"
	"os"

	"github.com/josuerosadeavila/qr/s3"
	"github.com/mdp/qrterminal"
	"github.com/spf13/cobra"
)

func ShowQRCode(cmd *cobra.Command, args []string) {
	filePath, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")
	endpoint, _ := cmd.Flags().GetString("endpoint")

	if filePath == "" {
		fmt.Println("please provide a file path using the --path flag")
		return
	}

	client := s3.NewClient(&s3.Config{
		Bucket:   bucket,
		Endpoint: endpoint,
		Region:   "us-east-1",
	})

	code, err := client.Get(filePath)
	if err != nil {
		fmt.Println("error downloading file from S3:", err)
		return
	}

	qrterminal.GenerateHalfBlock(string(code), qrterminal.L, os.Stdout)

}
