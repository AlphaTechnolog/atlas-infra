export interface DbRepository {
  create<T extends Record<string, any>>(body: T): Promise<void>;
  remove(pk: string): Promise<void>;
}