"use strict";
/// <reference types="../../../../../.npm/_npx/2db181330ea4b15b/node_modules/@vue/language-core/types/template-helpers.d.ts" />
/// <reference types="../../../../../.npm/_npx/2db181330ea4b15b/node_modules/@vue/language-core/types/props-fallback.d.ts" />
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
Object.defineProperty(exports, "__esModule", { value: true });
var notivue_1 = require("notivue");
var vue_1 = require("vue");
var vue_router_1 = require("vue-router");
var DonationPopup_vue_1 = require("./components/DonationPopup.vue");
var userStore_1 = require("./stores/userStore");
var donationPopupRef = (0, vue_1.ref)();
var router = (0, vue_router_1.useRouter)();
var userStore = (0, userStore_1.useUserStore)();
(0, vue_1.onMounted)(function () {
    router.afterEach(function () {
        var user = userStore === null || userStore === void 0 ? void 0 : userStore.user;
        var now = Date.now();
        var isVip = (user === null || user === void 0 ? void 0 : user.vipExpireAt) && new Date(user.vipExpireAt).getTime() > now;
        var isAdmin = (user === null || user === void 0 ? void 0 : user.isAdmin) === true;
        if (Math.random() < 0.05 && donationPopupRef.value && !isVip && !isAdmin && user) {
            donationPopupRef.value.show();
        }
    });
});
var __VLS_ctx = __assign(__assign({}, {}), {});
var __VLS_components;
var __VLS_intrinsics;
var __VLS_directives;
var __VLS_0;
/** @ts-ignore @type {typeof __VLS_components.Notivue | typeof __VLS_components.Notivue} */
notivue_1.Notivue;
// @ts-ignore
var __VLS_1 = __VLS_asFunctionalComponent1(__VLS_0, new __VLS_0({}));
var __VLS_2 = __VLS_1.apply(void 0, __spreadArray([{}], __VLS_functionalComponentArgsRest(__VLS_1), false));
{
    var __VLS_5 = __VLS_3.slots.default;
    var item = __VLS_vSlot(__VLS_5)[0];
    var __VLS_6 = void 0;
    /** @ts-ignore @type {typeof __VLS_components.Notification} */
    notivue_1.Notification;
    // @ts-ignore
    var __VLS_7 = __VLS_asFunctionalComponent1(__VLS_6, new __VLS_6({
        item: (item),
    }));
    var __VLS_8 = __VLS_7.apply(void 0, __spreadArray([{
            item: (item),
        }], __VLS_functionalComponentArgsRest(__VLS_7), false));
    __VLS_3.slots['' /* empty slot name completion */];
}
var __VLS_3;
var __VLS_11 = DonationPopup_vue_1.default;
// @ts-ignore
var __VLS_12 = __VLS_asFunctionalComponent1(__VLS_11, new __VLS_11({
    ref: "donationPopupRef",
}));
var __VLS_13 = __VLS_12.apply(void 0, __spreadArray([{
        ref: "donationPopupRef",
    }], __VLS_functionalComponentArgsRest(__VLS_12), false));
var __VLS_16 = {};
var __VLS_14;
var __VLS_18;
/** @ts-ignore @type {typeof __VLS_components.routerView | typeof __VLS_components.RouterView} */
routerView;
// @ts-ignore
var __VLS_19 = __VLS_asFunctionalComponent1(__VLS_18, new __VLS_18({}));
var __VLS_20 = __VLS_19.apply(void 0, __spreadArray([{}], __VLS_functionalComponentArgsRest(__VLS_19), false));
// @ts-ignore
var __VLS_17 = __VLS_16;
var __VLS_export = (await Promise.resolve().then(function () { return require('vue'); })).defineComponent({});
exports.default = {};
