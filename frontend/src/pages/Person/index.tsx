/* eslint-disable @typescript-eslint/ban-ts-comment */
import {
  Box,
  Flex,
  Icon,
  Text,
  TextArea,
  Skeleton,
  Popover,
} from "@gravity-ui/uikit";
import { ArrowDownToLine } from "@gravity-ui/icons";
import styles from "./index.module.css";
import { useParams } from "react-router-dom";
import { useCallback, useEffect, useState } from "react";

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
  conflicts: {
    dates: string;
    id: string;
    records: VeteranData[];
    title: string;
  }[];
  published: boolean;
}

interface Document {
  id: string;
  record_id: string;
  type: string;
  object_key: string;
}

function Person() {
  const [data, setData] = useState<VeteranData | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState();
  const params = useParams();

  const fetchData = useCallback(async () => {
    setLoading(true);
    console.log(params);
    try {
      const response = await fetch(`/api/records/${params.id}`);
      const result = await response.json();
      setData(result);
      console.log(result);
    } catch (err: unknown) {
      // @ts-ignore
      setError(err);
    } finally {
      setLoading(false);
    }
  }, [params.id]);

  useEffect(() => {
    fetchData();
  }, [fetchData, params.id]);

  if (loading) {
    return (
      <Flex className={styles.person}>
        <Flex className={styles.person__card_wrapper}>
          <Flex className={styles.person__card}>
            <Flex direction={"column"} gap={"2"}>
              <Skeleton className={styles.photo__wrapper} />
              <Skeleton className={styles.medals__item} />
              <Skeleton className={styles.medals__item} />
              <Skeleton className={styles.medals__item} />
            </Flex>
            <Flex direction={"column"}>
              <Skeleton />
              <Skeleton />
              <Skeleton />
              <Skeleton />
              <Skeleton />
              <Skeleton />
              <Skeleton />
              <Skeleton />
            </Flex>
          </Flex>
        </Flex>
        <Flex className={styles.person__photos}>
          <Text variant="display-1">Фотографии</Text>
          <Flex className={styles.photos}>
            <Skeleton className={styles.photo} />
            <Skeleton className={styles.photo} />
            <Skeleton className={styles.photo} />
          </Flex>
        </Flex>
        <Flex className={styles.person__docs + " " + styles.docs}>
          <Text variant="display-1">Документы</Text>
          <Flex className={styles.docs__wrapper}>
            <Skeleton className={styles.docs__item} />
            <Skeleton className={styles.docs__item} />
          </Flex>
        </Flex>
      </Flex>
    );
  }

  if (error) {
    return <Text>Произошла ошибка при загрузке данных</Text>;
  }

  return (
    <Flex className={styles.person}>
      <Flex className={styles.person__card_wrapper}>
        <Flex className={styles.person__card}>
          <Flex direction={"column"} gap={"2"}>
            <div className={styles.photo__wrapper}>
              <img
                width={200}
                height={200}
                style={{ objectFit: "contain" }}
                // eslint-disable-next-line no-constant-binary-expression
                src={`/api/files/${data?.documents.filter((doc) => doc.type === "logo")[0].object_key}` || "/person-default.jpg"}
                alt="фотография"
                className={styles.person__photo}
              />
            </div>
            <div className={styles.medals}>
              {data?.awards.map((medal, index) => (
                <Popover content={medal as string} key={index}>
                  <img
                    src="/Medal.svg"
                    alt="medal"
                    className={styles.medals__item}
                  />
                </Popover>
              ))}
            </div>
          </Flex>
          <Flex direction={"column"}>
            <Text variant="body-3">
              <b>Фамилия:</b> {data?.last_name || "–"}
            </Text>
            <Text variant="body-3">
              <b>Имя:</b> {data?.name}
            </Text>
            <Text variant="body-3">
              <b>Отчество:</b> {data?.middle_name || "–"}
            </Text>
            <Text variant="body-3">
              <b>Дата рождения:</b> {data?.birth_date || "–"}
            </Text>
            <Text variant="body-3">
              <b>Военный комиссариат:</b> {data?.commissariat || "–"}
            </Text>
            <Text variant="body-3">
              <b>Дата смерти:</b> {data?.death_date || "–"}
            </Text>
            <Text variant="body-3">
              <b>Военные конфликты:</b>{" "}
              {data?.conflicts.map((conflict) => conflict.title).join(", ") ||
                "–"}
            </Text>
            <Text variant="body-3">
              <b>Место захоронения:</b> {data?.burial_place || "–"}
            </Text>
            <Text variant="body-3">
              <b>Описание</b>
            </Text>
            <TextArea
              placeholder="Content"
              defaultValue={data?.bio || "Content about person"}
              style={{ width: "100%" }}
              rows={5}
              disabled
            ></TextArea>
          </Flex>
        </Flex>
      </Flex>
      <Flex className={styles.person__photos}>
        <Text style={{ fontWeight: "700" }} variant="display-1">
          Фотографии
        </Text>
        <Flex className={styles.photos}>
          {(data?.documents.length as number) > 0 ? (
            data?.documents.filter((doc) => doc.type === "image").map((doc) => (
              <Box
                key={doc.id}
                width={250}
                height={200}
                className={styles.photo}
              >
                <img src={`/api/files/${doc.object_key || ''}`} alt="" />
              </Box>
            ))
          ) : (
            <Text variant="body-3">Фотографии отсутствуют</Text>
          )}
        </Flex>
      </Flex>
      <Flex className={styles.person__docs + " " + styles.docs}>
        <Text style={{ fontWeight: "700" }} variant="display-1">
          Документы
        </Text>
        <Flex className={styles.docs__wrapper}>
          {data && data?.documents.length > 0 ? (
            data?.documents.map((doc) => (
              <a
                download
                key={doc.id}
                href={'/api/files/' + doc.object_key}
                className={styles.docs__item}
              >
                Документ
                <Icon data={ArrowDownToLine} />
              </a>
            ))
          ) : (
            <Text variant="body-3">Документы отсутствуют</Text>
          )}
        </Flex>
      </Flex>
    </Flex>
  );
}

export default Person;
