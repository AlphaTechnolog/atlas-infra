import type { APIGatewayProxyEvent, APIGatewayProxyResult } from "aws-lambda";

export const handler = async (event: APIGatewayProxyEvent): Promise<APIGatewayProxyResult> => {
  console.log("event is", { event });
  return Promise.resolve({
    statusCode: 200,
    body: JSON.stringify({ hello: "world" }),
  });
}