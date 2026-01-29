
// Import ApiResponse from api model
import type { ApiResponse } from './api.model';
import type { Chapter } from './chapter.model';
import type { PastPaper } from './pastPaper.model';
import type { Syllabus } from './syllabus.model';


export interface Question {
  id: number;
  syllabusId: number;
  stem: string;
  totalScore: number;
  difficult: number;
  status: number;
  syllabus?: Syllabus;
  pastPaperId?: number;
  pastPaper?: PastPaper;
  chapters?: Chapter[];
  questionContents?: QuestionContent[];
  createdAt?: string;
  updatedAt?: string;
  indexInPastPaper: number;
}

export interface QuestionContent {
  partLabel?: string;
  subpartLabel?: string;
  score?: number;
  questionTypeId?: number;
  analyze?: string;
  singleChoice?: SingleChoiceContent;
  multipleChoice?: MultipleChoiceContent;
  trueOrFalse?: TrueOrFalseContent;
  gapFilling?: GapFillingContent;
  shortAnswer?: ShortAnswerContent;
}

export interface SingleChoiceContent {
  options: ChoiceOption[];
  answer: string;
}

export interface MultipleChoiceContent {
  options: ChoiceOption[];
  answer: string[];
}

export interface TrueOrFalseContent {
  answer: number; // 1 true, 2 false
}

export interface GapFillingContent {
  answer: string[];
}

export interface ShortAnswerContent {
  answer: string;
}

export interface ChoiceOption {
  prefix: string;
  content: string;
}

export interface QuestionQuery {
  pageIndex?: number;
  pageSize?: number;
  syllabusId?: number;
  chapters?: number[]; // Changed to 'chapters' and type number[] based on feedback
  difficult?: number;
  status?: number;
  stem?: string;
  paperName?: string;
}

export interface QuestionUpdateRequest {
  id: number;
  stem: string;
  totalScore: number;
  difficult: number;
  status: number;
  questionContents: QuestionContent[];
}

export interface QuestionChapterRequest {
  questionId: number;
  chapters: number[];
}

export interface PaginatedQuestions {
  list: Question[];
  total: number;
}

export const DIFFICULTY_NAMES = {
  1: 'Easy',
  2: 'Medium',
  3: 'Hard',
  4: 'Very Hard',
  5: 'Extremely Hard'
} as const;

export const QUESTION_STATE_NORMAL = 1;
export const QUESTION_STATE_FORBIDDEN = 2;
export const QUESTION_STATE_DELETE = 3;

// Question Chapter Update Request
export interface QuestionChapterUpdateRequest {
  questionId: number;
  chapters: number[];
}

// API Response Types
export type QuestionListResponse = ApiResponse<PaginatedQuestions>;
export type QuestionResponse = ApiResponse<Question>;

export const QUESTION_TYPE_SINGLE_CHOICE = 1;
export const QUESTION_TYPE_MULTIPLE_CHOICE = 2;
export const QUESTION_TYPE_TRUE_FALSE = 3;
export const QUESTION_TYPE_GAP_FILLING = 4;
export const QUESTION_TYPE_SHORT_ANSWER = 5;

// Question Type Names Mapping
export const QUESTION_TYPE_NAMES = {
  [QUESTION_TYPE_SINGLE_CHOICE]: 'single choice question',
  [QUESTION_TYPE_MULTIPLE_CHOICE]: 'multiple choice question',
  [QUESTION_TYPE_TRUE_FALSE]: 'true/false question',
  [QUESTION_TYPE_GAP_FILLING]: 'gap filling question',
  [QUESTION_TYPE_SHORT_ANSWER]: 'short answer question'
};

// Question Status Names Mapping
export const QUESTION_STATUS_NAMES = {
  [QUESTION_STATE_NORMAL]: 'normal',
  [QUESTION_STATE_FORBIDDEN]: 'forbidden',
  [QUESTION_STATE_DELETE]: 'deleted'
};
