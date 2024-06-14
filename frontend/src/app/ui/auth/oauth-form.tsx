import React, { useEffect } from 'react';
import { useOAuth2 } from '@tasoskakour/react-use-oauth2';
import { Button } from '@nextui-org/react';
import { GrGithub, GrGoogle } from 'react-icons/gr';
import { OAuthProvider } from '@/app/types/auth';
import { fetchOAuthProviders, fetchToken } from '@/app/lib/data';
import toast from '@/app/lib/toast';

const Icons: Record<string, React.ReactNode> = {
  github: <GrGithub size={20} />,
  google: <GrGoogle size={20} />,
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
    const { getAuth } = useOAuth2({
      authorizeUrl: provider.authorize_url,
      clientId: provider.client_id,
      redirectUri: provider.redirect_uri,
      scope: provider.scope,
      responseType: provider.response_type,
      exchangeCodeForTokenQueryFn: async (callbackParameters) => {
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
