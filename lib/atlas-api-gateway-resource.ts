import * as apiGatewayV2 from 'aws-cdk-lib/aws-apigatewayv2';
import * as iam from 'aws-cdk-lib/aws-iam';
import * as dynamo from 'aws-cdk-lib/aws-dynamodb';
import { aws_lambda, Duration } from "aws-cdk-lib";
import { HttpLambdaIntegration, WebSocketLambdaIntegration } from 'aws-cdk-lib/aws-apigatewayv2-integrations'
import { Construct } from 'constructs';

export interface ApiGatewayProps {
  urlShortenerFunction: aws_lambda.Function,
  urlConsumerFunction: aws_lambda.Function,
  urlVisitCountSubscriptionManager: aws_lambda.Function,
  urlVisitCountStreamProcessor: aws_lambda.Function,
  wsConnectionsTable: dynamo.Table,
}

export class AtlasApiGatewayResource extends Construct {
  constructor(scope: Construct, id: string, props: ApiGatewayProps) {
    super(scope, id);

    const {
      urlShortenerFunction,
      urlConsumerFunction,
      urlVisitCountSubscriptionManager,
      urlVisitCountStreamProcessor,
      wsConnectionsTable,
    } = props;

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
    const urlConsumerIntegration = new HttpLambdaIntegration('UrlConsumerIntegration', urlConsumerFunction);

    api.addRoutes({
      path: '/',
      methods: [apiGatewayV2.HttpMethod.POST],
      integration: urlShortenerIntegration,
    });

    api.addRoutes({
      path: '/consume',
      methods: [apiGatewayV2.HttpMethod.POST],
      integration: urlConsumerIntegration,
    });

    const webSocketApi = new apiGatewayV2.WebSocketApi(this, 'AtlasWebSocketApi', {
      apiName: 'atlas-websocket-api',
      description: 'Websocket API which will serve to notify about changes in data like visitCount',
      connectRouteOptions: {
        integration: new WebSocketLambdaIntegration('ConnectIntegration', urlVisitCountSubscriptionManager),
      },
      disconnectRouteOptions: {
        integration: new WebSocketLambdaIntegration('DisconnectIntegration', urlVisitCountSubscriptionManager),
      },
      defaultRouteOptions: {
        integration: new WebSocketLambdaIntegration('DefaultIntegration', urlVisitCountSubscriptionManager),
      },
    });

    webSocketApi.addRoute('subscribe', {
      integration: new WebSocketLambdaIntegration('SubscribeIntegration', urlVisitCountSubscriptionManager),
    });

    const webSocketStage = new apiGatewayV2.WebSocketStage(this, 'AtlasWebSocketStage', {
      webSocketApi,
      stageName: 'dev',
      autoDeploy: true,
    });

    urlVisitCountSubscriptionManager.addPermission('ApiGatewayInvoke', {
      principal: new iam.ServicePrincipal('apigateway.amazonaws.com'),
      action: 'lambda:InvokeFunction',
      sourceArn: webSocketApi.arnForExecuteApiV2(),
    });

    webSocketStage.grantManagementApiAccess(urlVisitCountSubscriptionManager);
    wsConnectionsTable.grantReadWriteData(urlVisitCountSubscriptionManager);
    wsConnectionsTable.grantReadData(urlVisitCountStreamProcessor);
    urlVisitCountStreamProcessor.addEnvironment('WEBSOCKET_ENDPOINT_URL', webSocketStage.url);
  }
}
