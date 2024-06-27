import { render, screen } from '@testing-library/react';
import React from 'react';
import ResourcePage from '@/app/dashboard/resources/page';

describe('Resources Page', () => {
  it('renders resource page', () => {
    render(<ResourcePage />);

    const heading = screen.getByText('All Resources');

    expect(heading).toBeInTheDocument();
  });
});
