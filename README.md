# habapp watchdog

Restarts habapp if it does not react anymore

Environment:
```yaml
HABAPP_CONTAINER_NAME: "habapp" # Name of the habapp container
OPENHAB_LAST_PING_ITEM_NAME: HABApp_LastRulePing # Name of the item that contains the last time the ping rule was executed
OPENHAB_REST_BASE_URL: "http://192.168.178.33:8080/rest"
CHECK_INTERVAL_SECONDS: "10"
HABAPP_MAX_PING_DELAY_SECONDS: "60" # Maximum delay between the last ping and the current time
HABAPP_POST_RESTART_DELAY_SECONDS: "60" # Delay after restarting the habapp container
```

Mount the docker socket into the container

add the `ruleping.py` to habapp and create the HABApp_LastRulePing

## Known Bugs
It does not accept LetsEncrypt certificates, no idea why