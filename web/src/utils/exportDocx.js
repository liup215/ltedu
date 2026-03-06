"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g = Object.create((typeof Iterator === "function" ? Iterator : Object).prototype);
    return g.next = verb(0), g["throw"] = verb(1), g["return"] = verb(2), typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.exportExamPaperToDocx = exportExamPaperToDocx;
var docx_1 = require("docx");
/**
 * Export exam paper to docx format
 * @param paper ExamPaper object with name and questions
 */
function exportExamPaperToDocx(paper) {
    return __awaiter(this, void 0, void 0, function () {
        var paragraphs, timestamp, title, questions, _loop_1, i, doc, blob, url, a, error_1;
        var _a, _b, _c, _d, _e, _f, _g;
        return __generator(this, function (_h) {
            switch (_h.label) {
                case 0:
                    paragraphs = [];
                    timestamp = new Date().toISOString().replace(/[-:T]/g, '').slice(0, 14);
                    title = paper.name || 'Exam Paper';
                    // Add title
                    paragraphs.push(new docx_1.Paragraph({
                        children: [
                            new docx_1.TextRun({
                                text: title,
                                bold: true,
                                size: 32,
                                font: "Arial"
                            }),
                            new docx_1.TextRun({
                                text: ' ',
                                size: 32,
                                font: "Arial"
                            })
                        ],
                        spacing: { after: 400 },
                        alignment: "center"
                    }));
                    questions = paper.questions || [];
                    _loop_1 = function (i) {
                        var parser, doc_1, runs_1, _i, _j, node, text, el, pRuns, _k, _l, child, text, childEl, img, response, blob, buffer, imgWidth, imgHeight, maxWidth, maxHeight, finalWidth, finalHeight, _m, text, text, katexHtml, formulaText, text, img, response, blob, buffer, imgWidth, imgHeight, maxWidth, maxHeight, finalWidth, finalHeight, _o;
                        return __generator(this, function (_p) {
                            switch (_p.label) {
                                case 0:
                                    // Add question number
                                    paragraphs.push(new docx_1.Paragraph({
                                        children: [
                                            new docx_1.TextRun({
                                                text: "".concat(i + 1, "."),
                                                bold: true,
                                                size: 24,
                                                font: "Arial"
                                            }),
                                            new docx_1.TextRun({
                                                text: ' ',
                                                size: 24,
                                                font: "Arial"
                                            })
                                        ],
                                        alignment: "left"
                                    }));
                                    if (!questions[i].stem) return [3 /*break*/, 24];
                                    parser = new DOMParser();
                                    doc_1 = parser.parseFromString(questions[i].stem, 'text/html');
                                    runs_1 = [];
                                    _i = 0, _j = Array.from(doc_1.body.childNodes);
                                    _p.label = 1;
                                case 1:
                                    if (!(_i < _j.length)) return [3 /*break*/, 23];
                                    node = _j[_i];
                                    if (!(node.nodeType === Node.TEXT_NODE)) return [3 /*break*/, 2];
                                    text = (_a = node.textContent) === null || _a === void 0 ? void 0 : _a.trim();
                                    if (text) {
                                        runs_1.push(new docx_1.TextRun({
                                            text: text,
                                            size: 24,
                                            font: "Arial"
                                        }));
                                    }
                                    return [3 /*break*/, 22];
                                case 2:
                                    if (!(node.nodeType === Node.ELEMENT_NODE)) return [3 /*break*/, 22];
                                    el = node;
                                    if (!(el.tagName === 'P')) return [3 /*break*/, 14];
                                    pRuns = [];
                                    _k = 0, _l = Array.from(el.childNodes);
                                    _p.label = 3;
                                case 3:
                                    if (!(_k < _l.length)) return [3 /*break*/, 13];
                                    child = _l[_k];
                                    if (!(child.nodeType === Node.TEXT_NODE)) return [3 /*break*/, 4];
                                    text = (_b = child.textContent) === null || _b === void 0 ? void 0 : _b.trim();
                                    if (text) {
                                        pRuns.push(new docx_1.TextRun({
                                            text: text.replace(/\s+/g, ' '), // Normalize whitespace
                                            size: 24,
                                            font: "Arial",
                                            break: text.endsWith('\n') ? 1 : undefined
                                        }));
                                    }
                                    return [3 /*break*/, 12];
                                case 4:
                                    if (!(child.nodeType === Node.ELEMENT_NODE)) return [3 /*break*/, 12];
                                    childEl = child;
                                    if (!(childEl.tagName === 'IMG')) return [3 /*break*/, 11];
                                    _p.label = 5;
                                case 5:
                                    _p.trys.push([5, 9, , 10]);
                                    img = childEl;
                                    return [4 /*yield*/, fetch(img.src)];
                                case 6:
                                    response = _p.sent();
                                    return [4 /*yield*/, response.blob()];
                                case 7:
                                    blob = _p.sent();
                                    return [4 /*yield*/, blob.arrayBuffer()];
                                case 8:
                                    buffer = _p.sent();
                                    imgWidth = img.naturalWidth || 400;
                                    imgHeight = img.naturalHeight || 300;
                                    maxWidth = 400;
                                    maxHeight = 300;
                                    finalWidth = imgWidth;
                                    finalHeight = imgHeight;
                                    if (finalWidth > maxWidth) {
                                        finalHeight = (finalHeight * maxWidth) / finalWidth;
                                        finalWidth = maxWidth;
                                    }
                                    if (finalHeight > maxHeight) {
                                        finalWidth = (finalWidth * maxHeight) / finalHeight;
                                        finalHeight = maxHeight;
                                    }
                                    pRuns.push(new docx_1.TextRun({ text: "", break: 1 }), new docx_1.ImageRun({
                                        data: buffer,
                                        transformation: {
                                            width: Math.round(finalWidth),
                                            height: Math.round(finalHeight)
                                        },
                                        type: 'png'
                                    }), new docx_1.TextRun({ text: "", break: 1 }));
                                    return [3 /*break*/, 10];
                                case 9:
                                    _m = _p.sent();
                                    pRuns.push(new docx_1.TextRun({
                                        text: "[图片]",
                                        size: 24,
                                        font: "Arial"
                                    }));
                                    return [3 /*break*/, 10];
                                case 10: return [3 /*break*/, 12];
                                case 11:
                                    if (childEl.tagName === 'SUB') {
                                        text = (_c = childEl.textContent) === null || _c === void 0 ? void 0 : _c.trim();
                                        if (text) {
                                            pRuns.push(new docx_1.TextRun({
                                                text: text.replace(/\s+/g, ' '),
                                                size: 24,
                                                font: "Arial",
                                                subScript: true
                                            }));
                                        }
                                    }
                                    else if (childEl.tagName === 'SUP') {
                                        text = (_d = childEl.textContent) === null || _d === void 0 ? void 0 : _d.trim();
                                        if (text) {
                                            pRuns.push(new docx_1.TextRun({
                                                text: text.replace(/\s+/g, ' '),
                                                size: 24,
                                                font: "Arial",
                                                superScript: true
                                            }));
                                        }
                                    }
                                    else if (childEl.classList.contains('ql-formula')) {
                                        katexHtml = childEl.querySelector('.katex-html');
                                        formulaText = '';
                                        if (katexHtml) {
                                            formulaText = ((_e = katexHtml.textContent) === null || _e === void 0 ? void 0 : _e.trim()) || '';
                                        }
                                        else {
                                            formulaText = ((_f = childEl.textContent) === null || _f === void 0 ? void 0 : _f.trim()) || '';
                                        }
                                        if (formulaText) {
                                            pRuns.push(new docx_1.TextRun({
                                                text: formulaText,
                                                size: 24,
                                                font: "Arial"
                                            }));
                                        }
                                    }
                                    else {
                                        text = (_g = childEl.textContent) === null || _g === void 0 ? void 0 : _g.trim();
                                        if (text) {
                                            // Add the formatted text
                                            pRuns.push(new docx_1.TextRun({
                                                text: text.replace(/\s+/g, ' '), // Normalize whitespace
                                                size: 24,
                                                font: "Arial",
                                                bold: childEl.tagName === 'STRONG' || childEl.tagName === 'B',
                                                underline: childEl.tagName === 'U' ? { type: 'single' } : undefined,
                                                italics: childEl.tagName === 'I' || childEl.tagName === 'EM',
                                                break: text.endsWith('\n') ? 1 : undefined
                                            }));
                                            // Add space after bold text
                                            if (childEl.tagName === 'STRONG' || childEl.tagName === 'B') {
                                                pRuns.push(new docx_1.TextRun({
                                                    text: ' ',
                                                    size: 24,
                                                    font: "Arial"
                                                }));
                                            }
                                        }
                                    }
                                    _p.label = 12;
                                case 12:
                                    _k++;
                                    return [3 /*break*/, 3];
                                case 13:
                                    if (pRuns.length > 0) {
                                        // Add line break before paragraph if not the first run
                                        if (runs_1.length > 0) {
                                            runs_1.push(new docx_1.TextRun({ text: "", break: 1 }));
                                        }
                                        // Add paragraph content
                                        runs_1.push.apply(runs_1, pRuns);
                                        // Add line break after paragraph
                                        runs_1.push(new docx_1.TextRun({ text: "", break: 1 }));
                                    }
                                    return [3 /*break*/, 22];
                                case 14:
                                    if (!(el.tagName === 'IMG')) return [3 /*break*/, 21];
                                    _p.label = 15;
                                case 15:
                                    _p.trys.push([15, 19, , 20]);
                                    img = el;
                                    return [4 /*yield*/, fetch(img.src)];
                                case 16:
                                    response = _p.sent();
                                    return [4 /*yield*/, response.blob()];
                                case 17:
                                    blob = _p.sent();
                                    return [4 /*yield*/, blob.arrayBuffer()];
                                case 18:
                                    buffer = _p.sent();
                                    imgWidth = img.naturalWidth || 400;
                                    imgHeight = img.naturalHeight || 300;
                                    maxWidth = 400;
                                    maxHeight = 300;
                                    finalWidth = imgWidth;
                                    finalHeight = imgHeight;
                                    if (finalWidth > maxWidth) {
                                        finalHeight = (finalHeight * maxWidth) / finalWidth;
                                        finalWidth = maxWidth;
                                    }
                                    if (finalHeight > maxHeight) {
                                        finalWidth = (finalWidth * maxHeight) / finalHeight;
                                        finalHeight = maxHeight;
                                    }
                                    runs_1.push(new docx_1.TextRun({ text: "", break: 1 }), new docx_1.ImageRun({
                                        data: buffer,
                                        transformation: {
                                            width: Math.round(finalWidth),
                                            height: Math.round(finalHeight)
                                        },
                                        type: 'png'
                                    }), new docx_1.TextRun({ text: "", break: 1 }));
                                    return [3 /*break*/, 20];
                                case 19:
                                    _o = _p.sent();
                                    runs_1.push(new docx_1.TextRun({
                                        text: "[图片]",
                                        size: 24,
                                        font: "Arial"
                                    }));
                                    return [3 /*break*/, 20];
                                case 20: return [3 /*break*/, 22];
                                case 21:
                                    if (el.tagName === 'BR') {
                                        runs_1.push(new docx_1.TextRun({ text: "", break: 1 }));
                                    }
                                    else {
                                        // Process text content in other elements
                                        Array.from(el.childNodes).forEach(function (child) {
                                            var _a;
                                            if (child.nodeType === Node.TEXT_NODE) {
                                                var text = (_a = child.textContent) === null || _a === void 0 ? void 0 : _a.trim();
                                                if (text) {
                                                    runs_1.push(new docx_1.TextRun({
                                                        text: text,
                                                        size: 24,
                                                        font: "Arial"
                                                    }));
                                                }
                                            }
                                        });
                                    }
                                    _p.label = 22;
                                case 22:
                                    _i++;
                                    return [3 /*break*/, 1];
                                case 23:
                                    if (runs_1.length > 0) {
                                        paragraphs.push(new docx_1.Paragraph({
                                            children: runs_1,
                                            spacing: { after: 200 },
                                            alignment: "left"
                                        }));
                                    }
                                    _p.label = 24;
                                case 24: return [2 /*return*/];
                            }
                        });
                    };
                    i = 0;
                    _h.label = 1;
                case 1:
                    if (!(i < questions.length)) return [3 /*break*/, 4];
                    return [5 /*yield**/, _loop_1(i)];
                case 2:
                    _h.sent();
                    _h.label = 3;
                case 3:
                    i++;
                    return [3 /*break*/, 1];
                case 4:
                    doc = new docx_1.Document({
                        title: title,
                        description: "Exam paper export generated on ".concat(new Date().toLocaleString()),
                        styles: {
                            default: {
                                document: {
                                    run: {
                                        font: "Arial",
                                        size: 24
                                    }
                                }
                            }
                        },
                        sections: [{
                                properties: {
                                    page: {
                                        margin: {
                                            top: 1440,
                                            right: 1440,
                                            bottom: 1440,
                                            left: 1440
                                        }
                                    }
                                },
                                children: paragraphs
                            }]
                    });
                    _h.label = 5;
                case 5:
                    _h.trys.push([5, 7, , 8]);
                    return [4 /*yield*/, docx_1.Packer.toBlob(doc)];
                case 6:
                    blob = _h.sent();
                    url = window.URL.createObjectURL(blob);
                    a = document.createElement('a');
                    a.href = url;
                    a.download = "".concat(title, "_").concat(timestamp, ".docx");
                    a.click();
                    window.URL.revokeObjectURL(url);
                    return [3 /*break*/, 8];
                case 7:
                    error_1 = _h.sent();
                    console.error('Failed to generate document:', error_1);
                    throw new Error('Failed to create Word document');
                case 8: return [2 /*return*/];
            }
        });
    });
}
