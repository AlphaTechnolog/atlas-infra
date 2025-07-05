/**
 * @file This file defines the `ProcessUrlUseCase`. This use case will receive an url, shortens it via crypto,
 * then will emit it into the stream, and will allow other micros to take analytics for it.
 */
import type { ProcessUrlInput, ProcessUrlOutput } from "../dtos/process-url-dto";

export class ProcessUrlUseCase {
  constructor() {
    // No dependencies for this use case.
  }

  /**
   * Executes the use case to process a url
   * @param url The url to process.
   * @returns A promise that resolves to the output information defined by the `ProcessUrlOutput` DTO.
   */
  public async execute({ url }: ProcessUrlInput): Promise<ProcessUrlOutput> {
    if (!url || url.trim() === "") {
      throw new Error("URL is required");
    }

    console.log("[ShortenedLayer::ProcessUrlUseCase.execute()] Received url", { url });

    // TODO: Actually call the ShortenedUrl domain and then the Kafka publisher to send an actual response.
    return {
      processedUrl: url,
      message: "Shortened url has been published successfully",
      timestamp: new Date().toISOString(),
    }
  }
}