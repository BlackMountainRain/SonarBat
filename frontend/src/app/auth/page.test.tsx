import { render, screen } from '@testing-library/react';
import React from 'react';
import { useOAuth2 } from '@tasoskakour/react-use-oauth2';
import LoginPage from '@/app/auth/page';

jest.mock('@tasoskakour/react-use-oauth2');

const mockUseOAuth = useOAuth2 as jest.Mock;
mockUseOAuth.mockReturnValue({ getAuth: jest.fn() });

jest.mock('next/navigation', () => ({
  useRouter() {
    return {
      push: jest.fn(),
      prefetch: () => null,
    };
  },
}));

describe('renders auth page', () => {
  // Renders without crashing
  it('should render without crashing', () => {
    // Arrange
    render(<LoginPage />);

    // Act
    const logoImage = screen.getByAltText('logo');

    // Assert
    expect(logoImage).toBeInTheDocument();
  });

  // Renders correctly when screen width is exactly 'sm'
  it('should render correctly when screen width is exactly "sm"', () => {
    // Arrange
    Object.defineProperty(window, 'innerWidth', {
      writable: true,
      configurable: true,
      value: 640,
    });
    window.dispatchEvent(new Event('resize')); // Trigger resize event

    // Act
    render(<LoginPage />);

    // Assert
    expect(screen.getByText('Sonar Bat')).toBeInTheDocument();
  });

  // Renders correctly when screen width is exactly 'md'
  it('should render correctly when screen width is exactly "md"', () => {
    // Arrange
    Object.defineProperty(window, 'innerWidth', {
      writable: true,
      configurable: true,
      value: 768,
    });
    window.dispatchEvent(new Event('resize')); // Trigger resize event

    // Act
    render(<LoginPage />);

    // Assert
    expect(screen.getByText('Sonar Bat')).toBeInTheDocument();
  });
});
