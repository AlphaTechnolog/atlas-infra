import { DynamoDBClient } from '@aws-sdk/client-dynamodb';
import { DynamoDBDocumentClient, PutCommand, DeleteCommand } from '@aws-sdk/lib-dynamodb';
import { DbRepository } from '../../domain/repositories/db-repository';

// TODO: Implement more methods like get, fetchAll, etc
export class DynamoRepository implements DbRepository {
  private _docClient: DynamoDBDocumentClient;
  private readonly _client: DynamoDBClient;
  private readonly _primaryKeyName: string;
  private readonly _tableName: string;

  constructor(tableName: string, primaryKeyName: string) {
    this._tableName = tableName;
    this._primaryKeyName = primaryKeyName;
    this._client = new DynamoDBClient({});
    this._docClient = DynamoDBDocumentClient.from(this._client);
  }

  async create<T extends Record<string, any>>(body: T): Promise<void> {
    await this._docClient.send(new PutCommand({
      TableName: this._tableName,
      Item: body,
    }));
  }

  async remove(pk: string): Promise<void> {
    await this._docClient.send(new DeleteCommand({
      TableName: this._tableName,
      Key: { [this._primaryKeyName]: pk },
    }));
  }
}