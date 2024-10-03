"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.urlPatternNoHtml = exports.urlPatternHtmlHttp = exports.urlPatternNoHtmlWithHttps = exports.urlPatternHtml = void 0;
exports.urlPatternHtml = /^https?:\/\/(?:www\.)?([a-zA-Z0-9.-]*shop[a-zA-Z0-9.-]*)\/.*\.html$/;
exports.urlPatternNoHtmlWithHttps = /^https?:\/\/(?:www\.)?([a-zA-Z0-9.-]*shop[a-zA-Z0-9.-]*)\/.*/;
exports.urlPatternHtmlHttp = /^([a-zA-Z0-9.-]*shop[a-zA-Z0-9.-]*)\/.*\.html$/;
exports.urlPatternNoHtml = /^([a-zA-Z0-9.-]*shop[a-zA-Z0-9.-]*)\/.*/;
