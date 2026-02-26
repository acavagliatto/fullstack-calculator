import { useState } from 'react';
import type { FormEvent, ChangeEvent } from 'react';
import { OPERATIONS } from './types';
import { calculate } from './api';
import './Calculator.css';

function Calculator() {
  const [operation, setOperation] = useState<string>('add');
  const [operands, setOperands] = useState<string[]>(['', '']);
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);

  const config = OPERATIONS[operation];

  // Update operands array when operation changes
  const handleOperationChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const newOperation = e.target.value;
    setOperation(newOperation);
    const newConfig = OPERATIONS[newOperation];
    setOperands(new Array(newConfig.operandCount).fill(''));
    setResult(null);
    setError('');
  };

  const handleOperandChange = (index: number, value: string) => {
    const newOperands = [...operands];
    newOperands[index] = value;
    setOperands(newOperands);
  };

  const validateInputs = (): boolean => {
    // Check if all operands are filled
    for (let i = 0; i < operands.length; i++) {
      if (operands[i].trim() === '') {
        setError(`Please enter ${config.inputLabels[i]}`);
        return false;
      }
    }

    // Check if all operands are valid numbers
    for (let i = 0; i < operands.length; i++) {
      const num = parseFloat(operands[i]);
      if (isNaN(num)) {
        setError(`${config.inputLabels[i]} must be a valid number`);
        return false;
      }
    }

    return true;
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError('');
    setResult(null);

    if (!validateInputs()) {
      return;
    }

    const numericOperands = operands.map((op) => parseFloat(op));

    setLoading(true);
    try {
      const response = await calculate(operation, numericOperands);
      setResult(response.result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred');
    } finally {
      setLoading(false);
    }
  };

  const handleReset = () => {
    setOperands(new Array(config.operandCount).fill(''));
    setResult(null);
    setError('');
  };

  return (
    <div className="calculator">
      <h1>Calculator</h1>
      
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="operation">Operation:</label>
          <select
            id="operation"
            value={operation}
            onChange={handleOperationChange}
            disabled={loading}
          >
            {Object.entries(OPERATIONS).map(([key, op]) => (
              <option key={key} value={key}>
                {op.symbol} {op.label}
              </option>
            ))}
          </select>
        </div>

        {operands.map((operand, index) => (
          <div key={index} className="form-group">
            <label htmlFor={`operand-${index}`}>
              {config.inputLabels[index]}:
            </label>
            <input
              id={`operand-${index}`}
              type="text"
              inputMode="decimal"
              value={operand}
              onChange={(e) => handleOperandChange(index, e.target.value)}
              placeholder={`Enter ${config.inputLabels[index].toLowerCase()}`}
              disabled={loading}
            />
          </div>
        ))}

        <div className="button-group">
          <button type="submit" disabled={loading}>
            {loading ? 'Calculating...' : 'Calculate'}
          </button>
          <button type="button" onClick={handleReset} disabled={loading}>
            Reset
          </button>
        </div>
      </form>

      {result !== null && (
        <div className="result success" data-testid="result">
          <strong>Result:</strong> {result}
        </div>
      )}

      {error && (
        <div className="result error" data-testid="error">
          <strong>Error:</strong> {error}
        </div>
      )}
    </div>
  );
}

export default Calculator;
