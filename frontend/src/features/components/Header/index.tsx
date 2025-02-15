import {Avatar, Button, Flex} from "@gravity-ui/uikit";
import {GearPlay} from "@gravity-ui/icons";
import styles from './index.module.css'

function Header() {
  return (
    <Flex className={styles.header} alignItems={'center'} justifyContent={'space-between'} spacing={{p: '5'}}>
      <Flex className="header__left" gap={'3'} alignItems={'center'}>
        <Avatar icon={GearPlay} size='xl' />
        <a className={styles.link}>о проекте</a>
        <a className={styles.link}>участники</a>
        <a className={styles.link}>военные конфликты</a>
        <a className={styles.link}>+ добавить участника</a>
      </Flex>
      <Flex className="header__right">
        <Avatar icon={GearPlay} />
      </Flex>
    </Flex>
  );
}

export default Header;
