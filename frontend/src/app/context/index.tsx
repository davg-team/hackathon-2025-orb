import Context from "./Context"

const ContextProvider = ({children}: {children: React.ReactNode}) => {
	return (
		<Context.Provider value={{theme: 'light'}}>
			{children}
		</Context.Provider>
	)
}

export default ContextProvider
