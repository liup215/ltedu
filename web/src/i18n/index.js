"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var vue_i18n_1 = require("vue-i18n");
var en_1 = require("../locales/en");
var zh_1 = require("../locales/zh");
var getDefaultLocale = function () {
    var saved = localStorage.getItem('locale');
    if (saved)
        return saved;
    var browser = navigator.language.toLowerCase();
    if (browser === 'zh-cn' || browser === 'zh')
        return 'zh';
    return 'en';
};
var i18n = (0, vue_i18n_1.createI18n)({
    legacy: false,
    locale: getDefaultLocale(),
    fallbackLocale: 'en',
    messages: {
        en: en_1.default,
        zh: zh_1.default,
    },
    globalInjection: true,
});
exports.default = i18n;
