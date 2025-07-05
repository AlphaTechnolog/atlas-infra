import type { APIGatewayProxyEvent, APIGatewayProxyResult } from "aws-lambda";
import { container } from '/opt/nodejs/di/container';
import { ProcessUrlUseCase } from "/opt/nodejs/app/use-cases/process-url-use-case";

type Body = {
  url: string;
};

const decodeEvent = (ev: APIGatewayProxyEvent): Body | null => {
  let raw = ev.body;
  if (!raw) return null;
  if (ev.isBase64Encoded) raw = Buffer.from(raw, "base64").toString("ascii");
  const value = JSON.parse(raw) as { url?: string };
  if (!Boolean(value.url) || value.url?.trim() === "") return null;
  return { url: value.url! };
}

const defaultHeaders = (): { [key: string]: string } => ({
  'Content-Type': 'application/json',
  'Access-Control-Allow-Origin': '*',
  'Access-Control-Allow-Methods': '*',
  'Access-Control-Allow-Headers': '*',
});

export const handler = async (event: APIGatewayProxyEvent): Promise<APIGatewayProxyResult> => {
  const body = decodeEvent(event);
  if (!body) {
    return {
      statusCode: 400,
      headers: defaultHeaders(),
      body: JSON.stringify({
        err: true,
        message: "Invalid given body",
      }),
    }
  }

  const { url } = body;
  const useCase = container.resolve<ProcessUrlUseCase>("ProcessUrlUseCase");
  console.log("Resolved useCase?", Boolean(useCase));

  const output = await useCase.execute({ url });

  return {
    statusCode: 200,
    headers: defaultHeaders(),
    body: JSON.stringify(output),
  };
}