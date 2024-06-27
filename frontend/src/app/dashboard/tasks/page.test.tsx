import { render, screen } from '@testing-library/react';
import React from 'react';
import TaskPage from '@/app/dashboard/tasks/page';

describe('Tasks Page', () => {
  it('renders task page', () => {
    render(<TaskPage />);

    const heading = screen.getByText('All Tasks');

    expect(heading).toBeInTheDocument();
  });
});
