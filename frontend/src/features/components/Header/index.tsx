/* eslint-disable @typescript-eslint/ban-ts-comment */
import { useState, useEffect } from "react";
import { Avatar, Flex, Button, Icon, Text, TextInput, Modal } from "@gravity-ui/uikit";
import { Bars, Xmark, ArrowRightToSquare } from "@gravity-ui/icons";
import Cookies from "js-cookie";
import styles from "./index.module.css";
import { Link } from "react-router-dom";
import { isValid } from "shared";

function Header() {
  const [menuOpen, setMenuOpen] = useState(false);
  const [isMobile, setIsMobile] = useState(window.innerWidth < 768);
  const token = Cookies.get("access_token");
  const hasValidToken = token ? isValid(token) : false;
  const [state, setState] = useState({ username: "", password: "" });
  const [open, setOpen] = useState(false);

  let fullName = "";
  if (token && token.split(".").length === 3) {
    try {
      const payload = JSON.parse(atob(token.split(".")[1]));
      fullName = payload.full_name || "";
    } catch (error) {
      console.error("Ошибка декодирования токена:", error);
    }
  }

  useEffect(() => {
    const handleResize = () => {
      setIsMobile(window.innerWidth < 768);
      if (window.innerWidth >= 768) {
        setMenuOpen(false);
      }
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  async function handleLogin() {
    const { username, password } = state;
    const response = await fetch("/api/users/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });
    if (response.ok) {
      window.location.reload();
    } else {
      console.error("Ошибка входа");
    }
  }

  return (
    <Flex className={styles.header} alignItems={"center"} justifyContent={"space-between"} spacing={{ p: "5" }}>
      <Flex className={styles.left} gap={"3"} alignItems={"center"}>
        <Modal open={open} onOpenChange={() => setOpen(false)}>
          <Flex alignItems={"center"} spacing={{ p: "5" }} width={"40vw"} direction={"column"} gap={"3"}>
            <Text variant="header-2">Войдите в аккаунт</Text>
            <a href="https://lk.orb.ru/oauth/authorize_mobile_with_social/esia?response_type=code&client_id=29&scope=email" className={"g-button g-button_view_normal g-button_size_m g-button_pin_round-round g-button_width_max"}>
              <Flex alignItems={"center"} justifyContent={"center"}>
                <Text variant="body-2">Войти через ГИС ЕЛК</Text>
                <Icon data={ArrowRightToSquare} />
              </Flex>
            </a>
            <Text variant="body-2">Или войдите при помощи логина и пароля, полученного от администратора:</Text>
            <TextInput placeholder="Логин" value={state.username} onChange={(event) => setState({ ...state, username: event.target.value })} />
            <TextInput placeholder="Пароль" type="password" value={state.password} onChange={(event) => setState({ ...state, password: event.target.value })} />
            <Button width="max" onClick={handleLogin}>Войти</Button>
          </Flex>
        </Modal>
        <Link to={"/"}>
          <Avatar imgUrl="/logo.svg" size="xl" />
        </Link>
        {!isMobile && (
          <div className={styles.desktopMenu}>
            <Link to={"/#about"} className={styles.link} onClick={() => setMenuOpen(false)}>о проекте</Link>
            <Link to={"/#participants"} className={styles.link} onClick={() => setMenuOpen(false)}>участники</Link>
            <Link to={"/#conflicts"} className={styles.link} onClick={() => setMenuOpen(false)}>военные конфликты</Link>
            <Link to={'/person/add'} className={styles.link} onClick={() => setMenuOpen(false)}>+ добавить участника</Link>
          </div>
        )}
      </Flex>
      {menuOpen && isMobile && (
        <div className={styles.mobileMenu}>
          <Link to={"/#about"} className={styles.link} onClick={() => setMenuOpen(false)}>о проекте</Link>
          <Link to={"/#participants"} className={styles.link} onClick={() => setMenuOpen(false)}>участники</Link>
          <Link to={"/#conflicts"} className={styles.link} onClick={() => setMenuOpen(false)}>военные конфликты</Link>
          <Link to={'/person/add'} className={styles.link} onClick={() => setMenuOpen(false)}>+ добавить участника</Link>
        </div>
      )}
      <Flex alignItems={"center"}>
        {token && hasValidToken ? (
          <Link to="/profile">
            <Avatar size="xl" text={fullName} />
          </Link>
        ) : (
          <div onClick={() => setOpen(true)}>
            <Avatar icon={ArrowRightToSquare} size="xl" />
          </div>
        )}
        {isMobile && (
          <Button onClick={() => setMenuOpen(!menuOpen)}>
            {menuOpen ? <Icon data={Xmark} /> : <Icon data={Bars} />}
          </Button>
        )}
      </Flex>
    </Flex>
  );
}

export default Header;