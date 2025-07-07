import * as crypto from "crypto";

/**
 * @file This file defines the `ShortenedUrl` entity. The `ShortenedUrl` entity represents the core business
 * object for a shortened url. It encapsulates the properties and behavior of a shortened URL. This entity belongs
 * to the Domain layer and should have no external dependencies.
 */

export const DEFAULT_SHORT_CODE_LENGTH = 6;

export class ShortenedUrl {
  public readonly id: string;  // The unique id for the url, e.g: the short code.
  public readonly originalUrl: string;  // The original long url.
  public readonly createdAt: Date;  // Date of record creation.

  /**
   * Creates a new ShortenedUrl interface.
   * @param id
   * @param originalUrl
   * @param createdAt
   */
  constructor(id: string, originalUrl: string, createdAt: Date) {
    if (!id || id.trim() === '') {
      throw new Error('ShortenedURL id cannot be empty');
    }
    if (!originalUrl || originalUrl.trim() === '') {
      throw new Error('ShortenedUrl original url cannot be empty.');
    }
    try {
      new URL(originalUrl);
    } catch (e) {
      // TODO: Maybe throw
      console.warn(`Provided URL "${originalUrl}" is not a valid URL format.`);
    }
    if (isNaN(createdAt.getTime())) {
      throw new Error('Created at must be a valid Date object.');
    }

    this.id = id;
    this.originalUrl = originalUrl;
    this.createdAt = createdAt;
  }

  /**
   * Generates a random alphanumeric short code.
   * @param originalUrl
   * @param shortCodeLength
   * @returns ShortenedUrl
   */
  public static create(originalUrl: string, shortCodeLength: number = DEFAULT_SHORT_CODE_LENGTH) {
    const id = this.generateShortCode(shortCodeLength);
    const createdAt = new Date();
    return new ShortenedUrl(id, originalUrl, createdAt);
  }

  /**
   * Generates a unique short code given a max length.
   * @param length The max length of the code
   * @private
   */
  private static generateShortCode(length: number): string {
    const bytesNeeded = Math.ceil(length / 2);
    const randomBytes = crypto.randomBytes(bytesNeeded);
    let code = randomBytes.toString("hex");
    if (code.length > length) code = code.substring(0, length);
    return code;
  }

  public displayable(): { id: string; originalUrl: string; createdAt: string; } {
    return {
      id: this.id,
      originalUrl: this.originalUrl,
      createdAt: this.createdAt.toISOString(),
    };
  }
}