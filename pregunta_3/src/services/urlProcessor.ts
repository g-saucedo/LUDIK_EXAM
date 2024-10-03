import * as fs from 'fs';
import * as readline from 'readline';
import { urlPatternHtmlHttp, urlPatternHtml, urlPatternNoHtml, urlPatternNoHtmlWithHttps } from '../utils/patterns';

export class UrlProcessor {
  private filePath: string;

  constructor(filePath: string) {
    this.filePath = filePath;
  }


  async process(): Promise<void> {
    const fileStream = fs.createReadStream(this.filePath);

    const rl = readline.createInterface({
      input: fileStream,
      crlfDelay: Infinity
    });

    const urlSet = new Set<string>();

    for await (const line of rl) {
      const trimmedLine = line.trim();
      if (urlPatternHtml.test(trimmedLine)
        || urlPatternHtmlHttp.test(trimmedLine)
        || urlPatternNoHtmlWithHttps.test(trimmedLine)
        || urlPatternNoHtml.test(trimmedLine)) {
        urlSet.add(trimmedLine);
      }
    }

    console.log(`Numero total de URLs que cumplen con los criterios ${urlSet.size}`);
    console.log("Urls encontradas: ");
    for (const url of urlSet) {
      console.log(url);
    }
  }
}