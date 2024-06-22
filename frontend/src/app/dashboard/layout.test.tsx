import { render, screen } from '@testing-library/react';
import React from 'react';
import DashboardLayout from '@/app/dashboard/layout';

const pushMock = jest.fn();

jest.mock('next/navigation', () => ({
  useRouter() {
    return {
      push: pushMock,
      prefetch: () => null,
    };
  },
  usePathname() {
    return '/dashboard';
  },
}));

describe('DashboardLayout', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  // Renders children correctly
  it('should render children correctly', () => {
    // Act
    render(
      <DashboardLayout>
        <div>Test Child</div>
      </DashboardLayout>,
    );
    const childElement = screen.getByText('Test Child');

    // Assert
    expect(childElement).toBeInTheDocument();
  });

  // No token in localStorage
  it('should redirect to /auth when no token in localStorage', () => {
    // Arrange
    jest.spyOn(Storage.prototype, 'getItem').mockReturnValue(null);

    // Act
    render(
      <DashboardLayout>
        <div>Test Child</div>
      </DashboardLayout>,
    );

    expect(pushMock).toHaveBeenCalledTimes(1);

    // Assert
    expect(pushMock).toHaveBeenCalledWith('/auth');
  });

  // Token in localStorage is valid
  it('should not redirect when token in localStorage is valid', () => {
    // Arrange
    jest
      .spyOn(Storage.prototype, 'getItem')
      .mockReturnValue(
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc2MDAwMDAwMDB9.123',
      );

    // Act
    render(
      <DashboardLayout>
        <div>Test Child</div>
      </DashboardLayout>,
    );

    // Assert
    expect(pushMock).not.toHaveBeenCalled();
  });

  // Token in localStorage is expired
  it('should redirect to /auth when token in localStorage is expired', async () => {
    // Arrange
    jest
      .spyOn(Storage.prototype, 'getItem')
      .mockReturnValue(
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDAwMDAwMDB9.123',
      );

    // Act
    render(
      <DashboardLayout>
        <div>Test Child</div>
      </DashboardLayout>,
    );

    // Assert
    expect(pushMock).toHaveBeenCalledTimes(1);
    expect(pushMock).toHaveBeenCalledWith('/auth');
  });
});
