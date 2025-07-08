import type { APIGatewayProxyResult, APIGatewayProxyWebsocketEventV2 } from "aws-lambda";
import type { DbRepository } from "../../domain/repositories/db-repository";

export class VisitCountSubscriptionUseCase {
  private readonly _dynamoWSConnsTable: DbRepository;

  constructor(dynamoWSConnsTable: DbRepository) {
    this._dynamoWSConnsTable = dynamoWSConnsTable;
  }

  private _sendResponse<T extends Record<string, any>>(status: number, body: T): APIGatewayProxyResult {
    return {
      statusCode: status,
      body: JSON.stringify(body),
    }
  }

  public async execute(event: APIGatewayProxyWebsocketEventV2): Promise<APIGatewayProxyResult> {
    const { requestContext: { connectionId, routeKey } } = event;

    if (!connectionId) {
      return this._sendResponse(400, {
        ok: true,
        message: "Missing connectionId",
      });
    }

    switch (routeKey) {
      case "$connect":
        await this._dynamoWSConnsTable.create({
          connectionId,
          timestamp: Date.now(),
        });
        console.log(`Connected: ${connectionId}`);
        break;
      case "$disconnect":
        await this._dynamoWSConnsTable.remove(connectionId);
        console.log(`Disconnected: ${connectionId}`);
        break;
      case "subscribe":
        console.log(`Client ${connectionId} sent a 'subscribe' message. Acknowledging.`);
        break;
      default:
        console.log(`Default route message from ${connectionId}: ${event.body}`);
        break;
    }

    return this._sendResponse(200, {
      ok: true,
      message: 'Ok',
    });
  }
}