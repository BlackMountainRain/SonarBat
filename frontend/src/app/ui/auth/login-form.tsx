import React from 'react';
import { Button, Input } from '@nextui-org/react';
import { FaRegEye, FaRegEyeSlash } from 'react-icons/fa6';
import { Spacer } from '@nextui-org/spacer';
import toast from '@/app/lib/toast';
import { fetchToken, signUp } from '@/app/lib/data';
import OAuthForm from '@/app/ui/auth/oauth-form';

const LoginForm = () => {
  const [email, setEmail] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [isVisible, setIsVisible] = React.useState(false);
  const [isSignUp, setIsSignUp] = React.useState(false);

  const toggleVisibility = () => setIsVisible(!isVisible);
  const toggleSignUp = () => setIsSignUp(!isSignUp);

  const handleSignUp = async () => {
    try {
      const token = await signUp(email, password);
      localStorage.setItem('token', token);
    } catch (error) {
      toast.error(String(error));
    }
  };

  const handleSignIn = async () => {
    try {
      const token = await fetchToken('Self', '', email, password);
      localStorage.setItem('token', token);
    } catch (error) {
      toast.error(String(error));
    }
  };

  return (
    <>
      <Input
        key="email"
        type="email"
        label="Email"
        labelPlacement="outside"
        placeholder="Please enter your email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      <Input
        key="password"
        type={isVisible ? 'text' : 'password'}
        label="Password"
        labelPlacement="outside"
        placeholder="Please enter your password"
        endContent={(
          <button className="focus:outline-none" type="button" onClick={toggleVisibility}>
            {isVisible ? (
              <FaRegEyeSlash className="text-2xl text-default-400 pointer-events-none" />
            ) : (
              <FaRegEye className="text-2xl text-default-400 pointer-events-none" />
            )}
          </button>
        )}
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      <Spacer y={1} />

      <Button
        className="w-full dark:bg-blue-800"
        onClick={() => {
          isSignUp ? handleSignUp() : handleSignIn();
        }}
      >
        {isSignUp ? 'Sign Up' : 'Sign In'}
      </Button>

      <OAuthForm />

      <div className="flex items-center gap-x-2">
        <div>
          {isSignUp ? 'Already have an account?' : 'Don\'t have an account?'}
        </div>
        <div
          className="font-bold text-blue-700 dark:text-blue-500 cursor-pointer hover:underline"
          onClick={toggleSignUp}
        >
          {isSignUp ? 'Sign In' : 'Sign Up'}
        </div>
      </div>
    </>
  );
};

export default LoginForm;
