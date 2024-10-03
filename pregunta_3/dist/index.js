"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const urlProcessor_1 = require("./services/urlProcessor");
const filePath = 'data/urls.txt';
const processor = new urlProcessor_1.UrlProcessor(filePath);
processor.process().catch(err => {
    console.error(`Error al procesar las urls: ${err.message}`);
});
