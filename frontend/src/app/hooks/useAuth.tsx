import { useState } from 'react';
import { useRouter } from 'next/navigation';
import * as jose from 'jose';
import { Claims, User } from '@/app/types/auth';

const useAuth = () => {
  const [user, setUser] = useState<User | null>(null);
  const router = useRouter();

  const isTokenExpired = (token: string): boolean => {
    if (!token) {
      return true;
    }

    try {
      const claims = jose.decodeJwt(token);
      const exp = claims.exp ?? 0;
      if (exp === 0 || exp <= Date.now() / 1000) {
        return true;
      }
    } catch (error) {
      return true;
    }
    return false;
  };

  const getUserInfo = (token: string): User | null => {
    if (!token) {
      return null;
    }

    try {
      const claims: Claims = jose.decodeJwt(token);
      const exp = claims.exp ?? 0;
      if (exp === 0 || exp <= Date.now() / 1000) {
        return null;
      }
      return {
        uid: claims.uid,
        username: claims.username,
        email: claims.email,
        status: claims.status,
        avatar_url: claims.avatar_url,
      };
    } catch (error) {
      return null;
    }
  };

  const handleSignIn = (token: string) => {
    localStorage.setItem('token', token);
    router.push('/dashboard');
  };

  const handleSignOut = () => {
    localStorage.removeItem('token');
    router.push('/auth');
  };

  const verifyToken = () => {
    const token = localStorage.getItem('token') ?? '';
    if (isTokenExpired(token)) {
      handleSignOut();
    }
  };

  const refreshUserInfo = () => {
    const token = localStorage.getItem('token') ?? '';
    if (isTokenExpired(token)) {
      handleSignOut();
      return;
    }
    if (user === null) {
      const userInfo = getUserInfo(token);
      if (userInfo === null) {
        handleSignOut();
        return;
      }
      setUser(userInfo);
    }
  };

  return {
    user,
    setUser,
    isTokenExpired,
    getUserInfo,
    handleSignIn,
    handleSignOut,
    refreshUserInfo,
    verifyToken,
  };
};

export default useAuth;
