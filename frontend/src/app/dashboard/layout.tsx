'use client';

import React, { useEffect, useMemo, useState } from 'react';
import { usePathname } from 'next/navigation';
import SideBar from '@/app/ui/dashboard/sidebar';
import { DashboardContext } from '@/app/dashboard/layout-context';
import { useLockedBody } from '@/app/hooks/useBodyLock';
import NavBar from '@/app/ui/dashboard/navbar';
import useAuth from '@/app/hooks/useAuth';

const DashboardLayout = ({
  children,
}: {
  children: React.ReactNode;
}): React.ReactNode => {
  const pathname = usePathname();
  const [sidebarOpen, setSidebarOpen] = useState(false);
  const [, setLocked] = useLockedBody(false);
  const handleToggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
    setLocked(!sidebarOpen);
  };
  const { user, setUser, refreshUserInfo } = useAuth();

  useEffect(() => {
    refreshUserInfo();
  }, [pathname]);

  const value = useMemo(
    () => ({
      collapsed: sidebarOpen,
      setCollapsed: handleToggleSidebar,
      user,
      setUser,
    }),
    [sidebarOpen, handleToggleSidebar],
  );

  return (
    <DashboardContext.Provider value={value}>
      <div className="flex">
        <SideBar />
        <div className="flex flex-col flex-1">
          <NavBar />
          <div className="grow p-6 md:overflow-y-auto md:p-12">{children}</div>
        </div>
      </div>
    </DashboardContext.Provider>
  );
};

export default DashboardLayout;
