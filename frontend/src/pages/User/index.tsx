// import {
//   Avatar,
//   Flex,
//   Tab,
//   TabList,
//   TabPanel,
//   TabProvider,
//   Text,
// } from "@gravity-ui/uikit";
// import styles from "./index.module.css";
// import Cookies from "js-cookie";
// import PersonCard from "features/components/PersonCard";
// import { useState } from "react";

// function User() {
//   const token = Cookies.get("access_token");
//   const [activeTab, setActiveTab] = useState("");
//   const data = JSON.parse(atob((token as string).split(".")[1]));
//   console.log(data);
//   return (
//     <Flex direction={"column"} spacing={{ p: "7" }}>
//       <Flex gap={"2"} alignItems={"center"} justifyContent={"center"}>
//         <Avatar size="xl" text={data.full_name} />
//         <Flex direction={"column"}>
//           <Text variant="display-1">{data.full_name}</Text>
//           <Text variant="body-3">{data.organization}</Text>
//         </Flex>
//       </Flex>
//       <TabProvider value={activeTab} onUpdate={setActiveTab}>
//         <TabList>
//           <Tab value="org-request">Пользователи</Tab>
//           <Tab value="requests">Заявки</Tab>
//           <Tab value="added-people">Добавленные люди</Tab>
//         </TabList>
//         <TabPanel value="org-request">123 w</TabPanel>
//         <TabPanel value="requests">Будни</TabPanel>
//         <TabPanel value="added-people">
//           <Flex direction={"column"} gap={"2"}>
//             <Text variant="display-1" className={styles.addedPeople}>
//               Добавленные люди
//             </Text>
//             <PersonCard editabe={true} />
//             <PersonCard editabe={true} />
//             <PersonCard editabe={true} />
//             <PersonCard editabe={true} />
//           </Flex>
//         </TabPanel>
//       </TabProvider>
//     </Flex>
//   );
// }

// export default User;

import {
  Avatar,
  Flex,
  Tab,
  TabList,
  TabPanel,
  TabProvider,
  Text,
} from "@gravity-ui/uikit";
import styles from "./index.module.css";
import Cookies from "js-cookie";
import PersonCard from "features/components/PersonCard";
import { useState } from "react";

function User() {
  const token = Cookies.get("access_token");
  const [activeTab, setActiveTab] = useState("org-request");
  const data = JSON.parse(atob((token as string).split(".")[1]));
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
      </Flex>
      <TabProvider value={activeTab} onUpdate={setActiveTab}>
        <TabList>
          {isAdmin && <Tab value="org-request">Пользователи</Tab>}
          {isAdmin && <Tab value="requests">Заявки</Tab>}
          {isUser && <Tab value="added-people">Добавленные люди</Tab>}
        </TabList>
        {isAdmin && (
          <>
            <TabPanel value="org-request"></TabPanel>
            <TabPanel value="requests"></TabPanel>
          </>
        )}
        {isUser && (
          <TabPanel value="added-people">
            <Flex direction={"column"} gap={"2"}>
              <Text variant="display-1" className={styles.addedPeople}>
                Добавленные люди
              </Text>
              <PersonCard editabe={true} />
              <PersonCard editabe={true} />
              <PersonCard editabe={true} />
              <PersonCard editabe={true} />
            </Flex>
          </TabPanel>
        )}
      </TabProvider>
    </Flex>
  );
}

export default User;
