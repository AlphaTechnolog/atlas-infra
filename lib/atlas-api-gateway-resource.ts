import * as apiGatewayV2 from 'aws-cdk-lib/aws-apigatewayv2';
import { aws_lambda, Duration } from "aws-cdk-lib";
import { HttpLambdaIntegration } from 'aws-cdk-lib/aws-apigatewayv2-integrations'
import { Construct } from 'constructs';

export interface ApiGatewayProps {
  urlShortenerFunction: aws_lambda.Function,
}

export class AtlasApiGatewayResource extends Construct {
  constructor(scope: Construct, id: string, props: ApiGatewayProps) {
    super(scope, id);

    const { urlShortenerFunction } = props;

    const api = new apiGatewayV2.HttpApi(this, 'AtlasApiGateway', {
      apiName: 'atlas-api-gateway',
      description: 'API Gateway for the atlas project',
      corsPreflight: {
        allowHeaders: ['Content-Type', 'Authorization', 'Accept'],
        allowMethods: [apiGatewayV2.CorsHttpMethod.POST],
        allowOrigins: ['*'],
        maxAge: Duration.days(1),
      },
    });

    const urlShortenerIntegration = new HttpLambdaIntegration('UrlShortenerIntegration', urlShortenerFunction);

    api.addRoutes({
      path: '/',
      methods: [apiGatewayV2.HttpMethod.POST],
      integration: urlShortenerIntegration,
    });
  }
}
