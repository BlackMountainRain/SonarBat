import { fireEvent, render, screen } from '@testing-library/react';
import React from 'react';
import { useOAuth2 } from '@tasoskakour/react-use-oauth2';
import OAuthForm from '@/app/ui/auth/oauth-form';

jest.mock('@tasoskakour/react-use-oauth2');

const pushMock = jest.fn();
jest.mock('next/navigation', () => ({
  useRouter() {
    return {
      push: pushMock,
      prefetch: () => null,
    };
  },
}));

const mockUseOAuth = useOAuth2 as jest.Mock;

describe('OAuthForm', () => {
  beforeAll(() => {
    global.open = jest.fn();
  });

  beforeEach(() => {
    jest.clearAllMocks();
    mockUseOAuth.mockReturnValue({
      getAuth: jest.fn(),
    });
  });

  it('renders GitHub and Google sign-in buttons', () => {
    render(<OAuthForm />);
    expect(screen.getByTestId('github-button')).toBeInTheDocument();
    expect(screen.getByTestId('google-button')).toBeInTheDocument();
  });

  it('calls getGitHubAuth when GitHub button is clicked', () => {
    const getGitHubAuth = jest.fn();
    mockUseOAuth.mockReturnValue({ getAuth: getGitHubAuth });

    render(<OAuthForm />);
    const githubButton = screen.getByTestId('github-button');
    fireEvent.click(githubButton);

    expect(getGitHubAuth).toHaveBeenCalled();
  });

  it('calls getGoogleAuth when Google button is clicked', () => {
    const getGoogleAuth = jest.fn();
    mockUseOAuth.mockReturnValue({ getAuth: getGoogleAuth });

    render(<OAuthForm />);
    const googleButton = screen.getByTestId('google-button');
    fireEvent.click(googleButton);

    expect(getGoogleAuth).toHaveBeenCalled();
  });
});
