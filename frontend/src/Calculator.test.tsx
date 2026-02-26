import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import Calculator from './Calculator';
import * as api from './api';

// Mock the API module
vi.mock('./api');

describe('Calculator Component', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('renders calculator with default operation', () => {
    render(<Calculator />);
    expect(screen.getByText('Calculator')).toBeInTheDocument();
    expect(screen.getByLabelText('Operation:')).toBeInTheDocument();
    expect(screen.getByDisplayValue('+ Add')).toBeInTheDocument();
  });

  it('displays correct number of inputs for add operation', () => {
    render(<Calculator />);
    expect(screen.getByLabelText('First Number:')).toBeInTheDocument();
    expect(screen.getByLabelText('Second Number:')).toBeInTheDocument();
  });

  it('changes inputs when operation changes', async () => {
    const user = userEvent.setup();
    render(<Calculator />);
    
    const select = screen.getByLabelText('Operation:');
    await user.selectOptions(select, 'sqrt');
    
    expect(screen.getByLabelText('Number:')).toBeInTheDocument();
    expect(screen.queryByLabelText('Second Number:')).not.toBeInTheDocument();
  });

  it('performs successful addition', async () => {
    const user = userEvent.setup();
    const mockCalculate = vi.mocked(api.calculate);
    mockCalculate.mockResolvedValue({ result: 8 });

    render(<Calculator />);
    
    await user.type(screen.getByLabelText('First Number:'), '5');
    await user.type(screen.getByLabelText('Second Number:'), '3');
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('result')).toHaveTextContent('Result: 8');
    });

    expect(mockCalculate).toHaveBeenCalledWith('add', [5, 3]);
  });

  it('performs successful division', async () => {
    const user = userEvent.setup();
    const mockCalculate = vi.mocked(api.calculate);
    mockCalculate.mockResolvedValue({ result: 5 });

    render(<Calculator />);
    
    const select = screen.getByLabelText('Operation:');
    await user.selectOptions(select, 'divide');
    
    await user.type(screen.getByLabelText('Dividend:'), '10');
    await user.type(screen.getByLabelText('Divisor:'), '2');
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('result')).toHaveTextContent('Result: 5');
    });

    expect(mockCalculate).toHaveBeenCalledWith('divide', [10, 2]);
  });

  it('displays error for division by zero', async () => {
    const user = userEvent.setup();
    const mockCalculate = vi.mocked(api.calculate);
    mockCalculate.mockRejectedValue(new Error('division by zero'));

    render(<Calculator />);
    
    const select = screen.getByLabelText('Operation:');
    await user.selectOptions(select, 'divide');
    
    await user.type(screen.getByLabelText('Dividend:'), '10');
    await user.type(screen.getByLabelText('Divisor:'), '0');
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('error')).toHaveTextContent('division by zero');
    });
  });

  it('validates empty inputs', async () => {
    const user = userEvent.setup();
    render(<Calculator />);
    
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('error')).toHaveTextContent('Please enter First Number');
    });
  });

  it('validates invalid number inputs', async () => {
    const user = userEvent.setup();
    render(<Calculator />);
    
    await user.type(screen.getByLabelText('First Number:'), 'abc');
    await user.type(screen.getByLabelText('Second Number:'), '5');
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('error')).toHaveTextContent('must be a valid number');
    });
  });

  it('handles square root operation', async () => {
    const user = userEvent.setup();
    const mockCalculate = vi.mocked(api.calculate);
    mockCalculate.mockResolvedValue({ result: 3 });

    render(<Calculator />);
    
    const select = screen.getByLabelText('Operation:');
    await user.selectOptions(select, 'sqrt');
    
    await user.type(screen.getByLabelText('Number:'), '9');
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('result')).toHaveTextContent('Result: 3');
    });

    expect(mockCalculate).toHaveBeenCalledWith('sqrt', [9]);
  });

  it('displays error for negative square root', async () => {
    const user = userEvent.setup();
    const mockCalculate = vi.mocked(api.calculate);
    mockCalculate.mockRejectedValue(new Error('cannot compute square root of negative number'));

    render(<Calculator />);
    
    const select = screen.getByLabelText('Operation:');
    await user.selectOptions(select, 'sqrt');
    
    await user.type(screen.getByLabelText('Number:'), '-4');
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('error')).toHaveTextContent('cannot compute square root of negative number');
    });
  });

  it('resets form when reset button is clicked', async () => {
    const user = userEvent.setup();
    const mockCalculate = vi.mocked(api.calculate);
    mockCalculate.mockResolvedValue({ result: 8 });

    render(<Calculator />);
    
    await user.type(screen.getByLabelText('First Number:'), '5');
    await user.type(screen.getByLabelText('Second Number:'), '3');
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('result')).toBeInTheDocument();
    });

    await user.click(screen.getByText('Reset'));

    expect(screen.getByLabelText('First Number:')).toHaveValue('');
    expect(screen.getByLabelText('Second Number:')).toHaveValue('');
    expect(screen.queryByTestId('result')).not.toBeInTheDocument();
  });

  it('disables inputs while calculating', async () => {
    const user = userEvent.setup();
    const mockCalculate = vi.mocked(api.calculate);
    
    // Simulate a delayed response
    mockCalculate.mockImplementation(() => 
      new Promise(resolve => setTimeout(() => resolve({ result: 8 }), 100))
    );

    render(<Calculator />);
    
    await user.type(screen.getByLabelText('First Number:'), '5');
    await user.type(screen.getByLabelText('Second Number:'), '3');
    
    const calculateButton = screen.getByText('Calculate');
    await user.click(calculateButton);

    // Check that button shows "Calculating..." and is disabled
    expect(screen.getByText('Calculating...')).toBeInTheDocument();
    expect(calculateButton).toBeDisabled();

    await waitFor(() => {
      expect(screen.getByTestId('result')).toBeInTheDocument();
    });
  });

  it('handles percentage operation', async () => {
    const user = userEvent.setup();
    const mockCalculate = vi.mocked(api.calculate);
    mockCalculate.mockResolvedValue({ result: 10 });

    render(<Calculator />);
    
    const select = screen.getByLabelText('Operation:');
    await user.selectOptions(select, 'percentage');
    
    await user.type(screen.getByLabelText('Value:'), '100');
    await user.type(screen.getByLabelText('Percentage:'), '10');
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('result')).toHaveTextContent('Result: 10');
    });

    expect(mockCalculate).toHaveBeenCalledWith('percentage', [100, 10]);
  });

  it('clears error when switching operations', async () => {
    const user = userEvent.setup();
    render(<Calculator />);
    
    await user.click(screen.getByText('Calculate'));

    await waitFor(() => {
      expect(screen.getByTestId('error')).toBeInTheDocument();
    });

    const select = screen.getByLabelText('Operation:');
    await user.selectOptions(select, 'subtract');

    expect(screen.queryByTestId('error')).not.toBeInTheDocument();
  });
});
