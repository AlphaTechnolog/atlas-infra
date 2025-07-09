import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { AtlasLambdasResource } from "./atlas-lambdas-resource";
import { AtlasApiGatewayResource } from "./atlas-api-gateway-resource";
import { AtlasLayersResource } from "./atlas-layers-resource";
import { AtlasUrlSqsResource } from "./atlas-url-sqs-resource";
import { AtlasDynamoUrlsResource } from "./atlas-dynamo-urls-resource";
import { AtlasWsConnsTableResource } from "./atlas-ws-conns-table-resource";

export class AtlasInfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const urlsSqsQueue = new AtlasUrlSqsResource(this, 'AtlasUrlSqsResource');
    const urlAnalyticsDynamoTable = new AtlasDynamoUrlsResource(this, 'AtlasDynamoUrlsResource');
    const wsConnectionsTable = new AtlasWsConnsTableResource(this, 'AtlasWSConnsTableResource');
    const layers = new AtlasLayersResource(this, 'AtlasLayersResource');

    const lambdas = new AtlasLambdasResource(this, 'AtlasLambdasResource', {
      urlShortenerLayer: layers.urlShortenerLayer,
      urlsSqsQueue: urlsSqsQueue.urlSqsQueue,
      urlAnalyticsDynamoTable: urlAnalyticsDynamoTable.urlTable,
      wsConnectionsTable: wsConnectionsTable.wsConnectionsTable,
    });

    new AtlasApiGatewayResource(this, 'AtlasApiGatewayResource', {
      urlFetcherFunction: lambdas.urlFetcher,
      urlShortenerFunction: lambdas.urlShortener,
      urlConsumerFunction: lambdas.urlConsumer,
      urlVisitCountStreamProcessor: lambdas.urlVisitCountStreamProcessor,
      urlVisitCountSubscriptionManager: lambdas.urlVisitCountSubscriptionManager,
      wsConnectionsTable: wsConnectionsTable.wsConnectionsTable,
    });
  }
}
