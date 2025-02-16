import {
  Avatar,
  Flex,
  Tab,
  TabList,
  TabPanel,
  TabProvider,
  Text,
} from "@gravity-ui/uikit";
import Cookies from "js-cookie";
import { useState } from "react";
import AdminUsers from "features/components/AdminUsers";
import AdminReq from "features/components/AdminReq";
import UserPeople from "features/components/UserPeople";

function User() {
  const token = Cookies.get("access_token");
  const [activeTab, setActiveTab] = useState("org-request");
  const data = JSON.parse(window.atob((token as string).split(".")[1]));
  
  const isAdmin = data.role === "admin" || data.role === "root";
  const isUser = data.role === "user";

  return (
    <Flex direction={"column"} spacing={{ p: "7" }}>
      
      <Flex gap={"2"} alignItems={"center"} justifyContent={"center"}>
        <Avatar size="xl" text={data.full_name} />
        <Flex direction={"column"}>
          <Text variant="display-1">{data.full_name}</Text>
          <Text variant="body-3">{data.organization}</Text>
        </Flex>
        <Flex>
          <a href="/api/logs/">Посмотреть логи</a>
        </Flex>
      </Flex>
      <TabProvider value={activeTab} onUpdate={setActiveTab}>
        <TabList>
          {isAdmin && <Tab value="org-request">Пользователи</Tab>}
          {isAdmin && <Tab value="requests">Заявки</Tab>}
          {isUser && <Tab value="added-people">Добавленные люди</Tab>}
        </TabList>
        {isAdmin && (
          <>
            <TabPanel value="org-request">
              <AdminUsers />
            </TabPanel>
            <TabPanel value="requests" style={{paddingTop: '1rem'}}>
              <AdminReq />
            </TabPanel>
          </>
        )}
        {isUser && (
          <TabPanel value="added-people">
            <UserPeople />
          </TabPanel>
        )}
      </TabProvider>
    </Flex>
  );
}

export default User;
