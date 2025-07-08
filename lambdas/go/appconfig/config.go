package appconfig

import "os"

var DynamoUrlAnalyticsTableName = os.Getenv("DYNAMO_URL_ANALYTICS_TABLE_NAME")
var WebSocketEndpoint = os.Getenv("WEBSOCKET_ENDPOINT_URL")
var WSConnectionsTableName = os.Getenv("WEBSOCKET_CONNECTIONS_TABLE_NAME")
