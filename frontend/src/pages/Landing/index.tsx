import { Button, Flex, Icon, TextInput } from "@gravity-ui/uikit";
import { Magnifier } from "@gravity-ui/icons";
import styles from "./index.module.css";
import Map from 'features/components/Map'

function Landing() {
  return (
    <>
      <Flex
        style={{ height: "90vh" }}
        direction={"column"}
        alignItems={"center"}
        justifyContent={"center"}
        height={"100vh"}
        className={styles.screen + " " + styles.screen_1}
      >
        <h1>КНИГА ПАМЯТИ</h1>
        <h2>Оренбургской области</h2>
        <a style={{ width: "max-content" }} className={styles.link}>
          <Icon data={Magnifier} />
          <span>Найти человека</span>
        </a>
      </Flex>
      <Flex
        justifyContent={"space-between"}
        spacing={{ pl: "10", pr: "4" }}
        style={{ height: "100vh" }}
        className={styles.screen + " " + styles.screen_2}
      >
        <Flex alignItems={'center'} justifyContent={'center'}>
          <Flex
            justifyContent={"center"}
						alignItems={'center'}
            className={styles.findForm + " " + styles.findFormMargin}
            direction={"column"}
            gap={"4"}
          >
            <p className={styles.screen_2__title}>Найти родственника</p>
            <TextInput placeholder="Фамилия" />
            <TextInput placeholder="Имя" />
            <TextInput placeholder="Отчество" />
            <Button
              style={{
                backgroundColor: "#B3261E",
                color: "white",
                width: "max-content",
              }}
            >
              <Icon data={Magnifier} />
              Найти
            </Button>
          </Flex>
        </Flex>
        <Flex className={styles.landing2Picture}>
          <img src="/landing-2.png" />
        </Flex>
      </Flex>
			<Map />
			<Flex alignItems={'center'} direction={'column'} className={styles.screen_3_marginTop} height={'100vh'}>
				<p className={styles.screen_3__title}>
					Военные конфликты 
				</p>
				<Flex justifyContent={'center'} wrap='wrap' gap={'6'}>
					<a className={styles.bigButton}>ВОВ</a>
					<a className={styles.bigButton}>Чеченская</a>
					<a className={styles.bigButton}>Афганистан</a>
					<a className={styles.bigButton}>СВО</a>
				</Flex>
			</Flex>
    </>
  );
}

export default Landing;
