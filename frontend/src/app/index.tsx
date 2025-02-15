import {ThemeProvider} from '@gravity-ui/uikit'
import './styles/index.css'
import Router from './router'
import ContextProvider from './context'

const App = () => {
  return (
    <ThemeProvider theme="light">
      <ContextProvider>
        <Router />
      </ContextProvider>
    </ThemeProvider>
  );
};

export default App
