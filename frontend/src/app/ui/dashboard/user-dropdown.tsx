import {
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
  NavbarItem,
  User,
} from '@nextui-org/react';
import React from 'react';

const UserDropdown = () => {
  return (
    <Dropdown>
      <NavbarItem>
        <DropdownTrigger>
          <User
            name="Leo"
            description="ifuryst@gmail.com"
            avatarProps={{
              size: 'sm',
              src: 'https://i.pravatar.cc/150',
            }}
            className="cursor-pointer hover:opacity-80"
          />
        </DropdownTrigger>
      </NavbarItem>

      <DropdownMenu
        aria-label="User menu actions"
        className="select-none"
        onAction={(actionKey) => console.log({ actionKey })}
      >
        <DropdownItem
          key="profile"
          className="flex flex-col justify-start w-full items-start"
        >
          <p>Signed in as</p>
          <p>ifuryst@gmail.com</p>
        </DropdownItem>
        <DropdownItem key="settings">My Settings</DropdownItem>
        <DropdownItem key="help_and_feedback">Help & Feedback</DropdownItem>
        <DropdownItem key="logout" color="danger" className="text-danger ">
          Log Out
        </DropdownItem>
      </DropdownMenu>
    </Dropdown>
  );
};

export default UserDropdown;
