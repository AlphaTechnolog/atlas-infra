import * as cdk from 'aws-cdk-lib';
import * as dynamodb from 'aws-cdk-lib/aws-dynamodb';
import { Construct } from 'constructs';

export class AtlasWsConnsTableResource extends Construct {
  public wsConnectionsTable: dynamodb.Table;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.wsConnectionsTable = new dynamodb.Table(this, 'AtlasWSConnectionsTable', {
      tableName: 'atlas-ws-connections-table',
      partitionKey: { name: 'connectionId', type: dynamodb.AttributeType.STRING },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    });
  }
}