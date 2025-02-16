/* eslint-disable @typescript-eslint/ban-ts-comment */
import {
  Flex,
  Avatar,
  Text,
  Modal,
  TextInput,
  Button,
  Select,
} from "@gravity-ui/uikit";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";

interface IUser {
  full_name: string;
  organization: string;
  position: string;
  phone?: string;
  email: string;
  password: string;
  snils: string;
  role: "root" | "admin" | "user";
  id: number;
}

const AdminUsers = () => {
  const [modalOpen, setModalOpen] = useState(false);
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [loading, setLoading] = useState(false);
  const [changeForm, setChangeForm] = useState<IUser>({
    full_name: "",
    organization: "",
    position: "",
    password: "",
    phone: "",
    email: "",
    snils: "",
    role: "user",
    id: 0,
  });
  const [users, setUsers] = useState<IUser[]>([]);

  async function createUser() {
    const token = Cookies.get("access_token");
    const response = await fetch("/api/users/users/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        full_name: changeForm.full_name,
        organization: changeForm.organization,
        position: changeForm.position,
        password: changeForm.password,
        phone: changeForm.phone,
        email: changeForm.email,
        snils: changeForm.snils,
        role: changeForm.role,
      }),
    });

    if (!response.ok) {
      // TODO notify about error
    } else {
      await fetchUsers();
      afterChangeHandler();
      setCreateModalOpen(false);
    }

    
  }

  function afterChangeHandler() {
    setChangeForm({
      full_name: "",
      organization: "",
      position: "",
      password: "",
      phone: "",
      email: "",
      snils: "",
      role: "user",
      id: 0,
    });
  }

  async function deleteUser() {
    setLoading(true);
    await fetch(`/api/users/users/${changeForm.id}`, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${Cookies.get("access_token")}`,
      },
    });
    await fetchUsers();
    onCloseHandler();
  }

  async function fetchUsers() {
    const token = Cookies.get("access_token");
    const response = await fetch("/api/users/users/", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    const data = await response.json();

    if (!response.ok) {
      // TODO notify about error
    } else {
      setUsers(data);
    }
  }

  async function updateUser() {
    setLoading(true);
    await fetch(`/api/users/users/${changeForm.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${Cookies.get("access_token")}`,
      },
      body: JSON.stringify({
        full_name: changeForm.full_name,
        organization: changeForm.organization,
        position: changeForm.position,
        password: changeForm.password,
        phone: changeForm.phone,
        email: changeForm.email,
        snils: changeForm.snils,
        role: changeForm.role,
      }),
    });
    await fetchUsers();
    onCloseHandler();
  }

  useEffect(() => {
    fetchUsers();
  }, []);

  function onCloseHandler() {
    setLoading(false);
    afterChangeHandler();
    setModalOpen(false);
  }

  return (
    <Flex direction={"column"} gap={"2"} spacing={{ pb: "2", pt: "2" }}>
      <Button
        onClick={() => {
          setCreateModalOpen(true);
          setChangeForm({
            full_name: "",
            organization: "",
            position: "",
            password: "",
            phone: "",
            email: "",
            snils: "",
            role: "user",
            id: 0,
          })
        }}
      >
        Создать пользователя
      </Button>
      {users.map((user: IUser, index: number) => (
        <Flex
          onClick={() => {
            setChangeForm({ ...user, password: "" });
            setModalOpen(true);
          }}
          key={index}
          style={{
            backgroundColor: "#EFEFEF",
            padding: ".4rem",
            borderRadius: ".5rem",
            width: "100%",
          }}
          alignItems={"center"}
          gap={"2"}
        >
          <Avatar text={user.full_name[0]} />
          <Text variant="body-1">{user.full_name}</Text>
        </Flex>
      ))}

      <Modal
        open={modalOpen}
        onClose={() => {
          onCloseHandler();
        }}
      >
        <Flex spacing={{ p: "4" }} gap={"3"} direction={"column"}>
          <Text variant="display-1"> Редактирование пользователя</Text>
          <TextInput
            label="Полное имя"
            value={changeForm.full_name}
            onChange={(e) =>
              setChangeForm({ ...changeForm, full_name: e.target.value })
            }
          />
          <TextInput
            label="Почта"
            value={changeForm.email}
            onChange={(e) =>
              setChangeForm({ ...changeForm, email: e.target.value })
            }
          />
          <TextInput
            label="Пароль"
            value={changeForm.password}
            onChange={(e) =>
              setChangeForm({ ...changeForm, password: e.target.value })
            }
          />
          <TextInput
            label="Телефон"
            value={changeForm.phone}
            onChange={(e) =>
              setChangeForm({ ...changeForm, phone: e.target.value })
            }
          />
          <TextInput
            label="СНИЛС"
            value={changeForm.snils}
            onChange={(e) =>
              setChangeForm({ ...changeForm, snils: e.target.value })
            }
          />
          <TextInput
            label="Организация"
            value={changeForm.organization}
            onChange={(e) =>
              setChangeForm({ ...changeForm, organization: e.target.value })
            }
          />
          <TextInput
            label="Должность"
            value={changeForm.position}
            onChange={(e) =>
              setChangeForm({ ...changeForm, position: e.target.value })
            }
          />
          <Select 
            options={[{value: 'root', content: 'root'}, {value: 'admin', content: 'admin'}, {value: 'user', content: 'user'}]}
            // @ts-ignore
            onChange={(e) =>
              setChangeForm({ ...changeForm, role: e.target.value })
            }
            label="Роль"
          />
          <Button loading={loading} onClick={updateUser}>Сохранить</Button>
          <Button loading={loading} onClick={deleteUser}>Удалить</Button>
        </Flex>
      </Modal>
      <Modal
        open={createModalOpen}
        onClose={() => {
          onCloseHandler();
        }}
      >
        <Flex spacing={{ p: "4" }} gap={"3"} direction={"column"}>
          <Text variant="display-1">Создать пользователя</Text>
          <TextInput
            label="Полное имя"
            value={changeForm.full_name}
            onChange={(e) =>
              setChangeForm({ ...changeForm, full_name: e.target.value })
            }
          />
          <TextInput
            label="Почта"
            value={changeForm.email}
            onChange={(e) =>
              setChangeForm({ ...changeForm, email: e.target.value })
            }
          />
          <TextInput
            label="Пароль"
            value={changeForm.password}
            onChange={(e) =>
              setChangeForm({ ...changeForm, password: e.target.value })
            }
          />
          <TextInput
            label="Телефон"
            value={changeForm.phone}
            onChange={(e) =>
              setChangeForm({ ...changeForm, phone: e.target.value })
            }
          />
          <TextInput
            label="СНИЛС"
            value={changeForm.snils}
            onChange={(e) =>
              setChangeForm({ ...changeForm, snils: e.target.value })
            }
          />
          <TextInput
            label="Организация"
            value={changeForm.organization}
            onChange={(e) =>
              setChangeForm({ ...changeForm, organization: e.target.value })
            }
          />
          <TextInput
            label="Должность"
            value={changeForm.position}
            onChange={(e) =>
              setChangeForm({ ...changeForm, position: e.target.value })
            }
          />
          <Select 
            options={[{value: 'root', content: 'root'}, {value: 'admin', content: 'admin'}, {value: 'user', content: 'user'}]}
            onUpdate={(e) =>
              // @ts-ignore
              setChangeForm({ ...changeForm, role: e.target[0].value })
            }
            label="Роль"
          />
          <Button loading={loading} onClick={createUser}>Создать</Button>
        </Flex>
      </Modal>
    </Flex>
  );
};

export default AdminUsers;
