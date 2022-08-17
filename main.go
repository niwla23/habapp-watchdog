// https://openhab.b49.cloudserver.click/rest/items/HABApp_LastRulePing/state

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func restartContainerByName(containerName string) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if contains(container.Names, "/"+containerName) {
			if err := cli.ContainerRestart(ctx, container.ID, nil); err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	CONTAINER_NAME := os.Getenv("HABAPP_CONTAINER_NAME")
	OPENHAB_LAST_PING_ITEM_NAME := os.Getenv("OPENHAB_LAST_PING_ITEM_NAME")
	OPENHAB_REST_BASE_URL := os.Getenv("OPENHAB_REST_BASE_URL")

	CHECK_INTERVAL_SECONDS, err := strconv.ParseInt(os.Getenv("CHECK_INTERVAL_SECONDS"), 10, 0)
	if err != nil {
		panic(err)
	}
	HABAPP_MAX_PING_DELAY_SECONDS, err := strconv.ParseInt(os.Getenv("HABAPP_MAX_PING_DELAY_SECONDS"), 10, 0)
	if err != nil {
		panic(err)
	}
	HABAPP_POST_RESTART_DELAY_SECONDS, err := strconv.ParseInt(os.Getenv("HABAPP_POST_RESTART_DELAY_SECONDS"), 10, 0)
	if err != nil {
		panic(err)
	}
	for {
		resp, err := http.Get(fmt.Sprintf("%s/items/%s/state", OPENHAB_REST_BASE_URL, OPENHAB_LAST_PING_ITEM_NAME))
		if err != nil {
			fmt.Println("error fetching data")
			panic(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error reading body")
			panic(err)
		}
		timestamp, err := strconv.ParseInt(string(body), 10, 0)
		if err != nil {
			fmt.Println("error parsing data")
			restartContainerByName(CONTAINER_NAME)
			fmt.Println("Ping value is NULL, restarting HABApp")
			time.Sleep(time.Second * time.Duration(HABAPP_POST_RESTART_DELAY_SECONDS))
			continue
		}

		diff := time.Now().Unix() - timestamp
		if diff > HABAPP_MAX_PING_DELAY_SECONDS {
			fmt.Printf("HABApp does not react, restarting it, last heartbeat was %d seconds ago\n", diff)
			restartContainerByName(CONTAINER_NAME)
			time.Sleep(time.Second * time.Duration(HABAPP_POST_RESTART_DELAY_SECONDS))
		}

		time.Sleep(time.Second * time.Duration(CHECK_INTERVAL_SECONDS))
	}
}
