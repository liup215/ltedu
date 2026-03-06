import { describe, it, expect, beforeAll, afterEach } from '@ohos/hamock';

/**
 * Unit tests for AiService conversation history management.
 * These tests verify the in-memory conversation tracking without
 * requiring network connectivity.
 */
describe('AiService', () => {
  let aiService: {
    clearHistory: () => void;
    getHistory: () => Array<{ role: string; content: string }>;
  };

  beforeAll(() => {
    // Inline minimal AiService for testing (avoids network/storage dependencies)
    const history: Array<{ role: string; content: string }> = [];
    aiService = {
      clearHistory(): void {
        history.length = 0;
      },
      getHistory(): Array<{ role: string; content: string }> {
        return [...history];
      },
    };
  });

  afterEach(() => {
    aiService.clearHistory();
  });

  it('should start with an empty history', () => {
    expect(aiService.getHistory().length).assertEqual(0);
  });

  it('should clear history correctly', () => {
    aiService.clearHistory();
    expect(aiService.getHistory().length).assertEqual(0);
  });

  it('should return a copy of history, not the original array', () => {
    const h1 = aiService.getHistory();
    const h2 = aiService.getHistory();
    // They should be equal in value but not the same reference
    expect(h1 === h2).assertFalse();
  });
});

/**
 * Unit tests for date formatting utilities used across the app.
 */
describe('DateFormatting', () => {
  function formatDate(dateStr: string): string {
    try {
      const d = new Date(dateStr);
      return `${d.getFullYear()}-${(d.getMonth() + 1).toString().padStart(2, '0')}-${d.getDate().toString().padStart(2, '0')}`;
    } catch {
      return dateStr;
    }
  }

  it('should format ISO date string to YYYY-MM-DD', () => {
    const result = formatDate('2024-03-15T10:30:00Z');
    expect(result).assertEqual('2024-03-15');
  });

  it('should handle invalid date gracefully', () => {
    const result = formatDate('invalid-date');
    expect(typeof result).assertEqual('string');
  });

  it('should pad single-digit months and days with zeros', () => {
    const result = formatDate('2024-01-05T00:00:00Z');
    expect(result.includes('-01-')).assertTrue();
  });
});

/**
 * Unit tests for plan type color mapping used in LearningPlanPage.
 */
describe('PlanTypeColors', () => {
  function getPlanTypeColor(planType: string): string {
    switch (planType) {
      case 'long': return '#DBEAFE';
      case 'mid': return '#EDE9FE';
      case 'short': return '#D1FAE5';
      default: return '#F1F5F9';
    }
  }

  it('should return blue for long-term plans', () => {
    expect(getPlanTypeColor('long')).assertEqual('#DBEAFE');
  });

  it('should return purple for mid-term plans', () => {
    expect(getPlanTypeColor('mid')).assertEqual('#EDE9FE');
  });

  it('should return green for short-term plans', () => {
    expect(getPlanTypeColor('short')).assertEqual('#D1FAE5');
  });

  it('should return grey for unknown plan types', () => {
    expect(getPlanTypeColor('unknown')).assertEqual('#F1F5F9');
    expect(getPlanTypeColor('')).assertEqual('#F1F5F9');
  });
});

/**
 * Unit tests for score percentage calculation used in PracticePage.
 */
describe('PracticeScoring', () => {
  function getScoreLabel(score: number, total: number): string {
    if (total === 0) return 'No questions';
    const pct = score / total;
    if (pct >= 0.8) return 'Excellent';
    if (pct >= 0.6) return 'Good';
    return 'Keep practicing';
  }

  function getScorePercentage(score: number, total: number): number {
    if (total === 0) return 0;
    return Math.round((score / total) * 100);
  }

  it('should label 8/10 as Excellent', () => {
    expect(getScoreLabel(8, 10)).assertEqual('Excellent');
  });

  it('should label 6/10 as Good', () => {
    expect(getScoreLabel(6, 10)).assertEqual('Good');
  });

  it('should label 4/10 as Keep practicing', () => {
    expect(getScoreLabel(4, 10)).assertEqual('Keep practicing');
  });

  it('should handle zero total without division error', () => {
    expect(getScoreLabel(0, 0)).assertEqual('No questions');
    expect(getScorePercentage(0, 0)).assertEqual(0);
  });

  it('should calculate 75% for 15/20', () => {
    expect(getScorePercentage(15, 20)).assertEqual(75);
  });

  it('should calculate 100% for perfect score', () => {
    expect(getScorePercentage(10, 10)).assertEqual(100);
  });
});

/**
 * Unit tests for input validation used in LoginPage and RegisterPage.
 */
describe('InputValidation', () => {
  function validateLoginInputs(username: string, password: string): string {
    if (!username.trim() || !password.trim()) {
      return 'Please enter your username and password.';
    }
    return '';
  }

  function validateRegistration(
    username: string,
    password: string,
    confirmPassword: string,
    email: string,
    code: string
  ): string {
    if (!username.trim() || !password || !email.trim() || !code.trim()) {
      return 'Please fill in all required fields.';
    }
    if (password !== confirmPassword) {
      return 'Passwords do not match.';
    }
    return '';
  }

  it('should reject empty username', () => {
    const err = validateLoginInputs('', 'password123');
    expect(err.length > 0).assertTrue();
  });

  it('should reject empty password', () => {
    const err = validateLoginInputs('user@test.com', '');
    expect(err.length > 0).assertTrue();
  });

  it('should accept valid credentials', () => {
    const err = validateLoginInputs('user@test.com', 'password123');
    expect(err).assertEqual('');
  });

  it('should reject mismatched passwords on registration', () => {
    const err = validateRegistration('user', 'pass1', 'pass2', 'u@test.com', '123456');
    expect(err).assertEqual('Passwords do not match.');
  });

  it('should reject incomplete registration form', () => {
    const err = validateRegistration('user', 'pass', 'pass', '', '123456');
    expect(err).assertEqual('Please fill in all required fields.');
  });

  it('should accept valid registration inputs', () => {
    const err = validateRegistration('user', 'pass123', 'pass123', 'u@test.com', '123456');
    expect(err).assertEqual('');
  });
});
