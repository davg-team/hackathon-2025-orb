import { Button, Flex, Icon, Text, TextInput } from "@gravity-ui/uikit";
import { Magnifier } from "@gravity-ui/icons";
import styles from "./index.module.css";
import Map from "features/components/Map";
import { Link, useLocation, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

function Landing() {
  const { hash } = useLocation();
  const navigate = useNavigate()
  const [state, setState] = useState({
    name: '',
    middleName: '',
    lastName: ''
  })

  // Cookies.set('token', 'eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJlbWFpbCI6ImFkbWluIiwicm9sZSI6InJvb3QiLCJzbmlscyI6IjAwMC0wMDAtMDAwIDAwIiwiZnVsbF9uYW1lIjoiXHUwNDIxXHUwNDQzXHUwNDNmXHUwNDM1XHUwNDQwXHUwNDMwXHUwNDM0XHUwNDNjXHUwNDM4XHUwNDNkIFx1MDQyMVx1MDQ0M1x1MDQzZlx1MDQzNVx1MDQ0MFx1MDQzMFx1MDQzNFx1MDQzY1x1MDQzOFx1MDQzZCIsIm9yZ2FuaXphdGlvbiI6Ilx1MDQxMFx1MDQzNFx1MDQzY1x1MDQzOFx1MDQzZFx1MDQzOFx1MDQ0MVx1MDQ0Mlx1MDQ0MFx1MDQzMFx1MDQ0Nlx1MDQzOFx1MDQ0ZiIsInBvc2l0aW9uIjoiXHUwNDIxXHUwNDQzXHUwNDNmXHUwNDM1XHUwNDQwXHUwNDMwXHUwNDM0XHUwNDNjXHUwNDM4XHUwNDNkIiwicGhvbmUiOiIiLCJleHAiOjE3Mzk3MTE4MzR9.nrPwinQGm2baLF8LnDe4mmkh5o46wNkAjYdWU_DBCnunda8Kr4XDH9YtagFK5YC4-djfOekv1a4KtiT5J4Pho2kSRGA5Xx0ZASjRJXNlW0kpAHBfCmJeBZNTRUPZukTLOKaC50UMq6zBcCSw3ixd3Ua9CaU9Fy0Bq8_jrbJz9W9TQKyIMn4XtHwG76Fhr878VqtDyvju7FOxqpLIVy852lWllZIlvrW5BYOQdUdCQk2ai03ctl6tWJiU9HNw4y-5XSY-33hLcQ4GA-tpwDDUGY0z1BsyPqBwPBZ0ERwBdbm7fP8Iu4ftELcy1OieAzsFhL-ncjhMUI84erFahhnSEA')

  useEffect(() => {
    if (hash) {
      const element = document.getElementById(hash.slice(1));
      // delete hash
      window.location.hash = "";
      if (element) {
        element.scrollIntoView({ behavior: "smooth" });
      }
    }
  }, [hash]);
  return (
    <>
      <Flex
        style={{ height: "91.5vh" }}
        direction={"column"}
        alignItems={"center"}
        justifyContent={"center"}
        height={"100vh"}
        className={styles.screen + " " + styles.screen_1}
      >
        <h1>КНИГА ПАМЯТИ</h1>
        <h2>Оренбургской области</h2>
        <Link
          to="#participants"
          style={{ width: "max-content" }}
          className={styles.link}
        >
          <Icon data={Magnifier} />
          <span>Найти человека</span>
        </Link>
      </Flex>
      <Flex
        className={styles.screen_about}
        id="about"
        justifyContent={'center'}
      >
        <Flex maxWidth={'800px'} alignItems={"center"}
        justifyContent={"center"}
        direction={'column'}
        spacing={{p: '7'}}>

        <Text className={styles.screen_about__title} variant="display-1">О проекте "Книга памяти"</Text>
        <Text className={styles.screen_about__text} variant="body-3">
          Проект "Книга памяти" призван стать важным ресурсом для сохранения
          памяти о героях, погибших в разных войнах, и увековечивания их
          подвигов. Наша цель — создать онлайн-платформу, на которой жители
          нашей области смогут найти информацию о тех, кто отдал свою жизнь за
          Родину. Каждый пользователь сможет исследовать биографии погибших,
          узнать о местах их захоронений.
        </Text>
        <Text className={styles.screen_about__text} variant="body-3">
          Проект объединяет современные технологии с исторической памятью,
          позволяя создать цифровую книгу, доступную для всех, кто хочет узнать
          больше о прошлом своей области и тех людях, которые стояли на защите
          нашей свободы.
        </Text>
        </Flex>
      </Flex>
      <Flex className={styles.screen_2} id="participants">
        <Flex className={styles.screen_2__left}>
          <Flex className={styles.findForm}>
            <p className={styles.screen_2__title}>Найти родственника</p>
            <TextInput value={state.lastName} onChange={
              (event) => {
                setState({...state, lastName: event.target.value})
              }
            } placeholder="Фамилия" />
            <TextInput value={state.name} onChange={
              (event) => {
                setState({...state, name: event.target.value})
              }
            } placeholder="Имя" />
            <TextInput value={state.middleName} onChange={
              (event) => {
                setState({...state, middleName: event.target.value})
              }
            } placeholder="Отчество" />
            <Button
              style={{
                backgroundColor: "#B3261E",
                color: "white",
                width: "max-content",
              }}
              onClick={() => {
                const query = new URLSearchParams(state).toString();
                navigate(`/search?${query}`)
              }}
            >
              Найти
            </Button>
          </Flex>
        </Flex>
        <Flex className={styles.landing2Picture}>
          <img src="/landing-2.png" />
        </Flex>
      </Flex>
      <Map />
      <Flex
        className={styles.screen_3 + " " + styles.screen_3_marginTop}
        id="conflicts"
      >
        <p className={styles.screen_3__title}>Военные конфликты</p>
        <Flex className={styles.screen_3__cards}>
          <Link to={"/search?conflict=Вторая Мировая Война&conflictId=WWII"} className={styles.bigButton}>
            ВОВ
          </Link>
          <Link to={"/search?conflict=Карабахский конфликт&conflictId=KARABAKH"} className={styles.bigButton}>
            КАРАБАХ
          </Link>
          <Link to={"/search?conflict=Специальная военная операция&conflictId=SVO"} className={styles.bigButton}>
            СВО
          </Link>
          <Link to={"/search?conflict=Гражданская война в Сирии&conflictId=SYRIA"} className={styles.bigButton}>
            СИРИЯ
          </Link>
          <Link to={"/search?conflict=Война во Вьетнаме&conflictId=VIETNAM"} className={styles.bigButton}>
            ВИЕТНАМ
          </Link>
          <Link to={"/search?conflict=Корейская война&conflictId=KOREA"} className={styles.bigButton}>
            КОРЕЯ
          </Link>
          <Link to={"/search?conflict=Афганская война&conflictId=AFGHAN"} className={styles.bigButton}>
            АФГАН
          </Link>
        </Flex>
      </Flex>
    </>
  );
}

export default Landing;

// {id: 'WWII', title: 'Вторая Мировая Война', dates: '1939-1945', records: Array(1)}
// {id: 'KARABAKH', title: 'Карабахский конфликт', dates: '1988-2023', records: Array(0)}
// {id: 'SVO', title: 'Специальная военная операция', dates: '2022-н.в.', records: Array(0)}
// {id: 'SYRIA', title: 'Гражданская война в Сирии', dates: '2011-н.в.', records: Array(0)}
// {id: 'VIETNAM', title: 'Война во Вьетнаме', dates: '1955-1975', records: Array(0)}
// {id: 'KOREA', title: 'Корейская война', dates: '1950-1953', records: Array(0)}
// {id: 'AFGHAN', title: 'Афганская война', dates: '1979-1989', records: Array(0)}