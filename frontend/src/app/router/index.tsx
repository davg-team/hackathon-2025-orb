import Landing from "pages/Landing";
import Header from "features/components/Header";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Search from "pages/Search";
import Person from "pages/Person";
import User from "pages/User";
import Admin from "pages/Admin";
import UserAdd from "pages/UserAdd";

const Router = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={
            <>
              <Header />
              <Landing />
            </>
          }
        />
        <Route
          path="/search"
          element={
            <>
              <Header />
              <Search />
            </>
          }
        />
        <Route
          path="/person/:id"
          element={
            <>
              <Header />
              <Person />
            </>
          }
        />
				{/* early added people */}
        <Route
          path="/profile"
          element={
            <>
              <Header />
              <User />
            </>
          }
        />
				{/* add new people into system */}
        <Route
          path="/profile/:id/add"
          element={
            <>
              <Header />
              <UserAdd />
            </>
          }
        />
				{/* users for aprove */}
        <Route
          path="/admin/:id"
          element={
            <>
              <Header />
              <Admin />
            </>
          }
        />
				{/* requests from users */}
        <Route
          path="/admin/:id/requests"
          element={
            <>
              <Header />
              <Person />
            </>
          }
        />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
