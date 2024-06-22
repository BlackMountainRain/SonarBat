import React from 'react';
import { renderHook } from '@testing-library/react';
import {
  useSidebarContext,
  SidebarContext,
} from '@/app/dashboard/layout-context';

const wrapper = ({ children }: { children: React.ReactNode }) => (
  <SidebarContext.Provider value={{ collapsed: false, setCollapsed: () => {} }}>
    {children}
  </SidebarContext.Provider>
);

describe('useSidebarContext', () => {
  // returns default context values when no provider is used
  it('should return default context values when no provider is used', () => {
    // Arrange
    const { result } = renderHook(() => useSidebarContext(), { wrapper });

    // Act

    // Assert
    expect(result.current).not.toBeNull();
    expect(result.current).not.toBeUndefined();
    expect(result.current.collapsed).toBe(false);
    expect(result.current.setCollapsed).toBeInstanceOf(Function);
  });
});
