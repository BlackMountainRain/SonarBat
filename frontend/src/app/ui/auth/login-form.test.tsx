import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import React from 'react';
import { useOAuth2 } from '@tasoskakour/react-use-oauth2';
import {
  fetchToken as originalFetchToken,
  signUp as originalSignUp,
} from '@/app/lib/data';
import LoginForm from '@/app/ui/auth/login-form';
import toast from '@/app/helpers/toast';

jest.mock('@tasoskakour/react-use-oauth2');

const mockUseOAuth = useOAuth2 as jest.Mock;
mockUseOAuth.mockReturnValue({ getAuth: jest.fn() });

const pushMock = jest.fn();

jest.mock('next/navigation', () => ({
  useRouter() {
    return {
      push: pushMock,
      prefetch: () => null,
    };
  },
}));

jest.mock('@/app/lib/data', () => ({
  fetchToken: jest.fn(),
  signUp: jest.fn(),
}));

jest.mock('@/app/helpers/toast', () => ({
  success: jest.fn(),
  error: jest.fn(),
}));

const fetchToken = originalFetchToken as jest.MockedFunction<
  typeof originalFetchToken
>;

const signUp = originalSignUp as jest.MockedFunction<typeof originalSignUp>;

const correctEmail = 'test@example.com';
const correctPassword = 'password123';

describe('LoginForm', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  // User can successfully sign in with valid email and password
  it('should sign in successfully with valid email and password', async () => {
    // Arrange
    fetchToken.mockResolvedValueOnce('mock-token');
    jest.spyOn(Storage.prototype, 'setItem');
    render(<LoginForm />);

    // Act
    fireEvent.change(screen.getByLabelText(/email/i), {
      target: { value: correctEmail },
    });
    fireEvent.change(screen.getByLabelText(/password/i), {
      target: { value: correctPassword },
    });
    fireEvent.click(screen.getByTestId('submit-button'));

    // Assert
    await waitFor(() => {
      expect(fetchToken).toHaveBeenCalledWith(
        'SELF',
        '',
        correctEmail,
        correctPassword,
      );
      expect(localStorage.setItem).toHaveBeenCalledWith('token', 'mock-token');
      expect(pushMock).toHaveBeenCalledWith('/dashboard');
      expect(toast.success).toHaveBeenCalledWith('Sign in successfully');
    });
  });

  // Sign in with invalid email or password shows error toast
  it('should show error toast when signing in with invalid email or password', async () => {
    // Arrange
    fetchToken.mockRejectedValueOnce(new Error('Invalid email or password'));
    render(<LoginForm />);

    // Act
    fireEvent.change(screen.getByLabelText(/email/i), {
      target: { value: correctEmail },
    });
    fireEvent.change(screen.getByLabelText(/password/i), {
      target: { value: correctPassword },
    });
    fireEvent.click(screen.getByTestId('submit-button'));

    // Assert
    await waitFor(() => {
      expect(fetchToken).toHaveBeenCalledWith(
        'SELF',
        '',
        correctEmail,
        correctPassword,
      );
      expect(toast.error).toHaveBeenCalledWith(
        'Error: Invalid email or password',
      );
    });
  });

  // User can successfully sign up with valid email and password
  it('should sign up successfully with valid email and password', async () => {
    // Arrange
    fetchToken.mockResolvedValueOnce('mock-token');
    signUp.mockResolvedValueOnce('mock-token');
    jest.spyOn(Storage.prototype, 'setItem');
    render(<LoginForm />);

    // Act
    fireEvent.click(screen.getByTestId('toggle-sign-up'));
    fireEvent.change(screen.getByLabelText(/email/i), {
      target: { value: correctEmail },
    });
    fireEvent.change(screen.getByLabelText(/password/i), {
      target: { value: correctPassword },
    });
    fireEvent.click(screen.getByTestId('submit-button'));

    // Assert
    await waitFor(() => {
      expect(signUp).toHaveBeenCalledWith(correctEmail, correctPassword);
      expect(localStorage.setItem).toHaveBeenCalledWith('token', 'mock-token');
      expect(pushMock).toHaveBeenCalledWith('/dashboard');
      expect(toast.success).toHaveBeenCalledWith('Sign up successfully');
    });
  });

  // Sign up failed
  it('should show error toast when sign up failed', async () => {
    // Arrange
    signUp.mockRejectedValueOnce(new Error('Email already exists'));
    render(<LoginForm />);

    // Act
    fireEvent.click(screen.getByTestId('toggle-sign-up'));
    fireEvent.change(screen.getByLabelText(/email/i), {
      target: { value: correctEmail },
    });
    fireEvent.change(screen.getByLabelText(/password/i), {
      target: { value: correctPassword },
    });
    fireEvent.click(screen.getByTestId('submit-button'));

    // Assert
    await waitFor(() => {
      expect(signUp).toHaveBeenCalledWith(correctEmail, correctPassword);
      expect(toast.error).toHaveBeenCalledWith('Error: Email already exists');
    });
  });

  // User does not provide email or password
  it('should show error toast when email or password is not provided', async () => {
    // Arrange
    render(<LoginForm />);

    // Act
    fireEvent.click(screen.getByTestId('submit-button'));

    // Assert
    await waitFor(() => {
      expect(toast.error).toHaveBeenCalledWith(
        'Email and password are required',
      );
    });
  });

  // User can toggle password visibility
  it('should toggle password visibility', async () => {
    // Arrange
    render(<LoginForm />);

    // Act
    fireEvent.click(screen.getByTestId('toggle-visibility'));

    // Assert
    await waitFor(() => {
      expect(screen.getByLabelText(/password/i)).toHaveAttribute(
        'type',
        'text',
      );
    });
  });
});
