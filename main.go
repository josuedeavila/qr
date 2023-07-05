package main

import (
	"fmt"
	"os"

	"github.com/josuerosadeavila/qr/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var filePath string
	var bucket string
	var endpoint string

	rootCmd := &cobra.Command{
		Use:   "s3downloader",
		Short: "A command-line tool to download files from Amazon S3",
		Run:   cmd.ShowQRCode,
	}

	rootCmd.Flags().StringVarP(&filePath, "path", "p", "", "File path in S3")
	rootCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "S3 bucket name")
	rootCmd.Flags().StringVar(&endpoint, "endpoint", "", "S3 endpoint URL")

	rootCmd.MarkFlagRequired("path")
	rootCmd.MarkFlagRequired("bucket")
	rootCmd.MarkFlagRequired("endpoint")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
