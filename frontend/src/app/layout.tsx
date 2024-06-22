import type { Metadata } from 'next';
import { Lusitana } from 'next/font/google';
import { NextUIProvider } from '@nextui-org/react';
import './globals.css';
import React from 'react';
import { ThemeProvider } from 'next-themes';

const lusitana = Lusitana({
  weight: ['400', '700'],
  subsets: ['latin'],
});

export const metadata: Metadata = {
  title: 'Sonar Bat',
  description: 'Sonar Bat is a large-scale network detection system',
};

const Providers = ({ children }: { children: React.ReactNode }) => (
  <NextUIProvider>
    <ThemeProvider attribute="class" defaultTheme="dark">
      {children}
    </ThemeProvider>
  </NextUIProvider>
);

const RootLayout = ({
  children,
}: {
  children: React.ReactNode;
}): React.ReactNode => (
  <html lang="en">
    <body className={lusitana.className} suppressHydrationWarning>
      <Providers>{children}</Providers>
    </body>
  </html>
);

export default RootLayout;
export { Providers };
