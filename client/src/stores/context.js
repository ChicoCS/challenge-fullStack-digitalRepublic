import React, { useContext } from "react";

const StoreContext = React.createContext(null);

export const StoreProvider = StoreContext.Provider;

export const StoreConsumer = StoreContext.Consumer;

export const withStore = (WrappedComponent) => (props) => (
  <StoreConsumer>{(stores) => <WrappedComponent {...props} {...stores} />}</StoreConsumer>
);

export const useStore = (name) => {
  const context = useContext(StoreContext);
  return context[name];
};

export default StoreContext;