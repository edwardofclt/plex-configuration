package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/ryanwholey/go-pihole"
)

var piHoleClient *pihole.Client

const sleepTime = 5 * time.Second

func init() {
	godotenv.Load(".env")
	var err error
	piHoleClient, err = pihole.New(pihole.Config{
		BaseURL:  fmt.Sprintf("https://%s", os.Getenv("PIHOLE_HOST")),
		Password: os.Getenv("PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Error creating pihole client: %s", err.Error())
	}
}

func main() {
	ctx := context.Background()

	for {
		runningContainers, err := getContainerDNS(ctx)
		if err != nil {
			logrus.Fatalf("Error getting running containers: %s", err.Error())
		}

		localDnsEntries, err := getPiholeLocalDns(ctx)
		if err != nil {
			logrus.Fatalf("Error getting local DNS entries: %s", err.Error())
		}

		// compare the running containers with the local DNS entries
		for _, container := range runningContainers {
			found := false
		LOCALDNS:
			for _, entry := range localDnsEntries {
				if container == entry {
					found = true
					break LOCALDNS
				}
			}
			if !found {
				err := createDnsEntry(ctx, container)
				if err != nil {
					logrus.Errorf("Error creating DNS entry for %s: %s", container, err.Error())
				}
			}
		}

		time.Sleep(sleepTime)
	}
}

func createDnsEntry(ctx context.Context, container string) error {
	_, err := piHoleClient.LocalDNS.Create(ctx, container, os.Getenv("TARGET_IP"))
	if err != nil {
		return err
	}
	logrus.Infof("Added %s to local DNS", container)
	return nil
}

func getPiholeLocalDns(ctx context.Context) ([]string, error) {
	records, err := piHoleClient.LocalDNS.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	dnsEntries := []string{}

	for _, record := range records {
		dnsEntries = append(dnsEntries, record.Domain)
	}

	return dnsEntries, nil
}

func getContainerDNS(ctx context.Context) ([]string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		logrus.Errorf("Error listing containers: %s", err.Error())
		return nil, err
	}

	dnsEntries := []string{}

	for _, container := range containers {
		if caddyValue, ok := container.Labels["caddy"]; ok {
			dnsEntries = append(dnsEntries, caddyValue)
		}
	}

	return dnsEntries, nil
}
