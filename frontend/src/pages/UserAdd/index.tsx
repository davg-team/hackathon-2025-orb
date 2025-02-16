import { Box, Button, Flex, Select, TextArea, TextInput } from "@gravity-ui/uikit";
import {DateField} from '@gravity-ui/date-components';
import styles from "./index.module.css";

function UserAdd() {
  return <Flex justifyContent={"center"} spacing={{p: '7'}}>
    <Flex className={styles.addForm}>
      <Flex wrap className={styles.addForm__left} direction={"column"} spacing={{mb: '3'}}>
        <Box className={styles.photo__wrapper}>
          <img src="" alt="" className={styles.hidden} />
          <input type="file" />
        </Box>
        <Box className={styles.docs}>
          
          <Box className={styles.doc}>
            <input type="file" className={styles.hidden} />
          </Box>
        </Box>
      </Flex>
      <Flex wrap={'wrap'} gap={'3'} alignItems={'center'} justifyContent={'center'}>
        <TextInput style={{width: '100%'}} placeholder="Фамилия" />
        <TextInput style={{width: '100%'}} placeholder="Имя" />
        <TextInput style={{width: '100%'}} placeholder="Отчество" />
        <DateField style={{width: '100%'}} label="Дата рождения" />
        <TextInput style={{width: '100%'}} placeholder="Военный комиссариат" />
        <TextInput style={{width: '100%'}} placeholder="Военной звание" />
        <Select width={'max'} placeholder={'Военные конфликты'} />
        <Select width={'max'} placeholder={'Награды'} />
        <DateField style={{width: '100%'}} label="Дата смерти" />
        <TextInput style={{width: '100%'}} placeholder="Место захоронения" />
        <TextArea style={{width: '100%'}} placeholder="Описание иных фактов биографии" />
        <Button view='action' style={{color: 'white', justifySelf: 'center'}}>сохранить</Button>
      </Flex>
    </Flex>
  </Flex>;
}

export default UserAdd;
