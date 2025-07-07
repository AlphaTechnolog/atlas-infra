import { SQSClient, SendMessageCommand } from '@aws-sdk/client-sqs';
import { IUrlRepository } from '../../domain/repositories/url-repository';
import { ShortenedUrl } from "../../domain/entities/shortened-url";

export class UrlSQSPublisher implements IUrlRepository {
  private readonly _queueUrl: string;

  constructor(queueUrl: string) {
    this._queueUrl = queueUrl;
  }

  async save(url: ShortenedUrl) {
    const client = new SQSClient();

    const body = {
      id: url.id,
      url: url.originalUrl,
      createdAt: url.createdAt.toISOString(),
    };

    await client.send(new SendMessageCommand({
      QueueUrl: this._queueUrl,
      MessageBody: JSON.stringify(body),
    })).catch(err => {
      console.error("Unable to send message to sqs at", this._queueUrl, { err });
    });

    console.log("Sent message successfully to sqs", body);
  }
}