import React from 'react';
import { renderHook } from '@testing-library/react';
import {
  useDashboardContext,
  DashboardContext,
} from '@/app/dashboard/layout-context';

const wrapper = ({ children }: { children: React.ReactNode }) => (
  <DashboardContext.Provider
    value={{
      collapsed: false,
      setCollapsed: () => {},
      user: null,
      setUser: () => {},
    }}
  >
    {children}
  </DashboardContext.Provider>
);

describe('useSidebarContext', () => {
  // returns default context values when no provider is used
  it('should return default context values when no provider is used', () => {
    // Arrange
    const { result } = renderHook(() => useDashboardContext(), { wrapper });

    // Act

    // Assert
    expect(result.current).not.toBeNull();
    expect(result.current).not.toBeUndefined();
    expect(result.current.collapsed).toBe(false);
    expect(result.current.setCollapsed).toBeInstanceOf(Function);
  });
});
