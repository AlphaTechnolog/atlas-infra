import { Construct } from 'constructs';
import * as cdk from 'aws-cdk-lib';
import * as dynamo from 'aws-cdk-lib/aws-dynamodb';

export class AtlasDynamoUrlsResource extends Construct {
  public urlTable: dynamo.Table;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.urlTable = new dynamo.Table(this, 'AtlasUrlAnalyticsTable', {
      tableName: 'atlas-url-analytics',
      billingMode: dynamo.BillingMode.PAY_PER_REQUEST,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      stream: dynamo.StreamViewType.NEW_AND_OLD_IMAGES,
      partitionKey: {
        name: 'id',
        type: dynamo.AttributeType.STRING,
      },
    });
  }
}