import { createSignal, createContext, useContext, JSX } from "solid-js";
const LoggedInContext = createContext();

export function CounterProvider(props: {
  loggedIn: boolean;
  children:
    | number
    | boolean
    | Node
    | JSX.ArrayElement
    | (string & {})
    | null
    | undefined;
}) {
  const [loggedIn, setLoggedIn] = createSignal<boolean>(
      props.loggedIn || false
    ),
    counter = [
      loggedIn,
      {
        increment() {
          setLoggedIn(true);
        },
        decrement() {
          setLoggedIn(false);
        },
      },
    ];

  return (
    <LoggedInContext.Provider value={counter}>
      {props.children}
    </LoggedInContext.Provider>
  );
}

export function useCounter() {
  return useContext(LoggedInContext);
}
