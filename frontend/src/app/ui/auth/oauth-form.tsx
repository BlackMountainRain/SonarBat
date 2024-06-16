import React, { useEffect } from 'react';
import { useOAuth2 } from '@tasoskakour/react-use-oauth2';
import { Button } from '@nextui-org/react';
import {
  GrAmazon, GrApple, GrFacebook, GrGithub,
  GrGoogle, GrLinkedin, GrTwitter,
} from 'react-icons/gr';
import { FaMicrosoft } from 'react-icons/fa';
import { OAuthProvider } from '@/app/types/auth';
import { fetchOAuthProviders, fetchToken } from '@/app/lib/data';
import toast from '@/app/lib/toast';

const Icons: Record<string, React.ReactNode> = {
  GitHub: <GrGithub size={20} />,
  Google: <GrGoogle size={20} />,
  X: <GrTwitter size={20} />,
  Facebook: <GrFacebook size={20} />,
  Microsoft: <FaMicrosoft size={20} />,
  LinkedIn: <GrLinkedin size={20} />,
  Apple: <GrApple size={20} />,
  Amazon: <GrAmazon size={20} />,
};

const OAuthForm = () => {
  const [providers, setProviders] = React.useState<OAuthProvider[]>([]);

  useEffect(() => {
    (async () => {
      try {
        const data = await fetchOAuthProviders();
        setProviders(data);
      } catch (error) {
        toast.error(`Failed to fetch OAuth providers: ${error}`);
      }
    })();
  }, []);

  if (providers.length === 0) {
    return null;
  }

  const onSuccess = (token: string) => {
    localStorage.setItem('token', token);
    toast.info('Sign in successfully');
  };
  const onError = (err: string) => {
    toast.error(`Sign in failed: ${err}`);
  };

  const handleOAuthClick = (provider: OAuthProvider) => {
    const { getAuth } = useOAuth2<string>({
      authorizeUrl: provider.authorize_url,
      clientId: provider.client_id,
      redirectUri: provider.redirect_uri,
      scope: provider.scope,
      responseType: provider.response_type,
      exchangeCodeForTokenQueryFn: async (callbackParameters): Promise<string> => {
        try {
          return await fetchToken(provider.provider, callbackParameters.code);
        } catch (error) {
          throw new Error(String(error));
        }
      },
      onSuccess,
      onError,
    });
    getAuth();
  };

  return (
    <>
      {
        providers.map((provider) => (
          <Button
            key={provider.provider}
            className="w-full"
            onClick={() => {
              handleOAuthClick(provider);
            }}
            startContent={Icons[provider.provider]}
          >
            Sign in with
            {' '}
            {provider.provider}
          </Button>
        ))
      }
    </>
  );
};

export default OAuthForm;
