import { signUp, fetchToken } from './data';

const correctEmail = 'test@example.com';
const correctPassword = 'password123';
const mockToken = 'mock-token';

describe('fetchToken', () => {
  // fetches token successfully with SELF provider using email and password
  it('should fetch token successfully when using SELF provider with email and password', async () => {
    // Arrange
    const mockResponse = { token: mockToken };
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 200,
        json: () => Promise.resolve(mockResponse),
      }),
    ) as jest.Mock;

    // Act
    const token = await fetchToken('SELF', '', correctEmail, correctPassword);

    // Assert
    expect(token).toBe(mockResponse.token);
  });

  // handles non-200 response status gracefully
  it('should throw an error when response status is not 200', async () => {
    // Arrange
    const mockErrorMessage = { message: 'Invalid credentials' };
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 401,
        json: () => Promise.resolve(mockErrorMessage),
      }),
    ) as jest.Mock;

    // Act & Assert
    await expect(
      fetchToken('SELF', '', correctEmail, 'wrong-password'),
    ).rejects.toThrow('Invalid credentials');
  });

  // fetches token successfully with OAuth provider using code
  it('should fetch token successfully when using OAuth provider with code', async () => {
    // Arrange
    const mockResponse = { token: mockToken };
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 200,
        json: () => Promise.resolve(mockResponse),
      }),
    ) as jest.Mock;
    const provider = 'GOOGLE';
    const code = 'mock-code';

    // Act
    const token = await fetchToken(provider, code, '', '');

    // Assert
    expect(token).toBe(mockResponse.token);
  });

  // throws an error when parsing JSON fails, even if the response status is 200.
  it('should throw an error when response.json() throws an error', async () => {
    // Arrange
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 200,
        json: () => {
          throw new Error('Failed to parse JSON');
        },
      }),
    ) as jest.Mock;

    // Act & Assert
    await expect(
      fetchToken('SELF', '', correctEmail, correctPassword),
    ).rejects.toThrow('Failed to parse JSON');
  });

  // throws an error when response.json() throws an error, even if the response status is not 200.
  it('should throw an error when response.json() throws an error, even if the response status is not 200', async () => {
    // Arrange
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 401,
        json: () => {
          throw new Error('Failed to parse JSON');
        },
      }),
    ) as jest.Mock;

    // Act & Assert
    await expect(
      fetchToken('SELF', '', correctEmail, correctPassword),
    ).rejects.toThrow('Get token failed.');
  });
});

describe('signUp', () => {
  // successful sign-up returns a token
  it('should return a token when sign-up is successful', async () => {
    // Arrange
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 200,
        json: () => Promise.resolve({ token: mockToken }),
      }),
    ) as jest.Mock;

    // Act
    const result = await signUp(correctEmail, correctPassword);

    // Assert
    expect(result).toBe(mockToken);
  });

  // response status is not 200
  it('should throw an error when response status is not 200', async () => {
    // Arrange
    const errorMessage = 'Sign up failed.';
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 400,
        json: () => Promise.resolve({ message: errorMessage }),
      }),
    ) as jest.Mock;

    // Act & Assert
    await expect(signUp(correctEmail, correctPassword)).rejects.toThrow(
      errorMessage,
    );
  });

  // throws an error when parsing JSON fails
  it('should throw an error when response.json() throws an error', async () => {
    // Arrange
    global.fetch = jest.fn(() =>
      Promise.resolve({
        status: 400,
        json: () => {
          throw new Error('Failed to parse JSON');
        },
      }),
    ) as jest.Mock;

    // Act & Assert
    await expect(signUp(correctEmail, correctPassword)).rejects.toThrow(
      'Sign up failed.',
    );
  });
});
