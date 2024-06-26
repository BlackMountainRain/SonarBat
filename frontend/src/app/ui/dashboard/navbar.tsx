import React, { useContext } from 'react';
import { Input, Navbar, NavbarContent } from '@nextui-org/react';
import { GrGithub } from 'react-icons/gr';
import { CiSearch } from 'react-icons/ci';
import { HiBars3BottomLeft, HiBars3BottomRight } from 'react-icons/hi2';
import ThemeToggleButton from '@/app/ui/dashboard/theme-switcher';
import UserDropdown from '@/app/ui/dashboard/user-dropdown';
import { DashboardContext } from '@/app/dashboard/layout-context';

const NavBar = (): React.ReactNode => {
  const { collapsed, setCollapsed } = useContext(DashboardContext);

  return (
    <Navbar
      isBordered
      className="w-full max-w-full select-none z-[20] sticky top-0"
      classNames={{
        wrapper: 'w-full max-w-full',
      }}
    >
      <NavbarContent className="w-full">
        {collapsed ? (
          <HiBars3BottomRight
            size={24}
            className="cursor-pointer hover:opacity-80 md:hidden"
            onClick={setCollapsed}
          />
        ) : (
          <HiBars3BottomLeft
            size={24}
            className="cursor-pointer hover:opacity-80 md:hidden"
            onClick={setCollapsed}
          />
        )}

        <Input
          startContent={<CiSearch />}
          isClearable
          className="w-full max-w-full"
          classNames={{
            input: 'w-full max-w-full',
            mainWrapper: 'w-full max-w-full',
          }}
          placeholder="Quick Search..."
        />
      </NavbarContent>
      <NavbarContent
        justify="end"
        className="w-fit data-[justify=end]:flex-grow-0"
      >
        <ThemeToggleButton />

        <GrGithub
          size={24}
          className="cursor-pointer hover:opacity-80"
          onClick={() =>
            window.open(
              'https://github.com/BlackMountainRain/SonarBat',
              '_blank',
            )
          }
        />
        <NavbarContent>
          <UserDropdown />
        </NavbarContent>
      </NavbarContent>
    </Navbar>
  );
};

export default NavBar;
