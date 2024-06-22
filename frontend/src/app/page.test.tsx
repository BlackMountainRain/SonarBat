import { render, screen } from '@testing-library/react';
import Home from '@/app/page';

describe('Home Page', () => {
  it('renders home page', () => {
    // eslint-disable-next-line react/react-in-jsx-scope
    render(<Home />);

    const heading = screen.getByText('Main Page');

    expect(heading).toBeInTheDocument();
  });
});
