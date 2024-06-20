'use client';

import React, { useMemo } from 'react';
import SideBar from '@/app/ui/dashboard/sidebar';
import { SidebarContext } from '@/app/dashboard/layout-context';
import { useLockedBody } from '@/app/hooks/useBodyLock';

const DashboardLayout = ({
  children,
}: {
  children: React.ReactNode;
}): React.ReactNode => {
  const [sidebarOpen, setSidebarOpen] = React.useState(false);
  const [, setLocked] = useLockedBody(false);
  const handleToggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
    setLocked(!sidebarOpen);
  };

  const value = useMemo(
    () => ({
      collapsed: sidebarOpen,
      setCollapsed: handleToggleSidebar,
    }),
    [sidebarOpen, handleToggleSidebar],
  );

  return (
    <SidebarContext.Provider value={value}>
      <div className="flex h-screen flex-col md:flex-row md:overflow-hidden">
        <div className="w-full flex-none md:w-64">
          <SideBar />
        </div>
        <div className="grow p-6 md:overflow-y-auto md:p-12">{children}</div>
      </div>
    </SidebarContext.Provider>
  );
};

export default DashboardLayout;
