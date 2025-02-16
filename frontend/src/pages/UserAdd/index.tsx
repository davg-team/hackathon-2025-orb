import {
  Box,
  Button,
  Flex,
  TextArea,
  TextInput,
  Text,
  Select
} from "@gravity-ui/uikit";
import { DateField } from "@gravity-ui/date-components";
import styles from "./index.module.css";
import { useState } from "react";
import Cookies from "js-cookie";

interface Document {
  type: string;
  objectKey: string;
}

interface MemoryEntry {
  name: string;
  middleName?: string;
  lastName: string;
  birthDate: string;
  birthPlace: string;
  militaryRank: string;
  commissariat: string;
  awards: string[];
  deathDate: string;
  burialPlace: string;
  bio: string;
  documents: Document[];
  conflictId: string[];
}

interface IFile {
  file: File;
  type: string;
  objectKey: string;
  name: string;
  extension: string;
}

function UserAdd() {
  const [files, setFiles] = useState<IFile[]>([]);
  const [logo, setLogo] = useState<IFile>({
    file: new File([], ""),
    type: "",
    objectKey: "",
    name: "",
    extension: "",
  });
  const [state, setState] = useState<MemoryEntry>({
    name: "",
    middleName: "",
    lastName: "",
    birthDate: "",
    birthPlace: "",
    militaryRank: "",
    commissariat: "",
    awards: [],
    deathDate: "",
    burialPlace: "",
    bio: "",
    documents: [],
    conflictId: [],
  });

  // async function sendFiles() {
  //   let filesToSend = [logo, ...files];
  //   filesToSend = files.filter((file) => file.file.size > 0);
  //   for (const file of filesToSend) {
  //     // octet
  //     const response = await fetch("/user-documents", {
  //       method: "PUT",
  //       headers: { "Content-Type": "application/octet-stream" },
  //       body: await file.file.arrayBuffer(),
  //     });

  //     if (!response.ok) {
  //       // TODO notify about error
  //     } else {
  //       const data = await response.json();
  //       console.log(data);
  //     }
    
  //   }
  // }

  async function registerPeople() {
    // await sendFiles()

    const response = await fetch("/api/records/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${Cookies.get("access_token")}`,
      },
      body: JSON.stringify({
        name: state.name,
        middle_name: state.middleName,
        last_name: state.lastName,
        birth_date: state.birthDate,
        birth_place: state.birthPlace,
        military_rank: state.militaryRank,
        commissariat: state.commissariat,
        awards: [...state.awards],
        death_date: state.deathDate,
        burial_place: state.burialPlace,
        bio: state.bio,
        documents: [],
        conflict_id: [...state.conflictId],
      })
    });
    const result = await response.json();
    console.log(result);
  }

  function handleFilesUpload(event: React.ChangeEvent<HTMLInputElement>) {
    const uploadedFiles = event.target.files;
    if (uploadedFiles) {
      const newFiles = Array.from(uploadedFiles).map((file) => ({
        file,
        type: file.type,
        objectKey: URL.createObjectURL(file),
        name: file.name,
        extension: file.name.split(".").pop() || "",
      }));
      setFiles((prevFiles) => {
        const uniqueFiles = [...prevFiles, ...newFiles].reduce((acc, file) => {
          if (!acc.some((t) => t.objectKey === file.objectKey)) {
            acc.push(file);
          }
          return acc;
        }, [] as IFile[]);
        return uniqueFiles;
      });
      console.log("Uploaded files:", newFiles);
    }
  }

  function handleFileUpload(event: React.ChangeEvent<HTMLInputElement>) {
    const uploadedFile = event.target.files?.[0];
    if (uploadedFile) {
      const newFile = {
        file: uploadedFile,
        type: uploadedFile.type,
        objectKey: URL.createObjectURL(uploadedFile),
        name: uploadedFile.name,
        extension: uploadedFile.name.split(".").pop() || "",
      };
      setLogo(newFile);
      console.log("Uploaded file:", newFile);
    }
  }

  function removeFile(index: number) {
    setFiles((prevFiles) => prevFiles.filter((_, i) => i !== index));
  }

  return (
    <Flex justifyContent={"center"} spacing={{ p: "7" }}>
      <Flex className={styles.addForm}>
        <Flex
          wrap
          className={styles.addForm__left}
          direction={"column"}
          spacing={{ mb: "3" }}
        >
          <Box className={styles.photo__wrapper}>
            <img width={'200'} height={'250'} src={logo.objectKey} alt=""/>
            <Text variant="body-3">Выбрать фотографию</Text>
            <input type="file" onChange={handleFileUpload} />
          </Box>
          <br />
          <Box className={styles.docs}>
            <Box className={styles.doc}>
              <Text variant="body-3">Выбрать файлы</Text>
              <input type="file" className={styles.hidden} onChange={handleFilesUpload} multiple />
            </Box>
            {files.map((file, index) => (
              <Box key={index} className={styles.doc}>
                <span>{file.name}</span>
                <Button view="flat" onClick={() => removeFile(index)}>
                  Удалить
                </Button>
              </Box>
            ))}
          </Box>
        </Flex>
        <Flex
          wrap={"wrap"}
          gap={"3"}
          alignItems={"center"}
          justifyContent={"center"}
        >
          <TextInput
            onChange={(e) => setState({ ...state, name: e.target.value })}
            value={state.name}
            style={{ width: "100%" }}
            placeholder="Фамилия"
          />
          <TextInput
            onChange={(e) => setState({ ...state, middleName: e.target.value })}
            value={state.middleName}
            style={{ width: "100%" }}
            placeholder="Имя"
          />
          <TextInput
            onChange={(e) => setState({ ...state, lastName: e.target.value })}
            value={state.lastName}
            style={{ width: "100%" }}
            placeholder="Отчество"
          />
          <DateField
            onUpdate={(date) => {
              setState({ ...state, birthDate: (date?.format("YYYY-MM-DD") as string) });
            }}
            style={{ width: "100%" }}
            label="Дата рождения"
          />
          <TextInput
            style={{ width: "100%" }}
            placeholder="Место рождения"
            onChange={(e) => setState({ ...state, birthPlace: e.target.value })}
          />
          <TextInput
            style={{ width: "100%" }}
            placeholder="Военной звание"
            onChange={(e) =>
              setState({ ...state, militaryRank: e.target.value })
            }
          />
          <TextInput
            style={{ width: "100%" }}
            placeholder="Военный комиссариат"
            onChange={(e) =>
              setState({ ...state, commissariat: e.target.value })
            }
          />
          <Select
            multiple
            onUpdate={(e) => setState({ ...state, conflictId: e})}
            options={[
              { value: "WWII", content: "Вторая Мировая Война" },
              { value: "KARABAKH", content: "Карабахский конфликт" },
              { value: "SVO", content: "Специальная военная операция" },
              { value: "SYRIA", content: "Гражданская война в Сирии" },
              { value: "VIETNAM", content: "Война во Вьетнаме" },
              { value: "KOREA", content: "Корейская война" },
              { value: "AFGHAN", content: "Афганская война" }
            ]}
            placeholder="Военные конфликты"
          />

          <TextInput
            onChange={(e) => setState({ ...state, awards: e.target.value.split(",") })}
            placeholder={"Награды"}
            style={{ width: "100%" }}
          />
          <DateField
            onUpdate={(date) => {
              setState({ ...state, deathDate: (date?.format("YYYY-MM-DD") as string) });
            }}
            style={{ width: "100%" }}
            label="Дата смерти"
          />
          <TextInput
            onChange={(e) =>
              setState({ ...state, burialPlace: e.target.value })
            }
            value={state.burialPlace}
            style={{ width: "100%" }}
            placeholder="Место захоронения"
          />
          <TextArea
            onChange={(e) => setState({ ...state, bio: e.target.value })}
            value={state.bio}
            style={{ width: "100%" }}
            placeholder="Описание иных фактов биографии"
          />
          <Button
            view="action"
            onClick={registerPeople}
            style={{ color: "white", justifySelf: "center" }}
          >
            сохранить
          </Button>
        </Flex>
      </Flex>
    </Flex>
  );
}

export default UserAdd;
