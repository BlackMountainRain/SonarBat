'use client';

import React, { useContext } from 'react';
import { usePathname } from 'next/navigation';
import { Image } from '@nextui-org/react';
import { HiBugAnt, HiChartPie, HiMiniCube } from 'react-icons/hi2';
import { useTheme } from 'next-themes';
import { Sidebar } from '@/app/ui/dashboard/sidebar.style';
import { SidebarContext } from '@/app/dashboard/layout-context';
import { SidebarItem } from '@/app/ui/dashboard/sidebar-item';
import { SidebarMenu } from '@/app/ui/dashboard/sidebar-menu';

const links = [
  { name: 'Home', href: '/dashboard', icon: HiChartPie },
  {
    isGroup: true,
    name: 'Tasks',
    icon: HiBugAnt,
    items: [
      {
        name: 'Tasks',
        href: '/dashboard/tasks',
        icon: HiBugAnt,
      },
    ],
  },
  {
    isGroup: true,
    name: 'Resources',
    icon: HiMiniCube,
    items: [
      {
        name: 'Resources',
        href: '/dashboard/resources',
        icon: HiMiniCube,
      },
    ],
  },
];

const SideBar = (): React.ReactNode => {
  const { theme } = useTheme();
  const pathname = usePathname();
  const { collapsed, setCollapsed } = useContext(SidebarContext);

  return (
    <aside className="h-screen z-[20] sticky top-0 select-none">
      {collapsed ? (
        <div className={Sidebar.Overlay()} onClick={setCollapsed} />
      ) : null}
      <div
        className={Sidebar({
          collapsed,
        })}
      >
        <div className={Sidebar.Header()}>
          {theme === 'light' ? (
            <Image
              width={40}
              height={40}
              src="/logo.black.webp"
              className="select-none hover:scale-105"
            />
          ) : (
            <Image
              width={40}
              height={40}
              src="/logo.white.webp"
              className="hover:scale-105"
            />
          )}

          <div>Sonar Bat</div>
        </div>
        <div className="flex flex-col justify-between h-full">
          <div className={Sidebar.Body()}>
            {links.map((link) => {
              if (link.isGroup) {
                return (
                  <SidebarMenu key={link.name} title={link.name}>
                    {link.items.map((item) => (
                      <SidebarItem
                        key={item.name}
                        isActive={pathname === item.href}
                        title={item.name}
                        icon={<item.icon size={24} />}
                        href={item.href}
                      />
                    ))}
                  </SidebarMenu>
                );
              }
              return (
                <SidebarItem
                  key={link.name}
                  isActive={pathname === link.href}
                  title={link.name}
                  icon={<link.icon size={24} />}
                  href={link.href}
                />
              );
            })}
          </div>
        </div>
      </div>
    </aside>
  );
};

export default SideBar;
