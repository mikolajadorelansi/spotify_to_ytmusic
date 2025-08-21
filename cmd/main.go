package main

import (
    "log"
    "os"
    "spotify-to-ytmusic-go/internal/transfer"
    "spotify-to-ytmusic-go/config"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

    // Initialize transfer service
    transferService := transfer.NewTransferService(cfg)

    // Execute the transfer process
    err = transferService.TransferPlaylists()
    if err != nil {
        log.Fatalf("Error transferring playlists: %v", err)
    }

    log.Println("Transfer completed successfully.")
}