import path from 'path'
import * as lambda from 'aws-cdk-lib/aws-lambda';
import { Construct } from 'constructs';

export type AtlasLambdasResourceProps = {
  urlShortenerLayer: lambda.LayerVersion,
};

export class AtlasLambdasResource extends Construct {
  public urlShortener: lambda.Function;

  constructor(scope: Construct, id: string, props: AtlasLambdasResourceProps) {
    super(scope, id);

    const { urlShortenerLayer } = props;

    this.urlShortener = new lambda.Function(this, 'AtlasUrlShortener', {
      functionName: 'atlas-url-shortener',
      description: 'Url shortener lambda function which emits converted urls.',
      handler: 'index.handler',
      layers: [urlShortenerLayer],
      code: lambda.Code.fromAsset(path.join(__dirname, '../lambdas/atlas-url-shortener/dist/')),
      runtime: lambda.Runtime.NODEJS_22_X,
    });
  }
}
