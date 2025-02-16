import NgwMap from "@nextgis/ngw-ol";
import { useRef, useEffect } from "react";

// Добавим параметр для функции и состояния
const useMap = () => {
  const mapRef = useRef<NgwMap | null>(null);

  useEffect(() => {
    if (!mapRef.current) {
      const options = {
        baseUrl: "/api-map",
        target: "map",
        auth: {
          login: "hackathon_19",
          password: "hackathon_19_25",
        },
        adapterOptions: {
          selectable: true,
        },
      };

      const newMap = new NgwMap(options);
      mapRef.current = newMap;

      const vectorLayerStyle = 8892;

    newMap.addNgwLayer({
      resource: vectorLayerStyle,
      fit: true,
      adapterOptions: {
        selectable: true,
      },
    });

      newMap.onLoad().then(() => {
        console.log("Карта загружена");
      }).catch((e) => {
        console.error("Ошибка при загрузке карты:", e);
      });
    }
  }, []);

  return mapRef.current;
};

export default useMap;
