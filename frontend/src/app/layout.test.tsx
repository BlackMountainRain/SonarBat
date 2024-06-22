import React from 'react';
import { render, screen } from '@testing-library/react';
import { Providers } from '@/app/layout';

describe('Providers', () => {
  it('renders children correctly', () => {
    render(
      <Providers>
        <div data-testid="test-child">Test Child</div>
      </Providers>,
    );

    const childElement = screen.getByTestId('test-child');
    expect(childElement).toBeInTheDocument();
    expect(childElement).toHaveTextContent('Test Child');
  });
});
