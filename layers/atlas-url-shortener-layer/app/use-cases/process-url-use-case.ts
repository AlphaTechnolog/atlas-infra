/**
 * @file This file defines the `ProcessUrlUseCase`. This use case will receive an url, shortens it via crypto,
 * then will emit it into the stream, and will allow other micros to take analytics for it.
 */
import { ShortenedUrl } from "../../domain/entities/shortened-url";
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

    // TODO: Emit this to kafka stream.
    const shortenedURL = ShortenedUrl.create(url);
    console.log("[ProcessUrlUseCase::execute()] Created new shortened url", { shortenedURL });

    return {
      message: "Shortened url has been published successfully",
      shortenedURL: {
        ...shortenedURL,
        createdAt: shortenedURL.createdAt.toISOString(),
      },
    };
  }
}