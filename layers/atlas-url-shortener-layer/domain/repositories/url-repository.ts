/**
 * @file This file defines an interface for the URL repository.
 * This interface belongs to the Domain layer, ensuring that the Domain does not depend on any specific data
 * persistence implementation. Any concrete repository must implement this interface.
 */
import { ShortenedUrl } from '../entities/shortened-url';

export interface IUrlRepository {
  save(url: ShortenedUrl): Promise<void>;
}