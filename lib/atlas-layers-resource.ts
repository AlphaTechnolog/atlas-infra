import * as path from 'path';
import * as cdk from 'aws-cdk-lib';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { Construct } from 'constructs';

export class AtlasLayersResource extends Construct {
  public urlShortenerLayer: lambda.LayerVersion;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    this.urlShortenerLayer = new lambda.LayerVersion(this, 'UrlShortenerLayer', {
      layerVersionName: 'atlas-url-shortener-layer',
      code: lambda.Code.fromAsset(path.join(__dirname, '../layers/atlas-url-shortener-layer/dist/layer.zip')),
      compatibleRuntimes: [lambda.Runtime.NODEJS_22_X],
      description: 'Layer which contains business logic for the url shortener lambda written in typescript',
      removalPolicy: cdk.RemovalPolicy.RETAIN,
    });
  }
}