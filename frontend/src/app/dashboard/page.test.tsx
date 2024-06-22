import { render, screen } from '@testing-library/react';
import React from 'react';
import DashboardPage from '@/app/dashboard/page';

describe('Dashboard Page', () => {
  it('should render dashboard page', () => {
    render(<DashboardPage />);

    const heading = screen.getByText('Dashboard');

    expect(heading).toBeInTheDocument();
  });
});
