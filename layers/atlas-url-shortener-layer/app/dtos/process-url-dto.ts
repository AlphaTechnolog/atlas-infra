/**
 * @file This file defines the Data Transfer Objects (DTOs) for the `ProcessUrlUseCase`.
 */

import {ShortenedUrl} from "../../domain/entities/shortened-url";

/**
 * Input DTO for the `ProcessUrlUseCase`.
 */
export interface ProcessUrlInput {
  url: string;
}

/**
 * Output DTO for the ProcessUrlUseCase.
 */
export interface ProcessUrlOutput {
  message: string;
  shortenedURL: Omit<ShortenedUrl, 'createdAt'> & {
    createdAt: string;  // ISO 8601 string.
  };
}