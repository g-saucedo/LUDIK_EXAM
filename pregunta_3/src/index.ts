import { UrlProcessor } from "./services/urlProcessor";

const filePath = 'data/urls.txt';
const processor = new UrlProcessor(filePath);

processor.process().catch(err => {
    console.error(`Error al procesar las urls: ${err.message}`);
})