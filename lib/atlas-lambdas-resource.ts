import path from 'path'
import * as cdk from 'aws-cdk-lib';
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { Construct } from 'constructs';

export class AtlasLambdasResource extends Construct {
  public urlShortener: lambda.Function;

  constructor(scope: Construct, id: string) {
    super(scope, id);

    const layerExportID = 'atlas-layer-deployment-stack-UrlShortenerLayerVersionArn';
    const urlShortenerLayer = lambda.LayerVersion.fromLayerVersionArn(
      this,
      'UrlShortenerLayer',
      cdk.Fn.importValue(layerExportID),
    );

    const lambdaPath = path.join(
      __dirname,
      '../lambdas/atlas-url-shortener/dist/atlas-infra/lambdas/atlas-url-shortener/src',  // lmao
    );

    this.urlShortener = new lambda.Function(this, 'AtlasUrlShortener', {
      functionName: 'atlas-url-shortener',
      description: 'Url shortener lambda function which emits converted urls.',
      handler: 'index.handler',
      layers: [urlShortenerLayer],
      code: lambda.Code.fromAsset(lambdaPath),
      runtime: lambda.Runtime.NODEJS_22_X,
    });
  }
}
