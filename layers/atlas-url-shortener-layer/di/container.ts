import { ProcessUrlUseCase } from '../app/use-cases/process-url-use-case';
import { IUrlRepository } from "../domain/repositories/url-repository";
import { UrlSQSPublisher } from "../infrastructure/sqs/url-publisher";
import {DbRepository} from "../domain/repositories/db-repository";
import {DynamoRepository} from "../infrastructure/dynamo/dynamo-repository";
import {VisitCountSubscriptionUseCase} from "../app/use-cases/visit-count-subscription-use-case";

type DependencyMap = { [key: string]: any };

class Container {
  private dependencies: DependencyMap = {};

  /**
   * Registers a new dependency with the container.
   * @param name The name of the dependency (e.g.: `IUrlRepository`).
   * @param dependency The instance or function that returns the instance of the dependency.
   */
  public register<T>(name: string, dependency: T): void {
    this.dependencies[name] = dependency;
  }

  /**
   * Resolves a dependency from the container.
   * @param name The name of the dependency to resolve.
   * @returns The resolved dependency instance.
   * @throws Error if the dependency is not found.
   */
  public resolve<T>(name: string): T {
    if (!this.dependencies[name]) {
      throw new Error(`Dependency ${name} not found in container.`);
    }
    if (typeof this.dependencies[name] === 'function') {
      return this.dependencies[name]();
    }
    return this.dependencies[name];
  }
}

export const container = new Container();

const urlSQSUrl = process.env.SHORTENED_URLS_SQS_URL!;
const wsConnsTableName = process.env.WEBSOCKET_CONNECTIONS_TABLE_NAME!;

container.register('UrlSQSPublisher', (): IUrlRepository => new UrlSQSPublisher(urlSQSUrl));

container.register('DynamoWSConnsTable', (): DbRepository =>
  new DynamoRepository(wsConnsTableName, 'connectionId'));

container.register('ProcessUrlUseCase', () =>
  new ProcessUrlUseCase(container.resolve<IUrlRepository>('UrlSQSPublisher')));

container.register('VisitCountSubscriptionUseCase', () =>
  new VisitCountSubscriptionUseCase(container.resolve<DbRepository>('DynamoWSConnsTable')));
