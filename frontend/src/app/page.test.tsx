import { render, screen } from '@testing-library/react';
import Home from './page';

describe('Page', () => {
  it('renders home page', () => {
    // eslint-disable-next-line react/react-in-jsx-scope
    render(<Home />);

    const heading = screen.getByText('Main Page');

    expect(heading).toBeInTheDocument();
  });
});
