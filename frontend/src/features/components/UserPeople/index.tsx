import { Flex } from "@gravity-ui/uikit";
import PersonCard from "../PersonCard";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";

interface VeteranData {
  id: string;
  name: string;
  middle_name: string;
  last_name: string;
  birth_date: string;
  birth_place: string;
  military_rank: string;
  commissariat: string;
  awards: string[];
  death_date: string;
  burial_place: string;
  bio: string;
  map_id: string;
  documents: Document[];
  conflicts: string[];
  published: boolean;
}

interface Document {
  id: string;
  record_id: string;
  type: string;
  object_key: string;
}


const UserPeople = () => {
  const [people, setPeople] = useState<VeteranData[]>([]);

  async function fetchPeople() {
    const response = await fetch("/api/records/?published=true", {
      headers: {
        Authorization: `Bearer ${Cookies.get("access_token")}`,
      },
    });
    const data = await response.json();

    if (!response.ok) {
      // TODO notify about error
    } else {
      setPeople(data)
      console.log(data)
    }
  }

  useEffect(() => {
    fetchPeople()
  }, [])
  
  return (
    <Flex direction={"column"} gap={"2"} spacing={{pt: '2'}}>
      <Flex direction={'column'} gap={'2'}>
        {people.map((person, index) => (
          <PersonCard person={person} key={index}  editabe />
        ))}
        </Flex>
    </Flex>
  );
}

export default UserPeople