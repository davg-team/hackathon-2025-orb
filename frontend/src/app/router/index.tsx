import Landing from "pages/Landing";
import Header from "features/components/Header";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Search from "pages/Search";
import Person from "pages/Person";
import User from "pages/User";
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
              <Person mode="base" />
            </>
          }
        />
        <Route
          path="/person-print/:id"
          element={
            <>
              <Person mode="print"/>
            </>
          }
        />
        <Route
          path="/profile"
          element={
            <>
              <Header />
              <User />
            </>
          }
        />
        <Route
          path="/person/add"
          element={
            <>
              <Header />
              <UserAdd />
            </>
          }
        />
      </Routes>
    </BrowserRouter>
  );
};

export default Router;
