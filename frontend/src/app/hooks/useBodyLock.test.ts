import { act, renderHook } from '@testing-library/react';
import { useLockedBody } from '@/app/hooks/useBodyLock';

describe('useLockedBody', () => {
  beforeEach(() => {
    // Reset the body styles before each test
    document.body.style.overflow = '';
    document.body.style.paddingRight = '';
  });

  it('should lock the body scroll when locked is true', () => {
    const { result } = renderHook(() => useLockedBody(true));

    expect(document.body.style.overflow).toBe('hidden');

    const root = document.getElementById('___gatsby');
    const scrollBarWidth = root ? root.offsetWidth - root.scrollWidth : 0;

    if (scrollBarWidth) {
      expect(document.body.style.paddingRight).toBe(`${scrollBarWidth}px`);
    } else {
      expect(document.body.style.paddingRight).toBe('');
    }

    act(() => {
      result.current[1](false);
    });

    expect(document.body.style.overflow).toBe('');
    expect(document.body.style.paddingRight).toBe('');
  });

  it('should unlock the body scroll when locked is false', () => {
    const { result } = renderHook(() => useLockedBody(false));

    expect(document.body.style.overflow).toBe('');
    expect(document.body.style.paddingRight).toBe('');

    act(() => {
      result.current[1](true);
    });

    expect(document.body.style.overflow).toBe('hidden');

    const root = document.getElementById('___gatsby');
    const scrollBarWidth = root ? root.offsetWidth - root.scrollWidth : 0;

    if (scrollBarWidth) {
      expect(document.body.style.paddingRight).toBe(`${scrollBarWidth}px`);
    } else {
      expect(document.body.style.paddingRight).toBe('');
    }
  });

  it('should update state if initialLocked changes', () => {
    const { result, rerender } = renderHook(
      ({ initialLocked }) => useLockedBody(initialLocked),
      { initialProps: { initialLocked: false } },
    );

    expect(result.current[0]).toBe(false);
    expect(document.body.style.overflow).toBe('');

    rerender({ initialLocked: true });

    expect(result.current[0]).toBe(true);
    expect(document.body.style.overflow).toBe('hidden');

    const root = document.getElementById('___gatsby');
    const scrollBarWidth = root ? root.offsetWidth - root.scrollWidth : 0;

    if (scrollBarWidth) {
      expect(document.body.style.paddingRight).toBe(`${scrollBarWidth}px`);
    } else {
      expect(document.body.style.paddingRight).toBe('');
    }
  });
});
