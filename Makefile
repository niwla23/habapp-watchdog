dev:
	HABAPP_CONTAINER_NAME=habapp CHECK_INTERVAL_SECONDS=10 HABAPP_MAX_PING_DELAY_SECONDS=60 HABAPP_POST_RESTART_DELAY_SECONDS=60 go run main.go