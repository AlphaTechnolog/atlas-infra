import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { AtlasLambdasResource } from "./atlas-lambdas-resource";
import { AtlasApiGatewayResource } from "./atlas-api-gateway-resource";

export class AtlasInfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const lambdas = new AtlasLambdasResource(this, 'AtlasLambdasResource');

    new AtlasApiGatewayResource(this, 'AtlasApiGatewayResource', {
      urlShortenerFunction: lambdas.urlShortener,
    });
  }
}
