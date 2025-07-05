import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { AtlasLambdasResource } from "./atlas-lambdas-resource";
import { AtlasApiGatewayResource } from "./atlas-api-gateway-resource";
import {AtlasLayersResource} from "./atlas-layers-resource";

export class AtlasInfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const layers = new AtlasLayersResource(this, 'AtlasLayersResource');

    const lambdas = new AtlasLambdasResource(this, 'AtlasLambdasResource', {
      urlShortenerLayer: layers.urlShortenerLayer,
    });

    new AtlasApiGatewayResource(this, 'AtlasApiGatewayResource', {
      urlShortenerFunction: lambdas.urlShortener,
    });
  }
}
