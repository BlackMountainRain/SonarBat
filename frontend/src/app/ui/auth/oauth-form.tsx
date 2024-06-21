import React from 'react';
import { useOAuth2 } from '@tasoskakour/react-use-oauth2';
import { Button } from '@nextui-org/react';
import { GrGithub, GrGoogle } from 'react-icons/gr';
import { useRouter } from 'next/navigation';
import { fetchToken } from '@/app/lib/data';
import toast from '@/app/lib/toast';

const OAuthForm = () => {
  const router = useRouter();

  const onSuccess = (token: string) => {
    localStorage.setItem('token', token);
    router.push('/dashboard');
    toast.info('Sign in successfully');
  };
  const onError = (err: string) => {
    toast.error(`Sign in failed: ${err}`);
  };

  const { getAuth: getGitHubAuth } = useOAuth2<string>({
    authorizeUrl: 'https://github.com/login/oauth/authorize',
    clientId: process.env.NEXT_PUBLIC_GITHUB_CLIENT_ID || '',
    redirectUri: `${process.env.NEXT_PUBLIC_BASE_URL || ''}/auth/callback`,
    scope: 'read:user',
    responseType: 'code',
    exchangeCodeForTokenQueryFn: async (
      callbackParameters,
    ): Promise<string> => {
      try {
        return await fetchToken('GITHUB', callbackParameters.code);
      } catch (error) {
        throw new Error(String(error));
      }
    },
    onSuccess,
    onError,
  });

  const { getAuth: getGoogleAuth } = useOAuth2<string>({
    authorizeUrl: 'https://accounts.google.com/o/oauth2/auth',
    clientId: process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || '',
    redirectUri: `${process.env.NEXT_PUBLIC_BASE_URL || ''}/auth/callback`,
    scope: 'openid profile email',
    responseType: 'code',
    exchangeCodeForTokenQueryFn: async (
      callbackParameters,
    ): Promise<string> => {
      try {
        return await fetchToken('GOOGLE', callbackParameters.code);
      } catch (error) {
        throw new Error(String(error));
      }
    },
    onSuccess,
    onError,
  });

  return (
    <>
      <Button
        className="w-full"
        onClick={getGitHubAuth}
        startContent={<GrGithub size={20} />}
      >
        Sign in with GitHub
      </Button>
      <Button
        className="w-full"
        onClick={getGoogleAuth}
        startContent={<GrGoogle size={20} />}
      >
        Sign in with Google
      </Button>
    </>
  );
};

export default OAuthForm;
