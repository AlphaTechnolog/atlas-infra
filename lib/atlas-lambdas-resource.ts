import * as path from 'path'
import * as lambda from 'aws-cdk-lib/aws-lambda';
import * as lambdaEventSources from 'aws-cdk-lib/aws-lambda-event-sources';
import * as sqs from 'aws-cdk-lib/aws-sqs';
import * as dynamo from 'aws-cdk-lib/aws-dynamodb';
import { Construct } from 'constructs';

export type AtlasLambdasResourceProps = {
  urlShortenerLayer: lambda.LayerVersion,
  urlsSqsQueue: sqs.Queue,
  urlAnalyticsDynamoTable: dynamo.Table,
};

export class AtlasLambdasResource extends Construct {
  public urlShortener: lambda.Function;
  public urlListener: lambda.Function;
  public urlConsumer: lambda.Function;

  constructor(scope: Construct, id: string, props: AtlasLambdasResourceProps) {
    super(scope, id);

    const { urlShortenerLayer, urlsSqsQueue, urlAnalyticsDynamoTable } = props;

    this.urlShortener = new lambda.Function(this, 'AtlasUrlShortener', {
      functionName: 'atlas-url-shortener',
      description: 'Url shortener lambda function which emits converted urls.',
      handler: 'index.handler',
      layers: [urlShortenerLayer],
      code: lambda.Code.fromAsset(path.join(__dirname, '../lambdas/js/atlas-url-shortener/dist/')),
      runtime: lambda.Runtime.NODEJS_22_X,
      environment: {
        SHORTENED_URLS_SQS_URL: urlsSqsQueue.queueUrl,
      },
    });

    urlsSqsQueue.grantSendMessages(this.urlShortener);

    this.urlListener = new lambda.Function(this, 'AtlasUrlListener', {
      functionName: 'atlas-url-listener',
      description: 'Lambda which listens to emitted shortened urls and saves em in dynamodb',
      handler: 'bootstrap',
      code: lambda.Code.fromAsset(path.join(__dirname, '../lambdas/go/dist/atlas-url-listener.zip')),
      runtime: lambda.Runtime.PROVIDED_AL2023,
      architecture: lambda.Architecture.X86_64,
      loggingFormat: lambda.LoggingFormat.JSON,
      environment: {
        DYNAMO_URL_ANALYTICS_TABLE_NAME: urlAnalyticsDynamoTable.tableName,
      },
    });

    this.urlListener.addEventSource(
      new lambdaEventSources.SqsEventSource(urlsSqsQueue, {
        batchSize: 1,
        enabled: true,
      }),
    );

    urlAnalyticsDynamoTable.grantWriteData(this.urlListener);

    this.urlConsumer = new lambda.Function(this, 'AtlasUrlConsumer', {
      functionName: 'atlas-url-consumer',
      description: 'Lambda which returns the shortened url by its id and updates analytics based on usage',
      handler: 'bootstrap',
      code: lambda.Code.fromAsset(path.join(__dirname, '../lambdas/go/dist/atlas-url-consumer.zip')),
      runtime: lambda.Runtime.PROVIDED_AL2023,
      architecture: lambda.Architecture.X86_64,
      loggingFormat: lambda.LoggingFormat.JSON,
      environment: {
        DYNAMO_URL_ANALYTICS_TABLE_NAME: urlAnalyticsDynamoTable.tableName,
      },
    });

    urlAnalyticsDynamoTable.grantReadWriteData(this.urlConsumer);
  }
}
