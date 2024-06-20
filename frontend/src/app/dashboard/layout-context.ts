'use client';

import { createContext, useContext } from 'react';

interface ISidebarContext {
  collapsed: boolean;
  setCollapsed: () => void;
}

export const SidebarContext = createContext<ISidebarContext>({
  collapsed: false,
  setCollapsed: () => {},
});

export const useSidebarContext = () => {
  return useContext(SidebarContext);
};
