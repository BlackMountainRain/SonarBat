import { isTokenExpired } from './jwt';

describe('isTokenExpired', () => {
  // returns false for valid token
  it('should return false when the token is valid', () => {
    // Arrange
    const validToken =
      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc2MDAwMDAwMDB9.123';

    // Act
    const result = isTokenExpired(validToken);

    // Assert
    expect(result).toBe(false);
  });

  // returns true for expired token
  it('should return true when the token is expired', () => {
    // Arrange
    const expiredToken =
      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDAwMDAwMDB9.abc123';

    // Act
    const result = isTokenExpired(expiredToken);

    // Assert
    expect(result).toBe(true);
  });

  // returns true for empty token string
  it('should return true when the token string is empty', () => {
    // Arrange
    const emptyToken = '';

    // Act
    const result = isTokenExpired(emptyToken);

    // Assert
    expect(result).toBe(true);
  });

  // returns true for invalid token
  it('should return true when the token is invalid', () => {
    // Arrange
    const invalidToken = 'invalid-token';

    // Act
    const result = isTokenExpired(invalidToken);

    // Assert
    expect(result).toBe(true);
  });

  // returns true for the token with no expiration claim
  it('should return true when the token has no expiration claim', () => {
    // Arrange
    const tokenWithoutExp = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.123';

    // Act
    const result = isTokenExpired(tokenWithoutExp);

    // Assert
    expect(result).toBe(true);
  });
});
