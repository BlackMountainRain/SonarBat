import { render, screen } from '@testing-library/react';
import React from 'react';
import CallbackPage from '@/app/auth/callback/page';

window.opener = {
  postMessage: jest.fn(),
};

describe('Resources Page', () => {
  it('renders resource page', () => {
    render(<CallbackPage />);

    const ele = screen.getByText('Loading...');

    expect(ele).toBeInTheDocument();
  });
});
