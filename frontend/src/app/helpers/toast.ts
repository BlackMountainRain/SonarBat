import toast from 'react-hot-toast';

const idMap = new Map<string, string>();

export default {
  stickyError: (id: string, msg: string) => {
    if (idMap.has(id)) return; // prevent duplicate toast
    const toastId = toast.loading(msg, {
      style: {
        background: '#ef4444',
        color: '#fff',
        lineBreak: 'anywhere',
      },
      iconTheme: {
        primary: '#ef4444',
        secondary: '#fff',
      },
    });
    idMap.set(id, toastId);
  },
  stickySuccess: (id: string, msg: string) => {
    const toastId = idMap.get(id);
    if (toastId === undefined) return;
    if (toastId) {
      toast.success(msg, {
        id: toastId,
        style: {
          background: '#333',
          color: '#fff',
          lineBreak: 'anywhere',
        },
        iconTheme: {
          primary: '#333',
          secondary: '#fff',
        },
      });
    }
    idMap.delete(id);
  },

  info: (msg: string) => {
    toast(msg, {
      position: 'top-center',
      className: 'bg-neutral-700',

      style: {
        background: '#333',
        color: '#fff',
        lineBreak: 'anywhere',
      },
    });
  },

  success: (msg: string) => {
    toast.success(msg, {
      position: 'top-center',
      className: 'bg-neutral-700',

      style: {
        background: '#333',
        color: '#fff',
        lineBreak: 'anywhere',
      },
    });
  },

  error: (msg: string) => {
    toast.error(msg, {
      position: 'top-center',
      className: 'bg-red-700',

      style: {
        background: '#333',
        color: '#fff',
        lineBreak: 'anywhere',
      },
    });
  },
};
