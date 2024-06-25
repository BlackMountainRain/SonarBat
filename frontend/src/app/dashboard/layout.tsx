'use client';

import React, { useEffect, useMemo } from 'react';
import { usePathname, useRouter } from 'next/navigation';
import SideBar from '@/app/ui/dashboard/sidebar';
import { SidebarContext } from '@/app/dashboard/layout-context';
import { useLockedBody } from '@/app/hooks/useBodyLock';
import { isTokenExpired } from '@/app/utils/jwt';
import NavBar from '@/app/ui/dashboard/navbar';

const DashboardLayout = ({
  children,
}: {
  children: React.ReactNode;
}): React.ReactNode => {
  const pathname = usePathname();
  const router = useRouter();
  const [sidebarOpen, setSidebarOpen] = React.useState(false);
  const [, setLocked] = useLockedBody(false);
  const handleToggleSidebar = () => {
    setSidebarOpen(!sidebarOpen);
    setLocked(!sidebarOpen);
  };

  useEffect(() => {
    // verify token
    const token = localStorage.getItem('token') ?? '';
    if (isTokenExpired(token)) {
      router.push('/auth');
    }
  }, [pathname]);

  const value = useMemo(
    () => ({
      collapsed: sidebarOpen,
      setCollapsed: handleToggleSidebar,
    }),
    [sidebarOpen, handleToggleSidebar],
  );

  return (
    <SidebarContext.Provider value={value}>
      <div className="flex">
        <SideBar />
        <div className="flex flex-col flex-1">
          <NavBar />
          <div className="grow p-6 md:overflow-y-auto md:p-12">{children}</div>
        </div>
      </div>
    </SidebarContext.Provider>
  );
};

export default DashboardLayout;
