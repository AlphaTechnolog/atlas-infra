package appconfig

import "os"

var DynamoUrlAnalyticsTableName = os.Getenv("DYNAMO_URL_ANALYTICS_TABLE_NAME")
