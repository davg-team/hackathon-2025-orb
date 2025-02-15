import Index from "pages/Index"
import Landing from "pages/Landing"
import Header from "features/components/Header"
import {BrowserRouter, Route, Routes} from "react-router-dom"

const Router = () => {
	return (
		<BrowserRouter>
			<Routes>
				{ /*<Route path="/" element={<Index />} /> */ }
				<Route path="/" element={(<>
																	<Header />
																	<Landing />
																	</>)} />
			</Routes>
		</BrowserRouter>
	)
}

export default Router
