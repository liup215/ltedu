"use strict";
var _a, _b;
Object.defineProperty(exports, "__esModule", { value: true });
exports.QUESTION_STATUS_NAMES = exports.QUESTION_TYPE_NAMES = exports.QUESTION_TYPE_SHORT_ANSWER = exports.QUESTION_TYPE_GAP_FILLING = exports.QUESTION_TYPE_TRUE_FALSE = exports.QUESTION_TYPE_MULTIPLE_CHOICE = exports.QUESTION_TYPE_SINGLE_CHOICE = exports.QUESTION_STATE_DELETE = exports.QUESTION_STATE_FORBIDDEN = exports.QUESTION_STATE_NORMAL = exports.DIFFICULTY_NAMES = void 0;
exports.DIFFICULTY_NAMES = {
    1: 'Easy',
    2: 'Medium',
    3: 'Hard',
    4: 'Very Hard',
    5: 'Extremely Hard'
};
exports.QUESTION_STATE_NORMAL = 1;
exports.QUESTION_STATE_FORBIDDEN = 2;
exports.QUESTION_STATE_DELETE = 3;
exports.QUESTION_TYPE_SINGLE_CHOICE = 1;
exports.QUESTION_TYPE_MULTIPLE_CHOICE = 2;
exports.QUESTION_TYPE_TRUE_FALSE = 3;
exports.QUESTION_TYPE_GAP_FILLING = 4;
exports.QUESTION_TYPE_SHORT_ANSWER = 5;
// Question Type Names Mapping
exports.QUESTION_TYPE_NAMES = (_a = {},
    _a[exports.QUESTION_TYPE_SINGLE_CHOICE] = 'single choice question',
    _a[exports.QUESTION_TYPE_MULTIPLE_CHOICE] = 'multiple choice question',
    _a[exports.QUESTION_TYPE_TRUE_FALSE] = 'true/false question',
    _a[exports.QUESTION_TYPE_GAP_FILLING] = 'gap filling question',
    _a[exports.QUESTION_TYPE_SHORT_ANSWER] = 'short answer question',
    _a);
// Question Status Names Mapping
exports.QUESTION_STATUS_NAMES = (_b = {},
    _b[exports.QUESTION_STATE_NORMAL] = 'normal',
    _b[exports.QUESTION_STATE_FORBIDDEN] = 'forbidden',
    _b[exports.QUESTION_STATE_DELETE] = 'deleted',
    _b);
