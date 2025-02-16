/* eslint-disable @typescript-eslint/ban-ts-comment */
import { Flex, TextInput, Text, Loader } from "@gravity-ui/uikit";
import styles from "./index.module.css";
import PersonCard from "features/components/PersonCard";
import { useEffect, useState } from "react";
import { useSearchParams } from "react-router-dom";
import { debounce } from "lodash";

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

function Search() {
  const [data, setData] = useState<VeteranData[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');
  const [params] = useSearchParams();
  const [firstLoading, setFirstLoading] = useState(false);
  const [state, setState] = useState({
    name: "",
    middleName: "",
    lastName: "",
  });

  const fetchConflict = debounce(async () => {
    try {
      const response = await fetch(
        `/api/conflicts/${params.get("conflictId")}`
      );
      const result = await response.json();
      setData(result.records);
    } catch (err) {
      console.log(err);
      setError("Произошла ошибка");
    } finally {
      setLoading(false);
    }
  }, 500);

  const fetchData = debounce(async () => {
    setLoading(true);
    try {
      const query = new URLSearchParams(state).toString();
      const response = await fetch(`/api/records/?${query}`);
      const result = await response.json();
      setData(result);
    } catch (err) {
      console.log(err);
      setError("Произошла ошибка");
    } finally {
      setLoading(false);
    }
  }, 500);

  useEffect(() => {
    if (firstLoading === false) {
      setState({
        name: params.get("name") || "",
        middleName: params.get("middleName") || "",
        lastName: params.get("lastName") || "",
      });
      setFirstLoading(true);
    }

    if (params.has("conflict")) {
      fetchConflict();
    } else {
      fetchData();
    }
  }, [state]);

  return (
    <Flex spacing={{ p: "5" }} direction={"column"} alignItems={"center"}>
      <Flex alignItems={"flex-start"} width={"100%"}>
        {params.has("conflict") && (
          <Text
            style={{ color: "#B3261E", fontWeight: "bold" }}
            variant="display-1"
          >
            {params.get("conflict")}
          </Text>
        )}
      </Flex>
      <Flex wrap={"wrap"} gap={"3"} direction={"row"} width={"100%"}>
        {!params.has("conflict") && (
          <>
            <TextInput
              placeholder="Фамилия"
              value={state.lastName}
              onChange={(e) => {
                setState({ ...state, lastName: e.target.value });
              }}
            />
            <TextInput
              placeholder="Имя"
              value={state.name}
              onChange={(e) => setState({ ...state, name: e.target.value })}
            />
            <TextInput
              placeholder="Отчество"
              value={state.middleName}
              onChange={(e) =>
                setState({ ...state, middleName: e.target.value })
              }
            />
          </>
        )}
      </Flex>
      <br />
      {/* @ts-ignore */}
      {error && <p style={{ color: "red" }}>Ошибка: {error.message}</p>}
      <div className={styles.personCards}>
        {loading ? (
          <Loader />
        ) : error ? (
          // @ts-ignore
          <p style={{ color: "red" }}>Ошибка: {error.message}</p>
        ) : data.length > 0 ? (
          data.map((person) => (<PersonCard person={person} key={person.id} />))
        ) : (
          <p>Ничего не найдено</p>
        )}
      </div>
    </Flex>
  );
}

export default Search;
