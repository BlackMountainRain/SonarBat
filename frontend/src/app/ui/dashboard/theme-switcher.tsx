import React from 'react';
import { useTheme } from 'next-themes';
import { IoMdMoon } from 'react-icons/io';
import { PiSunDimFill } from 'react-icons/pi';

const ThemeToggleButton = () => {
  const { theme, setTheme } = useTheme();

  if (theme === 'light') {
    return (
      <IoMdMoon
        size={24}
        onClick={() => setTheme('dark')}
        className="cursor-pointer hover:opacity-80"
      />
    );
  }
  return (
    <PiSunDimFill
      size={24}
      onClick={() => setTheme('light')}
      className="cursor-pointer hover:opacity-80"
    />
  );
};

export default ThemeToggleButton;
