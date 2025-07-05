/**
 * @file This file defines the Data Transfer Objects (DTOs) for the `ProcessUrlUseCase`.
 */

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
  processedUrl: string;
  message: string;
  timestamp: string;  // ISO 8601 string.
}