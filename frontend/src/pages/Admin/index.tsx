import { Avatar, Flex, Text } from "@gravity-ui/uikit";
import styles from "./index.module.css";
import PersonCard from "features/components/PersonCard";

function Admin() {
  return (
    <Flex direction={"column"} spacing={{ p: "7" }}>
      <Flex gap={"2"} alignItems={"center"} justifyContent={"center"}>
        <Avatar size="xl" text="as" />
        <Flex direction={"column"}>
          <Text variant="display-1">Иван Иванов</Text>
          <Text variant="body-3">Администратор</Text>
        </Flex>
      </Flex>
      <Flex direction={"column"} gap={"2"}>
        <Text variant="display-1" className={styles.addedPeople}>
          Заявки
        </Text>
        <PersonCard accessable />
        <PersonCard accessable />
        <PersonCard accessable />
        <PersonCard accessable />
      </Flex>
    </Flex>
  );
}

export default Admin;


//<TabProvider value={activeTab} onUpdate={setActiveTab}>
//            <TabList>
//              <Tab value="monday">Понедельник</Tab>
//              <Tab value="weekday">Будни</Tab>
//              <Tab value="saturday">Суббота</Tab>
//            </TabList>
//            <TabPanel value="monday">
//              <Flex direction={"column"} gap={"2"}>
//                <DateField
//                  style={{ width: "max-content" }}
//                  label="Время начала записи:"
//                  placeholder=""
//                  format="HH:mm:ss"
//                />
//                <Flex direction={"column"} gap={"2"}>
//                  <Flex gap="2">
//                    <DateField
//                      style={{ width: "max-content" }}
//                      label="Начало:"
//                      placeholder=""
//                      format="HH:mm:ss"
//                    />
//                    <DateField
//                      style={{ width: "max-content" }}
//                      label="Конец:"
//                      placeholder=""
//                      format="HH:mm:ss"
//                    />
//                    <Select options={typeSelectOptions} label={"Объявление:"} />
//                    <Flex alignItems={"center"} gap={"2"}>
//                      Фиксировать время:
//                      <SegmentedRadioGroup>
//                        <SegmentedRadioGroupOption value="start">
//                          Начало
//                        </SegmentedRadioGroupOption>
//                        <SegmentedRadioGroupOption value="nd">
//                          Конец
//                        </SegmentedRadioGroupOption>
//                      </SegmentedRadioGroup>
//                    </Flex>
//                    <Flex gap={'2'} alignItems={'center'}>
//                      <Checkbox>Fade In</Checkbox>
//                      <Checkbox>Fade Out</Checkbox>
//                    </Flex>
//										<Button>
//										<Icon data={GearPlay} />
//										</Button>
//                  </Flex>
//                </Flex>
//                {/* popover */}
//                <Button style={{ width: "max-content" }} view="outlined">
//                  <Icon data={PencilToLine} />
//                  Добавить объявление
//                  <Icon data={ChevronDown} />
//                </Button>
//                <DateField
//                  style={{ width: "max-content" }}
//                  label="Время конца записи:"
//                  placeholder=""
//                  format="HH:mm:ss"
//                />
//								<Flex direction={'column'} style={{width: 'max-content'}} gap={'2'}>
//									<Text variant="body-3">
//										3. Параметры
//									</Text>
//									<Select label="Формат выходного файла" options={typeExtSelectOptions} />
//									<Select label="Битрейт" options={typeBitraitSelectOptions} />
//									<Button style={{width: "max-content"}} view="action">
//									<Icon data={PencilToLine} />
//									Сохранить
//									</Button>
//								</Flex>
//              </Flex>
//            </TabPanel>
//            <TabPanel value="weekday">Будни</TabPanel>
//            <TabPanel value="saturday">Суббота</TabPanel>
//          </TabProvider>