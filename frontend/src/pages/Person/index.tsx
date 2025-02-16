/* eslint-disable @typescript-eslint/ban-ts-comment */
import {
  Box,
  Flex,
  Icon,
  Text,
  TextArea,
  Skeleton,
  Popover,
  Button,
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

interface PersonProps {
  mode: string;
}

function Person({ mode }: PersonProps) {
  const [data, setData] = useState<VeteranData | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState();
  const [hasLogo, setHasLogo] = useState(false);
  const [images, setImages] = useState<Document[]>([]);
  const params = useParams();

  const fetchData = useCallback(async () => {
    setLoading(true);
    try {
      const response = await fetch(`/api/records/${params.id}`);
      const result = await response.json();
      setData(result);

      setHasLogo(result.documents.some((doc: Document) => doc.type === "logo"));

      setImages(
        result.documents
          .filter((doc: Document) => doc.type === "image")
          .map((doc: Document) => doc.object_key)
      );
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

  // После полной загрузки данных, если mode === "print", вызываем печать.
  useEffect(() => {
    if (mode === "print" && !loading && data) {
      // Небольшая задержка для гарантии, что всё отрендерилось
      setTimeout(() => {
        window.print();
      }, 3);
    }
  }, [mode, loading, data]);

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
                src={
                  hasLogo ? `/api/files/${data?.documents.filter((docs: Document) => docs.type === "logo")[0].object_key}` :
                  "/person-default.jpg"
                }
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
      
      {/* Кнопка вывода на печать при условии, что mode !== "print" */}
      {mode !== "print" && (
        <>
        <Flex className={styles.person__photos}>
        <Text style={{ fontWeight: "700" }} variant="display-1">
          Фотографии
        </Text>
        <Flex className={styles.photos}>
          {(data?.documents.length as number) > 0 ? (
            images.map((doc: Document) => (
              <Box
                key={doc.id}
                width={250}
                height={200}
                className={styles.photo}
              >
                <img src={`/api/files/${doc.object_key || ""}`} alt="" />
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
                href={"/api/files/" + doc.object_key}
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
      <br />
        <Flex>
          <Button
            onClick={() => window.location.assign(`/person-print/${data?.id}`)}
          >
            Вывести на печать
          </Button>
        </Flex>
        </>
      )}
    </Flex>
  );
}

export default Person;
