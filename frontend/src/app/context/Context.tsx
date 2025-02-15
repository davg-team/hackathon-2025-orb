import {createContext} from "react";

interface IContext {
	theme: 'light' | 'dark';
}

const Context = createContext<IContext>({theme: 'light'})
export default Context
