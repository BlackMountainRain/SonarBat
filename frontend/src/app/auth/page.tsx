'use client';

import React from 'react';
import { Toaster } from 'react-hot-toast';
import LoginForm from '@/app/ui/auth/login-form';

const LeftPanel = (): React.ReactNode => (
  <div
    className="border-5 sm:w-[30%] md:w-[40%] lg:w-[50%] h-full rounded-l-3xl border-slate-50
                    dark:border-neutral-900  flex-col-reverse hidden sm:flex"
  >
    <div className="m-5 text-5xl text-slate-50 dark:text-neutral-900 select-none">
      Sonar Bat
    </div>
  </div>
);

const RightPanel = (): React.ReactNode => (
  <div
    className="border-5 w-[100%] sm:w-[70%] md:w-[60%] lg:w-[50%] h-full rounded-3xl sm:rounded-none sm:rounded-r-3xl
                    bg-slate-50 dark:bg-neutral-900 border-slate-50 dark:border-neutral-900
                    flex justify-center items-center"
  >
    <div className="w-[80%] flex flex-col justify-center items-center gap-5">
      <img src="/logo.white.webp" alt="logo" className="w-20 h-20" />
      <div className="text-5xl">Welcome Back</div>
      <div className="text-sm">Please sign in to continue</div>
      <LoginForm />
    </div>
  </div>
);

const LoginPage = (): React.ReactNode => (
  <>
    <div className="h-[100dvh] w-[100dvw] bg-login-bg bg-cover bg-center flex justify-center items-center p-10 min-h-[800px]">
      <LeftPanel />
      <RightPanel />
    </div>

    <Toaster />
  </>
);

export default LoginPage;
