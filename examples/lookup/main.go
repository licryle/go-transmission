package main

import (
	"log"

	"github.com/licryle/go-transmission"
)

func main() {
	// Note: This example assumes a running Transmission instance
	// Note: Transmission RPC is usually at /transmission/rpc
	client, err := transmission.New("http://localhost:9091/transmission/rpc", "", "")
	if err != nil {
		log.Fatalf("Failed to connect to Transmission: %v", err)
	}

	torrent, err := client.GetTorrent("da72e275823dba275b0b219dd52fg6b3a51fcd27")
	if err != nil {
		log.Panic(err)
	}

	log.Println("Torrent:")
	log.Println("   ID:            ", torrent.ID)
	log.Println("   Name:          ", torrent.Name)
	log.Println("   Status:        ", torrent.Status)
	log.Println("   LeftUntilDone: ", torrent.LeftUntilDone)
	log.Println("   Eta:           ", torrent.Eta)
	log.Println("   UploadRatio:   ", torrent.UploadRatio)
	log.Println("   RateDownload:  ", torrent.RateDownload)
	log.Println("   RateUpload:    ", torrent.RateUpload)
	log.Println("   DownloadDir:   ", torrent.DownloadDir)
	log.Println("   IsFinished:    ", torrent.IsFinished)
	log.Println("   PercentDone:   ", torrent.PercentDone)
	log.Println("   SeedRatioMode: ", torrent.SeedRatioMode)

	torrent2, err2 := client.GetTorrent("33")
	if err2 != nil {
		log.Panic(err2)
	}

	log.Println("Torrent:")
	log.Println("   ID:            ", torrent2.ID)
	log.Println("   Name:          ", torrent2.Name)
	log.Println("   Status:        ", torrent2.Status)
	log.Println("   LeftUntilDone: ", torrent2.LeftUntilDone)
	log.Println("   Eta:           ", torrent2.Eta)
	log.Println("   UploadRatio:   ", torrent2.UploadRatio)
	log.Println("   RateDownload:  ", torrent2.RateDownload)
	log.Println("   RateUpload:    ", torrent2.RateUpload)
	log.Println("   DownloadDir:   ", torrent2.DownloadDir)
	log.Println("   IsFinished:    ", torrent2.IsFinished)
	log.Println("   PercentDone:   ", torrent2.PercentDone)
	log.Println("   SeedRatioMode: ", torrent2.SeedRatioMode)

	torrent3, err3 := client.GetTorrent(33)
	if err3 != nil {
		log.Panic(err3)
	}

	log.Println("Torrent:")
	log.Println("   ID:            ", torrent3.ID)
	log.Println("   Name:          ", torrent3.Name)
	log.Println("   Status:        ", torrent3.Status)
	log.Println("   LeftUntilDone: ", torrent3.LeftUntilDone)
	log.Println("   Eta:           ", torrent3.Eta)
	log.Println("   UploadRatio:   ", torrent3.UploadRatio)
	log.Println("   RateDownload:  ", torrent3.RateDownload)
	log.Println("   RateUpload:    ", torrent3.RateUpload)
	log.Println("   DownloadDir:   ", torrent3.DownloadDir)
	log.Println("   IsFinished:    ", torrent3.IsFinished)
	log.Println("   PercentDone:   ", torrent3.PercentDone)
	log.Println("   SeedRatioMode: ", torrent3.SeedRatioMode)
}
