export interface CalculateRequest {
  operation: string;
  operands: number[];
}

export interface CalculateResponse {
  result: number;
}

export interface ErrorResponse {
  error: string;
}

export interface OperationConfig {
  label: string;
  operandCount: number;
  inputLabels: string[];
  symbol: string;
}

export const OPERATIONS: Record<string, OperationConfig> = {
  add: {
    label: 'Add',
    operandCount: 2,
    inputLabels: ['First Number', 'Second Number'],
    symbol: '+',
  },
  subtract: {
    label: 'Subtract',
    operandCount: 2,
    inputLabels: ['First Number', 'Second Number'],
    symbol: '-',
  },
  multiply: {
    label: 'Multiply',
    operandCount: 2,
    inputLabels: ['First Number', 'Second Number'],
    symbol: '×',
  },
  divide: {
    label: 'Divide',
    operandCount: 2,
    inputLabels: ['Dividend', 'Divisor'],
    symbol: '÷',
  },
  exponentiation: {
    label: 'Power',
    operandCount: 2,
    inputLabels: ['Base', 'Exponent'],
    symbol: '^',
  },
  sqrt: {
    label: 'Square Root',
    operandCount: 1,
    inputLabels: ['Number'],
    symbol: '√',
  },
  percentage: {
    label: 'Percentage',
    operandCount: 2,
    inputLabels: ['Value', 'Percentage'],
    symbol: '%',
  },
};
