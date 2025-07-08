import type { APIGatewayProxyWebsocketEventV2, APIGatewayProxyResult } from 'aws-lambda';
import type { VisitCountSubscriptionUseCase } from '/opt/nodejs/app/use-cases/visit-count-subscription-use-case';
import { container } from '/opt/nodejs/di/container';

const defaultHeaders = (): { [key: string]: string } => ({
  'Content-Type': 'application/json',
  'Access-Control-Allow-Origin': '*',
  'Access-Control-Allow-Methods': '*',
  'Access-Control-Allow-Headers': '*',
});

export const handler = async (event: APIGatewayProxyWebsocketEventV2): Promise<APIGatewayProxyResult> => {
  const useCase = container.resolve<VisitCountSubscriptionUseCase>('VisitCountSubscriptionUseCase');

  try {
    return await useCase.execute(event);
  } catch (err: any) {
    return {
      statusCode: 500,
      headers: defaultHeaders(),
      body: JSON.stringify({
        err: true,
        message: String(err),
      }),
    }
  }
}
