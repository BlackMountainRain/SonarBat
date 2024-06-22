import toast from 'react-hot-toast';
import toastNotifications from './toast';

jest.mock('react-hot-toast', () => jest.fn());

describe('toastNotifications', () => {
  beforeAll(() => {
    toast.loading = jest.fn(() => 'fake toast id');
    toast.success = jest.fn();
    toast.error = jest.fn();
  });

  beforeEach(() => {
    jest.clearAllMocks();
  });

  test('stickyError should display a loading toast and prevent duplicates', () => {
    const id = 'random1';
    const msg = 'Error message';

    toastNotifications.stickyError(id, msg);
    expect(toast.loading).toHaveBeenCalledWith(msg, expect.any(Object));
    expect(toast.loading).toHaveBeenCalledTimes(1);

    toastNotifications.stickyError(id, msg);
    expect(toast.loading).toHaveBeenCalledTimes(1); // No duplicate toast
  });

  test('stickySuccess should display a success toast', () => {
    const id = 'random2';
    const msg = 'Success message';

    toastNotifications.stickyError(id, msg);
    expect(toast.loading).toHaveBeenCalledWith(msg, expect.any(Object));
    expect(toast.loading).toHaveBeenCalledTimes(1);

    toastNotifications.stickySuccess(id, msg);
    expect(toast.success).toHaveBeenCalledWith(msg, expect.any(Object));
    expect(toast.success).toHaveBeenCalledTimes(1);
  });

  test('stickySuccess should not display a toast if the id is not found', () => {
    const id = 'random3';
    const msg = 'Success message';

    toastNotifications.stickySuccess(id, msg);
    expect(toast.success).not.toHaveBeenCalled();
  });

  test('success should display a success toast', () => {
    const msg = 'Success message';

    toastNotifications.success(msg);
    expect(toast.success).toHaveBeenCalledWith(msg, expect.any(Object));
    expect(toast.success).toHaveBeenCalledTimes(1);
  });

  test('error should display an error toast', () => {
    const msg = 'Error message';

    toastNotifications.error(msg);
    expect(toast.error).toHaveBeenCalledWith(msg, expect.any(Object));
    expect(toast.error).toHaveBeenCalledTimes(1);
  });

  test('info should display an info toast', () => {
    const msg = 'Info message';

    toastNotifications.info(msg);
    expect(toast).toHaveBeenCalledWith(msg, expect.any(Object));
    expect(toast).toHaveBeenCalledTimes(1);
  });
});
