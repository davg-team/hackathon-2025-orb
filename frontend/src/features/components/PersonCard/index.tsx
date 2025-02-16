import { Button, Icon } from "@gravity-ui/uikit";
import styles from "./index.module.css";
import { PencilToLine } from "@gravity-ui/icons";
import { Link } from "react-router-dom";

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
}: {
  editabe?: boolean;
  accessable?: boolean;
  person?: VeteranData;
}) {  
  return (
    <Link to={`/person/${person?.id}`} style={{ width: "100%" }}>
      <div className={styles.personeCard}>
        <div>
          <p className={styles.title}>{`${person?.name || "–"} ${
            person?.middle_name || "–"
          } ${person?.last_name || "–"}`}</p>
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
              <Button view="action" style={{ color: "white" }}>
                Принять
              </Button>{" "}
              / <Button view="outlined">отклонить</Button>
            </>
          )}
        </div>
      </div>
    </Link>
  );
}

export default PersonCard;
