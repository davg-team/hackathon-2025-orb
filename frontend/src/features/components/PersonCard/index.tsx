import { Button, Icon } from "@gravity-ui/uikit";
import styles from "./index.module.css";
import { PencilToLine } from "@gravity-ui/icons";
import { Link } from "react-router-dom";
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

function PersonCard({
  editabe,
  accessable,
  person,
  fetchRequests
}: {
  editabe?: boolean;
  accessable?: boolean;
  person?: VeteranData;
  fetchRequests?: () => void
}) {  
  async function acceptRequest() {
      await fetch(`/api/records/${person?.id}/publish`, {
        method: "PATCH",
        headers: {
          Authorization: `Bearer ${Cookies.get("access_token")}`,
        }
      })
      if (fetchRequests) {
        fetchRequests()
      }
    }
  return (
    
      <div className={styles.personeCard}>
        <div>
          <p className={styles.title}><Link to={`/person/${person?.id}`} style={{ width: "100%" }}>{`${person?.name || "–"} ${
            person?.middle_name || "–"
          } ${person?.last_name || "–"}`}</Link></p>
          <p className={styles.text}>{person?.birth_date || "–"}</p>
          <p className={styles.text}>район {person?.birth_place || "–"}</p>
        </div>
        <div>
          {editabe && (
            <Button view="outlined">
              <Icon data={PencilToLine} />
              Изменить
            </Button>
          )}
          {accessable && (
            <>
              <Button view="action" onClick={acceptRequest} style={{ color: "white" }}>
                Принять
              </Button>{" "}
            </>
          )}
        </div>
      </div>
  );
}

export default PersonCard;
