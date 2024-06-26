'use client';

import { createContext, useContext } from 'react';
import { User } from '@/app/types/auth';

interface Context {
  collapsed: boolean;
  setCollapsed: () => void;
  user: User | null;
  setUser: (user: User | null) => void;
}

export const DashboardContext = createContext<Context>({
  collapsed: false,
  setCollapsed: () => {},
  user: null,
  setUser: () => {},
});

export const useDashboardContext = () => {
  return useContext(DashboardContext);
};
