import { Flex } from "@gravity-ui/uikit";
import PersonCard from "../PersonCard";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";

const AdminReq = () => {
  const [requests, setRequests] = useState([]);

  async function fetchRequests() {
    const response = await fetch("/api/records/?published=false", {
      headers: {
        Authorization: `Bearer ${Cookies.get("access_token")}`,
      },
    })
    const data = await response.json();

    if (!response.ok) {
      // TODO notify about error
    } else {
      setRequests(data)
      console.log(data)
    }
  }

  

  useEffect(() => {
    fetchRequests()
  }, [])

  return (
    <Flex direction={"column"} gap={"2"}>
      {requests.length > 0 ? requests.map((request, index) => (
      <PersonCard fetchRequests={fetchRequests} person={request} key={index} accessable />  
      )) : <p>Нет новых заявок</p>}
    </Flex>
  );
};

export default AdminReq;
