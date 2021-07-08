import create from "zustand";

interface LoggedInStore {
  loggedIn: boolean;
  loading: boolean;
  setStatus: (loggedIn: boolean) => void;
}

export const useLoggedIn = create<LoggedInStore>((set) => ({
  loggedIn: false,
  loading: true,
  setStatus: (loggedIn) =>
    set((_) => ({
      loggedIn,
      loading: false,
    })),
}));
