import {
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
  NavbarItem,
  User,
} from '@nextui-org/react';
import React, { useContext } from 'react';
import { DashboardContext } from '@/app/dashboard/layout-context';

const UserDropdown = () => {
  const { user, handleSignOut } = useContext(DashboardContext);

  return (
    <Dropdown>
      <NavbarItem>
        <DropdownTrigger>
          <User
            name={user?.username}
            description={user?.email}
            avatarProps={{
              size: 'sm',
              src: user?.avatar_url,
            }}
            className="cursor-pointer hover:opacity-80"
          />
        </DropdownTrigger>
      </NavbarItem>

      <DropdownMenu aria-label="User menu actions" className="select-none">
        <DropdownItem
          key="profile"
          className="flex flex-col justify-start w-full items-start"
        >
          <p>Signed in as</p>
          <p>{user?.email}</p>
        </DropdownItem>
        <DropdownItem key="settings">My Settings</DropdownItem>
        <DropdownItem key="help_and_feedback">Help & Feedback</DropdownItem>
        <DropdownItem
          key="logout"
          color="danger"
          className="text-danger"
          onClick={handleSignOut}
        >
          Log Out
        </DropdownItem>
      </DropdownMenu>
    </Dropdown>
  );
};

export default UserDropdown;
