import { Construct } from 'constructs';
import * as cdk from "aws-cdk-lib";
import * as sqs from 'aws-cdk-lib/aws-sqs';

export class AtlasUrlSqsResource extends Construct {
  public urlSqsQueue: sqs.Queue;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.urlSqsQueue = new sqs.Queue(this, 'AtlasUrlQueue', {
      queueName: 'atlas-url-notification-queue',
      visibilityTimeout: cdk.Duration.seconds(300),
      deadLetterQueue: {
        maxReceiveCount: 3,
        queue: new sqs.Queue(this, 'AtlasUrlDql', {
          queueName: 'atlas-url-dlq',
        }),
      },
    });
  }
}