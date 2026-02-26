import type { CalculateRequest, CalculateResponse, ErrorResponse } from './types';

const API_BASE_URL = '/api';

export async function calculate(
  operation: string,
  operands: number[]
): Promise<CalculateResponse> {
  const response = await fetch(`${API_BASE_URL}/calculate`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ operation, operands } as CalculateRequest),
  });

  if (!response.ok) {
    const errorData: ErrorResponse = await response.json();
    throw new Error(errorData.error || 'Calculation failed');
  }

  return response.json();
}

export async function healthCheck(): Promise<{ status: string }> {
  const response = await fetch(`${API_BASE_URL}/health`);
  
  if (!response.ok) {
    throw new Error('Health check failed');
  }

  return response.json();
}
